package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]interface{}{
		"status":  statusCode,
		"message": message,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
