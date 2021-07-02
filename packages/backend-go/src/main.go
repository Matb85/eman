package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	const GQLPATH = "/graphql/"

	router.HandleFunc(GQLPATH, g.ServeHTTP).Methods("POST")
	router.Methods("GET").PathPrefix(GQLPATH).Handler(http.StripPrefix(GQLPATH, http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":"+PORT, router)
}
