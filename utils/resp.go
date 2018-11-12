package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResp responds error message in JSON format
func ErrorResp(w http.ResponseWriter, code int, message string) {
	JSONResp(w, code, map[string]string{"error": message})
}

// JSONResp responds in JSON format
func JSONResp(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// TextResp response in plain text
func TextResp(w http.ResponseWriter, code int, text string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(text))
}
