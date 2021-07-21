package graphql

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserQuery struct{}

func (r *UserQuery) User(ctx context.Context) (*usergql, error) {
	return findUser(ctx.Value(User_id).(primitive.ObjectID))
}

type usergql struct {
	ID          graphql.ID    `json:"id" bson:"_id,omitempty"`
	EMAIL       string        `json:"email"`
	FIRSTNAME   string        `json:"firstname"`
	LASTNAME    string        `json:"lastname"`
	ENTERPRISES *[]graphql.ID `json:"enterprises"`
}
