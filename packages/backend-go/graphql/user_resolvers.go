package graphql

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"matb85.github.io/eman-core/utils"
)

type UserResolvers struct{}

func (r *UserResolvers) User(ctx context.Context) (*UserGQL, error) {
	return FindUser(ctx.Value(utils.User_id).(primitive.ObjectID))
}
