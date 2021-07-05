package main

import (
	"fmt"
	"net/http"
	"os"

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

	http.ListenAndServe(":"+PORT, router)
}
