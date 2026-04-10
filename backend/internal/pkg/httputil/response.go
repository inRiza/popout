package httputil

import (
	"net/http"
	"encoding/json"
)

type envelope struct {
	Data any	`json:"data,omitempty"`
	Error any	`json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(envelope{
		Data: data,
	})
}

func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(envelope{
		Error: message,
	})
}

func NoContent (w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}