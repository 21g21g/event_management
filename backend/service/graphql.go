package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
	"golang.org/x/crypto/bcrypt"
)

// http://graphql-engine:8080/v1/graphql
var client = graphql.NewClient("http://graphql-engine:8080/v1/graphql")

// var jwtSecret = []byte("your-secret-key") // Replace with a secure key

func RunGraphQLQuery(ctx context.Context, userQuery *graphql.Request, userResp interface{}) error {

	err := client.Run(ctx, userQuery, userResp)
	if err != nil {
		log.Println("GraphQL query failed:", err)
		return err
	}
	log.Printf("GraphQL query succeeded. Response: %+v\n", userResp)

	return nil
}

func UserExists(email string) (bool, error) {
	req := graphql.NewRequest(`
	query user($email:String!){
		users(where:{email:{_eq:$email}}){
		  id
		  email
		}
	  }
    `)
	req.Var("email", email)

	var respData struct {
		Data struct {
			Users []struct {
				ID string `json:"id"`
			} `json:"users"`
		} `json:"data"`
	}

	ctx := context.Background()
	if err := RunGraphQLQuery(ctx, req, &respData); err != nil {
		return false, err
	}
	return len(respData.Data.Users) > 0, nil
}

func RegisterUser(username, email, password string) error {
	exists, err := UserExists(email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req := graphql.NewRequest(`
	mutation registerUser($email:String!,$username:String!,$password:String!){
		insert_users(objects:{email:$email,username:$username,password:$password}){
		  returning{
			username
		  }
		}
	  }
    `)
	req.Var("username", username)
	req.Var("email", email)
	req.Var("password", string(hashedPassword))

	var respData struct {
		Data struct {
			RegisterUser struct {
				Username string `json:"username"`
			} `json:"registerUser"`
		} `json:"data"`
	}

	ctx := context.Background()
	if err := RunGraphQLQuery(ctx, req, &respData); err != nil {
		return err
	}
	return nil
}

var jwtSecretKey []byte

type JWTSecret struct {
	Type                string `json:"type"`
	Key                 string `json:"key"`
	ClaimsNamespacePath string `json:"claims_namespace_path"`
}

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Parse the JWT secret from environment
	jwtSecretEnv := os.Getenv("HASURA_GRAPHQL_JWT_SECRET")
	if jwtSecretEnv == "" {
		log.Fatal("HASURA_GRAPHQL_JWT_SECRET not set in environment")
	}

	var jwtSecret JWTSecret
	if err := json.Unmarshal([]byte(jwtSecretEnv), &jwtSecret); err != nil {
		log.Fatalf("Error parsing HASURA_GRAPHQL_JWT_SECRET: %v", err)
	}
	jwtSecretKey = []byte(jwtSecret.Key)
}
func LoginUser(email, password string) (string, error) {
	fmt.Printf("email is: %+v\n", email)
	fmt.Printf("password is: %+v\n", password)
	req := graphql.NewRequest(`
    query user($email:String!){
        users(where:{email:{_eq:$email}}){
          id
          password
        }
      }
    `)
	req.Var("email", email)
	//this is used to handle the response comes from graphql server.
	var respData struct {
		Data struct {
			Users []struct {
				ID       string `json:"id"`
				Password string `json:"password"`
			} `json:"users"`
		} `json:"data"`
	}

	ctx := context.Background()
	if err := RunGraphQLQuery(ctx, req, &respData); err != nil {
		return "", err
	}

	if len(respData.Data.Users) == 0 {
		return "", errors.New("invalid email or password")
	}

	storedPassword := respData.Data.Users[0].Password
	userID := respData.Data.Users[0].ID

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
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expiry time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
