package graphql_test

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/graphql"
)

func TestAddEnterpriseOK(t *testing.T) {
	callback, result, adderr := createUserAndEnterprise()
	defer callback(t)
	if adderr != nil {
		t.Fatal(adderr)
	}
	if result.enterprise.Name != MockedEnterprise.Name {
		t.Error("given and created enterprise are not equal", result.enterprise.Name, MockedEnterprise.Name)
	}
}

func TestEnterpriseOK(t *testing.T) {
	callback, result, adderr := createUserAndEnterprise()
	defer callback(t)
	if adderr != nil {
		t.Fatal(adderr)
	}
	// call the Enterprise method
	foundenterprise, finderr := result.resolvers.Enterprise(result.authctx, graphql.GetEnterpriseArgs{0})
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
	message, deleteerr := result.resolvers.DeleteEnterprise(result.authctx, graphql.DeleteEnterpriseArgs{MockedPasswordString, result.enterprise.ID})
	if deleteerr != nil {
		t.Error(deleteerr)
	}
	if message.Message != "enterprise successfully deleted" {
		t.Error("message is not as expected:", message.Message)
	}
}

type utils struct {
	authctx    context.Context
	userID     primitive.ObjectID
	enterprise *graphql.EnterpriseGQL
	resolvers  *graphql.EnterpriseResolvers
}

// create a user and an enterprise
func createUserAndEnterprise() (func(t *testing.T), *utils, error) {
	dbctx, cancel := graphql.CreateDBContext()

	resolvers := &graphql.EnterpriseResolvers{}
	// create a mock user
	userresult, usererr := database.UserCol.InsertOne(dbctx, *MockedAuthUser)
	if usererr != nil {
		return nil, nil, usererr
	}
	// mock the auth context
	authctx := context.WithValue(context.Background(), graphql.User_id, userresult.InsertedID)
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
	u := &utils{authctx, userresult.InsertedID.(primitive.ObjectID), createdenterprise, resolvers}
	return callback, u, adderr
}
