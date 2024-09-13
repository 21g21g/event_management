package utils
import (
    "encoding/json"
    "net/http"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    // Encode the data as JSON and write it to the response writer
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}
