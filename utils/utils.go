package utils

import (
	"encoding/json"
	"net/http"
)

type contextKey string

const ContextKeyUser contextKey = "user_id"

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{"data": data})
}

func WriteJSONError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
