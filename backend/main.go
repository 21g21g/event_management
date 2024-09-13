// package main

// import (
// 	"backend/middleware"
// 	"backend/service"
// 	"backend/utils"
// 	"bytes"
// 	"context"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/cloudinary/cloudinary-go/v2"
// 	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
// 	"github.com/google/uuid"
// )

// func main() {
// 	// Define HTTP routes for the web server
// 	http.HandleFunc("/register", registerHandler)
// 	http.HandleFunc("/login", loginHandler)
// 	http.Handle("/protected", middleware.SessionAuthMiddleware(http.HandlerFunc(protectedHandler)))
// 	http.HandleFunc("/upload", uploadHandler) // New upload handler route

// 	log.Println("Server starting on port 4000...")
// 	if err := http.ListenAndServe(":4000", nil); err != nil {
// 		log.Fatalf("Server failed: %v", err)
// 	}
// }

// // Register user handler
// type HasuraRequestBody struct {
// 	Input struct {
// 		Username string `json:"username"`
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	} `json:"input"`
// }

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	var reqBody HasuraRequestBody

// 	// Decode the request body into the struct
// 	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
// 		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid request body")
// 		return
// 	}

// 	log.Printf("Request body: %+v", reqBody)

// 	// Access the username, email, and password from the nested input struct
// 	err := service.RegisterUser(reqBody.Input.Username, reqBody.Input.Email, reqBody.Input.Password)
// 	if err != nil {
// 		if err.Error() == "user already exists" {
// 			utils.SendJSONResponse(w, http.StatusConflict, "Email already exists")
// 		} else {
// 			utils.SendJSONResponse(w, http.StatusInternalServerError, "Failed to register user")
// 		}
// 		return
// 	}

// 	// Send a success response
// 	utils.SendJSONResponse(w, http.StatusOK, "User registered successfully")
// }

// // Login user handler
// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	var reqBody struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
// 		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid request")
// 		return
// 	}

// 	token, err := service.LoginUser(reqBody.Email, reqBody.Password)
// 	if err != nil {
// 		utils.SendJSONResponse(w, http.StatusUnauthorized, "Invalid email or password")
// 		return
// 	}

// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "session_token",
// 		Value:    token,
// 		Path:     "/",
// 		HttpOnly: true,
// 	})

// 	utils.SendJSONResponse(w, http.StatusOK, "Login successful")
// }

// // Protected handler that requires authentication
// func protectedHandler(w http.ResponseWriter, r *http.Request) {
// 	utils.SendJSONResponse(w, http.StatusOK, "Access granted")
// }

// // UploadRequest represents the expected JSON request body for file upload
// type UploadRequest struct {
// 	Base64Str string `json:"base64_str"`
// }

// // UploadResponse represents the JSON response body for file upload
// type UploadResponse struct {
// 	URL string `json:"url"`
// }

// // UploadHandler handles the upload of a base64 image to Cloudinary
// func uploadHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var req UploadRequest
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&req); err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Decode the base64 image
// 	fileBytes, err := decodeBase64(req.Base64Str)
// 	if err != nil {
// 		http.Error(w, "Failed to decode base64 string", http.StatusBadRequest)
// 		return
// 	}

// 	// Initialize Cloudinary credentials
// 	cld, ctx := credentials()

// 	// Upload the base64 image and get the URL
// 	url, err := uploadBase64Image(cld, ctx, fileBytes)
// 	if err != nil {
// 		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
// 		return
// 	}

// 	if url == "" {
// 		http.Error(w, "Received empty URL from Cloudinary", http.StatusInternalServerError)
// 		return
// 	}

// 	// Prepare the response
// 	response := UploadResponse{URL: url}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }

// // decodeBase64 decodes a base64 string into bytes
// func decodeBase64(data string) ([]byte, error) {
// 	// Remove any data URL prefix
// 	if strings.HasPrefix(data, "data:") {
// 		data = strings.SplitN(data, ",", 2)[1]
// 	}
// 	return base64.StdEncoding.DecodeString(data)
// }

// // credentials initializes and returns the Cloudinary client and context
// func credentials() (*cloudinary.Cloudinary, context.Context) {
// 	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
// 	if err != nil {
// 		fmt.Printf("Error initializing Cloudinary: %v\n", err)
// 	}
// 	ctx := context.Background()
// 	return cld, ctx
// }

// // uploadBase64Image uploads the binary image data to Cloudinary
// func uploadBase64Image(cld *cloudinary.Cloudinary, ctx context.Context, fileBytes []byte) (string, error) {
// 	// Generate a unique identifier for the file
// 	publicID := "file-" + uuid.NewString()

// 	uploadResult, err := cld.Upload.Upload(ctx, bytes.NewReader(fileBytes), uploader.UploadParams{
// 		Folder:       "event",
// 		PublicID:     publicID,
// 		ResourceType: "auto",
// 	})
// 	if err != nil {
// 		fmt.Printf("Upload error: %v\n", err)
// 		return "", fmt.Errorf("upload error: %v", err)
// 	}

