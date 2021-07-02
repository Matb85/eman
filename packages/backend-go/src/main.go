package main

import (
	"fmt"
	"net/http"
	"os"

	"redinnlabs.com/redinn-core/graphql"
)

func main() {
	PORT := os.Getenv("PORT")
	fmt.Println("listening on port " + PORT)

	http.Handle("/graphql", graphql.Setup())
	http.ListenAndServe(":"+PORT, nil)
}
