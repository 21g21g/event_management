package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid" // Import the UUID package
	"github.com/joho/godotenv"
	// "github.com/machinebox/graphql"
)

// Struct to match the Chapa API success response format
type ChapaResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}

// Struct for receiving payment details from the client request
type Payment struct {
	Input struct {
		PhoneNumber string `json:"phoneNumber"`
		Amount      string `json:"amount"`
	} `json:"input"`
}

// var clients = graphql.NewClient("http://graphql-engine:8080/v1/graphql")

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func PaymentsHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming request body
	var payment Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Request body: %+v", payment.Input)

	// Generate a dynamic tx_ref using UUID
	txRef := "chewatatest-" + uuid.New().String()

	url := "https://api.chapa.co/v1/transaction/initialize"
	method := "POST"

	// Prepare payload with dynamic tx_ref and user inputs
	const returnURL = "http://localhost:3000/user"
	payload := map[string]string{
		"amount":                     payment.Input.Amount,
		"currency":                   "ETB",
		"email":                      "assegagebeyehu21@gmail.com", // Can be dynamic
		"first_name":                 "Gebeyehu",                   // Can be dynamic
		"last_name":                  "assega",                     // Can be dynamic
		"phone_number":               payment.Input.PhoneNumber,
		"tx_ref":                     txRef, // Use the dynamic tx_ref
		"callback_url":               "https://webhook.site/077164d6-29cb-40df-ba29-8a00e59a7e60",
		"return_url":                 returnURL,
		"customization[title]":       "Payment for my favourite merchant",
		"customization[description]": "I love online payments",
		"meta[hide_receipt]":         "true",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payloadBytes))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Load the Chapa API key from the environment variable
	chapaKey := os.Getenv("CHAPA_API_KEY")
	if chapaKey == "" {
		http.Error(w, "CHAPA_API_KEY is not set in the environment", http.StatusInternalServerError)
		return
	}
	log.Printf("chapakey: %+v", chapaKey)

	// Set request headers
	req.Header.Add("Authorization", "Bearer "+os.Getenv("CHAPA_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	// Unmarshal the response into the ChapaResponse struct
	var chapaResponse ChapaResponse
	err = json.Unmarshal(body, &chapaResponse)
	if err != nil {
		http.Error(w, "Error unmarshalling response", http.StatusInternalServerError)
		return
	}

	// Check if the response status is "success"
	if chapaResponse.Status == "success" {
		// Send response back to the client
		response := map[string]interface{}{
			"message":     chapaResponse.Message,
			"tx_ref":      txRef,
			"checkoutUrl": chapaResponse.Data.CheckoutURL,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		// Handle error case from Chapa API
		http.Error(w, "Error in payment initialization", http.StatusBadGateway)
		fmt.Println("Response:", string(body)) // Print full response for debugging
	}
}
