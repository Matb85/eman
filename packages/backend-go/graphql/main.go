package graphql

import (
	"io/ioutil"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type rootquery struct {
	UserQuery
}

func Setup() *relay.Handler {
	// stich schemas
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
	return &relay.Handler{Schema: schema}
}
