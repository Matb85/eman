package graphql

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/utils"
)

func FindUser(id primitive.ObjectID) (*UserGQL, error) {
	dbctx, cancel := utils.CreateDBContext()
	defer cancel()

	user := &UserGQL{}

	fetchErr := database.UserCol.FindOne(dbctx, bson.M{"_id": id}).Decode(&user)

	return user, fetchErr

}

func FindEnterprise(id primitive.ObjectID) (*EnterpriseGQL, error) {
	dbctx, cancel := utils.CreateDBContext()
	defer cancel()

	enterprise := &EnterpriseGQL{}
	finderr := database.EnterpriseCol.FindOne(dbctx, bson.M{"_id": id}).Decode(enterprise)

	return enterprise, finderr
}
