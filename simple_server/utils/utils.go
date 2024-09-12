package utils

import (
	"encoding/json"
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func BadRequest(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusBadRequest)
}

func Conflict(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusConflict)
}

func Unauthorized(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusUnauthorized)
}

func NotFound(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusInternalServerError)
}

func OK(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": message})
	if err != nil {
		return
	}
}
