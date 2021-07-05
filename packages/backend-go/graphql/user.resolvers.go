package graphql

import (
	"context"
	"fmt"
	"time"

	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
	"redinnlabs.com/redinn-core/database"
)

type UserQuery struct{}

func (r *UserQuery) User(ctx context.Context) (*User, error) {
	fmt.Println(ctx.Value("user_id"))

	dbctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.Connect()

	user := &User{}

	fetchErr := db.Collection("users").FindOne(dbctx, bson.M{"uuid": ctx.Value("user_id")}).Decode(&user)
	fmt.Println(user)
	if fetchErr != nil {
		return &User{}, fetchErr
	}
	fmt.Println(ctx.Value("user_id"))
	return user, nil
}

type User struct {
	ID          graphql.ID
	EMAIL       string
	FIRSTNAME   string
	LASTNAME    string
	ENTERPRISES *[]graphql.ID
}
