package graphql_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"matb85.github.io/eman-core/assets"
	"matb85.github.io/eman-core/database"
	"matb85.github.io/eman-core/graphql"
	"matb85.github.io/eman-core/utils"
)

var resolvers = &graphql.EnterpriseResolvers{}

func TestAddEnterpriseOK(t *testing.T) {
	callback, result, adderr := createUserAndEnterprise()
	defer callback(t)
	if adderr != nil {
		t.Fatal(adderr)
	}
	if result.enterprise.Name != MockedEnterprise.Name {
		t.Error("given and created enterprise are not equal", result.enterprise.Name, MockedEnterprise.Name)
	}
	folderPath := assets.ASSETS_DIR + result.enterprise.Folder
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		t.Error("the assets folder for the enterprise does not exist: ", folderPath)
	}
}

func TestEnterpriseOK(t *testing.T) {
	callback, result, adderr := createUserAndEnterprise()
	defer callback(t)
	if adderr != nil {
		t.Fatal(adderr)
	}
	// call the Enterprise method
	foundenterprise, finderr := resolvers.Enterprise(result.authctx, graphql.GetEnterpriseArgs{0})
	if finderr != nil {
		t.Fatal(finderr)
	}
	if result.enterprise.Name != foundenterprise.Name {
		t.Fatal("given and created enterprise are not equal", result.enterprise.Name, foundenterprise.Name)
	}
}
func TestDeleteEnterpriseOK(t *testing.T) {
	callback, result, adderr := createUserAndEnterprise()
	defer callback(t)
	if adderr != nil {
		t.Fatal(adderr)
	}
	// call the DeleteEnterprise method
	message, deleteerr := resolvers.DeleteEnterprise(result.authctx, graphql.DeleteEnterpriseArgs{MockedPasswordString, result.enterprise.ID})
	if deleteerr != nil {
		t.Error(deleteerr)
	}
	if message.Message != "enterprise successfully deleted" {
		t.Error("message is not as expected:", message.Message)
	}
}

func TestEnterpriseLogoOK(t *testing.T) {
	const FILENAME = "image.jpg"
	callback, result, addErr := createUserAndEnterprise()
	defer callback(t)
	if addErr != nil {
		t.Fatal(addErr)
	}
	response, logoErr := resolvers.EnterpriseLogo(result.authctx, graphql.EnterpriseLogoArgs{0, FILENAME})
	if logoErr != nil {
		t.Fatal(logoErr)
	}
	if len(strings.Split(response.Uploadtoken, ".")) != 3 {
		t.Fatal("upload token has an incorrect structure (does not have 3 parts): ", response.Uploadtoken)
	}
	// find the enterprise
	id, findErr := graphql.FindEnterpriseByIndex(result.userID, 0)
	if findErr != nil {
		t.Fatal(findErr)
	}
	enterprise, _ := graphql.FindEnterprise(*id)
	if filepath.Ext(enterprise.Logo) != filepath.Ext(FILENAME) {
		t.Fatal("file extensions do not match: ", enterprise.Logo, FILENAME)
	}
}

type cutils struct {
	authctx    context.Context
	userID     primitive.ObjectID
	enterprise *graphql.EnterpriseGQL
}

// create a user and an enterprise
func createUserAndEnterprise() (func(t *testing.T), *cutils, error) {
	dbctx, cancel := utils.CreateDBContext()

	// create a mock user
	userresult, usererr := database.UserCol.InsertOne(dbctx, *MockedAuthUser)
	if usererr != nil {
		return nil, nil, usererr
	}
	// mock the auth context
	authctx := context.WithValue(context.Background(), utils.User_id, userresult.InsertedID)
	// add an enterprise
	createdenterprise, adderr := resolvers.AddEnterprise(authctx, graphql.AddEnterpriseArgs{*MockedEnterprise})
	// a callback for clearing up the db
	var callback = func(t *testing.T) {
		defer cancel()
		_, edeleteErr := database.EnterpriseCol.DeleteOne(dbctx, bson.M{"name": MockedEnterprise.Name})
		if edeleteErr != nil {
			t.Fatal(edeleteErr)
		}
		_, udeleteErr := database.UserCol.DeleteOne(dbctx, bson.M{"_id": userresult.InsertedID})
		if udeleteErr != nil {
			t.Fatal(edeleteErr)
		}
	}
	u := &cutils{authctx, userresult.InsertedID.(primitive.ObjectID), createdenterprise}
	return callback, u, adderr
}
