package graphql

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResolvers struct{}

func (r *UserResolvers) User(ctx context.Context) (*UserGQL, error) {
	return FindUser(ctx.Value(User_id).(primitive.ObjectID))
}
