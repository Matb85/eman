package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"redinnlabs.com/redinn-core/auth"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/graphql"
)

func main() {
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}
	fmt.Println("listening on port " + PORT)

	// connect to the database
	database.Connect()

	router := mux.NewRouter()

	graphql.Setup(router)
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
