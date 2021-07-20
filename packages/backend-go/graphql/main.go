package graphql

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"redinnlabs.com/redinn-core/auth"
)

type rootquery struct {
	UserQuery
	EnterpriseMutation
	EnterpriseQuery
}

type contextKey int

const User_id = contextKey(1)

func createDBContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}

func Setup(router *mux.Router) {
	// stitch schemas
	s := ""
	files, err := ioutil.ReadDir("./graphql/schema")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(files); i++ {
		content, e := ioutil.ReadFile("./graphql/schema/" + files[i].Name())
		if e != nil {
			panic(e)
		}
		s += string(content)
	}

	// setup the graphql package
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(s, &rootquery{}, opts...)
	g := relay.Handler{Schema: schema}

	router.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		// pass auth validate the token
		claims, err := auth.VerifyToken(r.Header.Get("Authorization"))
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			auth.SendResponse(w, http.StatusInternalServerError, &map[string]string{
				"message": err.Error(),
			})
			return
		}
		ctx := context.WithValue(context.Background(), User_id, claims.ID)
		g.ServeHTTP(w, r.WithContext(ctx))
	}).Methods("POST")
}
