package main

import (
	"context"
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

	g := graphql.Setup()
	const (
		GQLPATH     = "/graphql"
		FULLGQLPATH = GQLPATH + "/"
	)
	router.HandleFunc(GQLPATH, func(w http.ResponseWriter, r *http.Request) {
		// pass auth validate the token
		claims, err := auth.VerifyToken(r.Header.Get("Authorization"))
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			auth.SendResponse(w, http.StatusInternalServerError, &map[string]string{
				"message": err.Error(),
			})
			return
		}
		ctx := context.WithValue(context.Background(), "user_id", claims.ID)
		g.ServeHTTP(w, r.WithContext(ctx))
	}).Methods("POST")

	router.Methods("GET").PathPrefix(FULLGQLPATH).Handler(http.StripPrefix(FULLGQLPATH, http.FileServer(http.Dir("./static"))))

	auth.Setup(router.PathPrefix("/auth").Subrouter())
	http.ListenAndServe(":"+PORT, router)
}
