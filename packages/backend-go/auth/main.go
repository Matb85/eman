package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {
	router.Use(jsonHeader)
	router.HandleFunc("/register", Register).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")

}

func jsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
