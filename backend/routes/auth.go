package routes

import (
	"backend/service"
	"backend/utils"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"input"`
}

type JWTSecret struct {
	Type                string `json:"type"`
	Key                 string `json:"key"`
	ClaimsNamespacePath string `json:"claims_namespace_path"`
}

var (
	client       *graphql.Client
	jwtSecretKey []byte
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	jwtSecretEnv := os.Getenv("HASURA_GRAPHQL_JWT_SECRET")
	if jwtSecretEnv == "" {
		log.Fatal("HASURA_GRAPHQL_JWT_SECRET not set in environment")
	}

	var jwtSecret JWTSecret
	if err := json.Unmarshal([]byte(jwtSecretEnv), &jwtSecret); err != nil {
		log.Fatalf("Error parsing HASURA_GRAPHQL_JWT_SECRET: %v", err)
	}
	jwtSecretKey = []byte(jwtSecret.Key)

	hasuraURL := "http://graphql-engine:8080/v1/graphql"
	client = graphql.NewClient(hasuraURL)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Input struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"input"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	log.Printf("Request body: %+v", reqBody.Input)

	err := service.RegisterUser(reqBody.Input.Username, reqBody.Input.Email, reqBody.Input.Password)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.SendJSONResponse(w, http.StatusConflict, map[string]string{"message": "Email already exists"})
		} else {
			utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Failed to register user"})
		}
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, map[string]string{"message": "User registered successfully"})
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody SignInRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	email := reqBody.Input.Email
	password := reqBody.Input.Password

	token, err := LoginUser(email, password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func LoginUser(email, password string) (string, error) {
	req := graphql.NewRequest(`
		query user($email: String!) {
			users(where: {email: {_eq: $email}}) {
				id
				password
			}
		}
	`)
	req.Var("email", email)

	var respData struct {
		Users []struct {
			ID       string `json:"id"`
			Password string `json:"password"`
		} `json:"users"`
	}

	ctx := context.Background()
	if err := RunGraphQLQuery(ctx, req, &respData); err != nil {
		return "", err
	}

	if len(respData.Users) == 0 {
		return "", errors.New("invalid email or password")
	}

	storedPassword := respData.Users[0].Password
	userID := respData.Users[0].ID

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := generateJWT(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": []string{"user", "admin", "anonymous"},
			"x-hasura-default-role":  "anonymous",
			"x-hasura-user-id":       userID,
		},
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RunGraphQLQuery(ctx context.Context, req *graphql.Request, respData interface{}) error {
	if err := client.Run(ctx, req, respData); err != nil {
		return err
	}
	return nil
}
