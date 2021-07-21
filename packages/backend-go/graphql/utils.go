package graphql

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/database"
)

func findUser(id primitive.ObjectID) (*userGQL, error) {
	dbctx, cancel := createDBContext()
	defer cancel()

	user := &userGQL{}

	fetchErr := database.UserCol.FindOne(dbctx, bson.M{"_id": id}).Decode(&user)

	return user, fetchErr

}

func findEnterprise(id primitive.ObjectID) (*enterpriseGQL, error) {
	dbctx, cancel := createDBContext()
	defer cancel()

	enterprise := &enterpriseGQL{}
	finderr := database.EnterpriseCol.FindOne(dbctx, bson.M{"_id": id}).Decode(enterprise)

	return enterprise, finderr
}
