package graphql

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"matb85.github.io/eman-core/database"
	"matb85.github.io/eman-core/utils"
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
func FindEnterpriseByIndex(user_id primitive.ObjectID, index int) (*primitive.ObjectID, error) {
	// fetch the user
	user, fetchErr := FindUser(user_id)
	if fetchErr != nil {
		return nil, fetchErr
	}
	// get the id of the enterprise
	id, primerr := primitive.ObjectIDFromHex(string((*user.Enterprises)[index]))
	if primerr != nil {
		return nil, primerr
	}
	// get the id of the enterprise and fetch it
	return &id, nil
}
