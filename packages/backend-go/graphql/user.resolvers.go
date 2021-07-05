package graphql

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/database"
)

type UserQuery struct{}

func (r *UserQuery) User(ctx context.Context) (*User, error) {
	dbctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := &User{}

	id, _ := primitive.ObjectIDFromHex(ctx.Value("user_id").(string))

	fetchErr := database.UserCol.FindOne(dbctx, bson.M{"_id": id}).Decode(&user)
	if fetchErr != nil {
		return &User{}, fetchErr
	}
	return user, nil
}

type User struct {
	ID          graphql.ID    `json:"id" bson:"_id,omitempty"`
	EMAIL       string        `json:"email"`
	FIRSTNAME   string        `json:"firstname"`
	LASTNAME    string        `json:"lastname"`
	ENTERPRISES *[]graphql.ID `json:"enterprises"`
}
