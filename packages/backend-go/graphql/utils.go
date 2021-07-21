package graphql

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/database"
)

func findUser(id primitive.ObjectID) (*usergql, error) {
	dbctx, cancel := createDBContext()
	defer cancel()

	user := &usergql{}

	fetchErr := database.UserCol.FindOne(dbctx, bson.M{"_id": id}).Decode(&user)

	return user, fetchErr

}

func findEnterprise(id primitive.ObjectID) (*Enterprise, error) {
	dbctx, cancel := createDBContext()
	defer cancel()

	enterprise := &Enterprise{}
	finderr := database.EnterpriseCol.FindOne(dbctx, bson.M{"_id": id}).Decode(enterprise)

	return enterprise, finderr
}
