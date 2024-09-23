package routes

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary

func init() {
	var err error

	// Load Cloudinary parameters from environment variables

	cloudinaryName := os.Getenv("CLOUDINARY_NAME")
	cloudinaryAPIKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudinaryAPISecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err = cloudinary.NewFromParams(cloudinaryName, cloudinaryAPIKey, cloudinaryAPISecret)
	if err != nil {
		log.Fatalf("Failed to create Cloudinary instance: %v", err)
	}
}

type UploadRequest struct {
	Input struct {
		Base64Str string `json:"base64_str"`
	} `json:"input"`
}

func UploadBase64Image(w http.ResponseWriter, r *http.Request) {
	var request UploadRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	prefix := "data:image/png;base64,"
	if strings.HasPrefix(request.Input.Base64Str, prefix) {
		request.Input.Base64Str = strings.TrimPrefix(request.Input.Base64Str, prefix)
	}

	decodedImage, err := base64.StdEncoding.DecodeString(request.Input.Base64Str)
	if err != nil {
		log.Printf("Failed to decode base64 string: %v", err)
		http.Error(w, "Failed to decode base64 string", http.StatusInternalServerError)
		return
	}

	uploadResp, err := cld.Upload.Upload(r.Context(), bytes.NewReader(decodedImage), uploader.UploadParams{Folder: "images"})
	if err != nil {
		log.Printf("Error while uploading to Cloudinary: %v", err)
		http.Error(w, "Failed to upload image to Cloudinary", http.StatusInternalServerError)
		return
	}

	if uploadResp.SecureURL == "" {
		log.Printf("Cloudinary upload response did not return a valid URL: %+v", uploadResp)
		http.Error(w, "Upload response missing URL", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"url": uploadResp.SecureURL}
	responseData, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
