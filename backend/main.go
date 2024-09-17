package main

import (
	"backend/routes" // Import your routes package
	"log"
	"net/http"
)

func main() {
	log.Printf("The server is running correctly")
	// Register routes
	http.HandleFunc("/register", routes.SignupHandler)
	http.HandleFunc("/login", routes.SigninHandler)
	http.HandleFunc("/uploads", routes.UploadBase64Image)
	log.Println("Server starting on port 4000...")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
