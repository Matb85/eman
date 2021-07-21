package graphql

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserQuery struct{}

func (r *UserQuery) User(ctx context.Context) (*UserGQL, error) {
	return FindUser(ctx.Value(User_id).(primitive.ObjectID))
}
