package graphql

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserQuery struct{}

func (r *UserQuery) User(ctx context.Context) (*userGQL, error) {
	return findUser(ctx.Value(User_id).(primitive.ObjectID))
}
