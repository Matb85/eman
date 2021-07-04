package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"redinnlabs.com/redinn-core/auth"
	"redinnlabs.com/redinn-core/graphql"
)

func main() {
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}
	fmt.Println("listening on port " + PORT)

	router := mux.NewRouter()

	g := graphql.Setup()
	const (
		GQLPATH     = "/graphql"
		FULLGQLPATH = GQLPATH + "/"
	)
	router.HandleFunc(GQLPATH, func(w http.ResponseWriter, r *http.Request) {
		// pass auth token to resolvers
		ctx := context.WithValue(context.Background(), "token", "to jest klucz")
		g.ServeHTTP(w, r.WithContext(ctx))
	}).Methods("POST")

	router.Methods("GET").PathPrefix(FULLGQLPATH).Handler(http.StripPrefix(FULLGQLPATH, http.FileServer(http.Dir("./static"))))

	auth.Setup(router.PathPrefix("/auth").Subrouter())
	http.ListenAndServe(":"+PORT, router)
}
