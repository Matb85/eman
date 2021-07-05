package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"redinnlabs.com/redinn-core/database"
)

func SendResponse(w http.ResponseWriter, status int, message *map[string]string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func Setup(router *mux.Router) {
	db := database.Connect()

	router.Use(jsonHeader)
	router.HandleFunc("/register", Register(db)).Methods("POST")
	router.HandleFunc("/login", Login(db)).Methods("POST")

}

func jsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
