package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"matb85.github.io/eman-core/assets"
	"matb85.github.io/eman-core/auth"
	"matb85.github.io/eman-core/database"
	"matb85.github.io/eman-core/graphql"
	"matb85.github.io/eman-core/utils"
)

func main() {
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}
	fmt.Println("listening on port " + PORT)

	// connect to the database
	database.Connect()

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	api := router.PathPrefix("/secure").Subrouter()
	api.Use(authMiddleware)

	graphql.Setup(api)
	assets.Setup(api)
	auth.Setup(router.PathPrefix("/auth").Subrouter())

	router.Use(wrapHandlerWithLogging)
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// auth - validate the token
		claims, err := auth.VerifyToken(r.Header.Get("Authorization"))
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			utils.SendMessage(w, http.StatusBadRequest, err.Error())
			return
		}
		ctx := context.WithValue(context.Background(), utils.User_id, claims.ID)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)

		statusCode := lrw.statusCode
		log.Printf("- %s %s %d %s", req.Method, req.URL.Path, statusCode, http.StatusText(statusCode))
	})
}