// 	if uploadResult.SecureURL == "" {
// 		fmt.Println("Empty URL received from Cloudinary")
// 	}

//		return uploadResult.SecureURL, nil
//	}
package main

import (
	"backend/service"
	"backend/utils"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	log.Printf("The server is running correctly")
	http.HandleFunc("/register", signupHandler)
	http.HandleFunc("/login", signinHandler)
	http.HandleFunc("/uploads", uploadBase64Image)
	log.Println("Server starting on port 4000...")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// Register handler adjusted for Hasura's expected input structure
func signupHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Input struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"input"`
	}

	// Decode the request body (specifically the input field)
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	log.Printf("Request body: %+v", reqBody.Input) // Debugging purpose

	// Pass the decoded values from reqBody.Input to service.RegisterUser
	err := service.RegisterUser(reqBody.Input.Username, reqBody.Input.Email, reqBody.Input.Password)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.SendJSONResponse(w, http.StatusConflict, map[string]string{"message": "Email already exists"})
		} else {
			utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Failed to register user"})
		}
		return
	}

	// Send success response in the format of GraphQL mutation
	utils.SendJSONResponse(w, http.StatusOK, map[string]string{"message": "User registered successfully"})
}

// Login handler adjusted for Hasura's expected input structure
// func signinHandler(w http.ResponseWriter, r *http.Request) {
// 	var reqBody struct {
// 		Input struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		} `json:"input"`
// 	}

// 	// Decode the request body (specifically the input field)
// 	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
// 		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	// Call the service to handle user login
// 	token, err := service.LoginUser(reqBody.Input.Email, reqBody.Input.Password)
// 	if err != nil {
// 		utils.SendJSONResponse(w, http.StatusUnauthorized, "Invalid email or password")
// 		return
// 	}

// 	// Set the JWT token as a cookie
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "session_token",
// 		Value:    token,
// 		Path:     "/",
// 		HttpOnly: true,
// 	})

// 	utils.SendJSONResponse(w, http.StatusOK, "Login successful")
// }

//	func protectedHandler(w http.ResponseWriter, r *http.Request) {
//		utils.SendJSONResponse(w, http.StatusOK, "Access granted")
//	}

type SignInRequest struct {
	Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"input"`
}

// JWTSecret struct for parsing Hasura JWT secret
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
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Parse the JWT secret from the environment
	jwtSecretEnv := os.Getenv("HASURA_GRAPHQL_JWT_SECRET")
	if jwtSecretEnv == "" {
		log.Fatal("HASURA_GRAPHQL_JWT_SECRET not set in environment")
	}

	// Unmarshal the JWT secret
	var jwtSecret JWTSecret
	if err := json.Unmarshal([]byte(jwtSecretEnv), &jwtSecret); err != nil {
		log.Fatalf("Error parsing HASURA_GRAPHQL_JWT_SECRET: %v", err)
	}
	jwtSecretKey = []byte(jwtSecret.Key)

	// Directly set the GraphQL client URL
	hasuraURL := "http://graphql-engine:8080/v1/graphql"

	// Initialize the GraphQL client
	client = graphql.NewClient(hasuraURL)
}

