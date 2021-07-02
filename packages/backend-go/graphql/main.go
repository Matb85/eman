package graphql

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)


type query struct{}

func (*query) Hello() string { return "Hello, world!" }

func Setup() *relay.Handler {
	
	s := `
			type Query {
					hello: String!
			}
	`
	schema := graphql.MustParseSchema(s, &query{})
	return &relay.Handler{Schema: schema}
}