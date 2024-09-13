

// SendJSONResponse sends a well-structured JSON response with a flexible interface type
package utils

import (
	"encoding/json"
	"net/http"
)

// SendJSONResponses sends a well-structured JSON response
func SendJSONResponses(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// Use json.Encoder to ensure proper JSON formatting
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// In case of an encoding failure, send an internal server error
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