// signinHandler handles the login requests from Hasura actions
func signinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody SignInRequest

	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Extract user input
	email := reqBody.Input.Email
	password := reqBody.Input.Password

	// Validate user login and generate token
	token, err := LoginUser(email, password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Respond with the JWT token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// LoginUser handles the logic of querying the user from the database and validating the password
func LoginUser(email, password string) (string, error) {
	fmt.Printf("email is: %+v\n", email)

	// Create GraphQL query to fetch the user by email
	req := graphql.NewRequest(`
		query user($email: String!) {
			users(where: {email: {_eq: $email}}) {
				id
				password
			}
		}
	`)

	req.Var("email", email)

	// GraphQL response struct
	var respData struct {
		Users []struct {
			ID       string `json:"id"`
			Password string `json:"password"`
		} `json:"users"`
	}

	// Execute GraphQL query
	ctx := context.Background()
	if err := RunGraphQLQuery(ctx, req, &respData); err != nil {
		return "", err
	}

	// Check if the user exists
	if len(respData.Users) == 0 {
		return "", errors.New("invalid email or password")
	}

	// Verify password
	storedPassword := respData.Users[0].Password
	userID := respData.Users[0].ID

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := generateJWT(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateJWT generates a JWT token for the authenticated user
func generateJWT(userID string) (string, error) {
	// Set claims for the token
	claims := jwt.MapClaims{
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": []string{"user", "admin", "anonymous"},
			"x-hasura-default-role":  "user",
			"x-hasura-user-id":       userID,
		},
		"exp": time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// RunGraphQLQuery executes a GraphQL query
func RunGraphQLQuery(ctx context.Context, req *graphql.Request, respData interface{}) error {
	// Send the request to the GraphQL API and get the response
	if err := client.Run(ctx, req, respData); err != nil {
		return err
	}
	return nil
}

// RunGraphQLQuery executes a GraphQL query and handles the response
// func RunGraphQLQuery(ctx context.Context, userQuery *graphql.Request, userResp interface{}) error {
// 	err := client.Run(ctx, userQuery, userResp)
// 	if err != nil {
// 		log.Println("GraphQL query failed:", err)
// 		return err
// 	}
// 	log.Printf("GraphQL query succeeded. Response: %+v\n", userResp)
// 	return nil
// }

//this is for upload image functionality

var cld *cloudinary.Cloudinary

func init() {
	// cloudName := os.Getenv("CLOUDINARY_NAME")
	// apiKey := os.Getenv("CLOUDINARY_API_KEY")
	// apiSecret := os.Getenv("API_KEY")
	var err error
	cld, err = cloudinary.NewFromParams("dnfl1adwg", "988944114923561", "_ejh9K0mvcD-o6P2nedgE7XUWV4")
	if err != nil {
		log.Fatalf("Failed to create Cloudinary instance: %v", err)
	}
}

type UploadRequest struct {
	Input struct {
		Base64Str string `json:"base64_str"`
	} `json:"input"`
}

func uploadBase64Image(w http.ResponseWriter, r *http.Request) {
	var request UploadRequest

	// Decode JSON request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	prefix := "data:image/png;base64,"
	if strings.HasPrefix(request.Input.Base64Str, prefix) {
		request.Input.Base64Str = strings.TrimPrefix(request.Input.Base64Str, prefix)
	}

	// Decode base64 image
	decodedImage, err := base64.StdEncoding.DecodeString(request.Input.Base64Str)
	if err != nil {
		log.Printf("Failed to decode base64 string: %v", err)
		http.Error(w, "Failed to decode base64 string", http.StatusInternalServerError)
		return
	}

	// Upload to Cloudinary
	uploadResp, err := cld.Upload.Upload(r.Context(), bytes.NewReader(decodedImage), uploader.UploadParams{Folder: "images"})
	if err != nil {
		log.Printf("Error while uploading to Cloudinary: %v", err)
		http.Error(w, "Failed to upload image to Cloudinary", http.StatusInternalServerError)
		return
	}

	// Check if the URL is empty
	if uploadResp.SecureURL == "" {
		log.Printf("Cloudinary upload response did not return a valid URL: %+v", uploadResp)
		http.Error(w, "Upload response missing URL", http.StatusInternalServerError)
		return
	}

	// Send response with image URL
	response := map[string]string{"url": uploadResp.SecureURL}
	responseData, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

// func init() {
// 	// Load environment variables from .env file
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}

// 	// Retrieve Cloudinary credentials from environment variables
// 	cloudName := os.Getenv("CLOUDINARY_NAME")
// 	apiKey := os.Getenv("CLOUDINARY_API_KEY")
// 	apiSecret := os.Getenv("API_KEY")

// 	// Initialize Cloudinary client
// 	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)

// 	if err != nil {
// 		log.Fatalf("Error initializing Cloudinary client: %v", err)
// 	}
// 	cloudinaryClient = cld
// }

// func uploadImagesHandler(w http.ResponseWriter, r *http.Request) {
// 	var requestBody struct {
// 		Input struct {
// 			Images []string `json:"images"`
// 		} `json:"input"`
// 	}

// 	// Decode the request body
// 	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	// Upload images to Cloudinary
// 	var uploadedUrls []string
// 	for _, imageURL := range requestBody.Input.Images {
// 		resp, err := cloudinaryClient.Upload.Upload(context.Background(), imageURL, uploader.UploadParams{})
// 		if err != nil {
// 			http.Error(w, "Error uploading image", http.StatusInternalServerError)
// 			return
// 		}
// 		uploadedUrls = append(uploadedUrls, resp.SecureURL)
// 	}

// 	// Send response
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string][]string{"urls": uploadedUrls})
// }

// func signinHandler(w http.ResponseWriter, r *http.Request) {
// 	var reqBody struct {
// 		Input struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		} `json:"input"`
// 	}

// 	// Decode the request body
// 	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
// 		// Send a properly structured JSON error response
// 		utils.SendJSONResponses(w, http.StatusBadRequest, map[string]interface{}{
// 			"errors": []map[string]interface{}{
// 				{
// 					"message": "Invalid request payload",
// 					"extensions": map[string]interface{}{
// 						"code": "bad-request",
// 					},
// 				},
// 			},
// 		})
// 		return
// 	}

// 	// Call the service to handle user login
// 	token, err := service.LoginUser(reqBody.Input.Email, reqBody.Input.Password)
// 	if err != nil {
// 		// Send a properly structured JSON error response
// 		utils.SendJSONResponses(w, http.StatusUnauthorized, map[string]interface{}{
// 			"errors": []map[string]interface{}{
// 				{
// 					"message": "Invalid email or password",
// 					"extensions": map[string]interface{}{
// 						"code": "unauthorized",
// 					},
// 				},
// 			},
// 		})
// 		return
// 	}

// 	// Return the JWT token in the response body in the expected format
// 	utils.SendJSONResponses(w, http.StatusOK, map[string]interface{}{
// 		"data": map[string]interface{}{
// 			"token": token,
// 		},
// 	})
// }
