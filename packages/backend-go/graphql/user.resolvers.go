package graphql

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

type UserQuery struct{}

func (r *UserQuery) User(ctx context.Context) (*User, error) {
	fmt.Println(ctx.Value("token"))
	return &User{"", "fdsfsd", "cos", "cos", &[]graphql.ID{"fdads"}}, nil
}

type User struct {
	ID          graphql.ID
	EMAIL       string
	FIRSTNAME   string
	LASTNAME    string
	ENTERPRISES *[]graphql.ID
}
