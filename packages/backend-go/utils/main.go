package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type contextKey int

const User_id = contextKey(1)

type Message struct {
	Message string `json:"message"`
}

func SendResponse(w http.ResponseWriter, status int, message *map[string]string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func CreateDBContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}
