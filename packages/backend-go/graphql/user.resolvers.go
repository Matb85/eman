package graphql

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/database"
)

type UserQuery struct{}

func (r *UserQuery) User(ctx context.Context) (*usergql, error) {
	dbctx, cancel := createDBContext()
	defer cancel()

	user := &usergql{}

	fetchErr := database.UserCol.FindOne(dbctx, bson.M{"_id": ctx.Value(User_id).(primitive.ObjectID)}).Decode(&user)
	if fetchErr != nil {
		return &usergql{}, fetchErr
	}
	return user, nil
}

type usergql struct {
	ID          graphql.ID    `json:"id" bson:"_id,omitempty"`
	EMAIL       string        `json:"email"`
	FIRSTNAME   string        `json:"firstname"`
	LASTNAME    string        `json:"lastname"`
	ENTERPRISES *[]graphql.ID `json:"enterprises"`
}
