package graphql

import (
	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type rootResolvers struct {
	UserResolvers
	EnterpriseResolvers
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
	schema := graphql.MustParseSchema(s, &rootResolvers{}, opts...)
	g := relay.Handler{Schema: schema}

	router.HandleFunc("/graphql", g.ServeHTTP).Methods("POST")
}
