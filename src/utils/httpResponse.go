package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HttpResponse(data interface{}, w http.ResponseWriter, statusCode int) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal response to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
	fmt.Printf("%d - ", statusCode)
}
