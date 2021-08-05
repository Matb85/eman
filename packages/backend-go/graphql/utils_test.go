package graphql_test

import (
	"testing"

	gql "github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/graphql"
	"redinnlabs.com/redinn-core/utils"
)

var MockedUser = &graphql.UserGQL{
	Email:       "user@example.com",
	Firstname:   "example",
	Lastname:    "user",
	Enterprises: &[]gql.ID{},
}

const MockedPasswordString = "example-password"

var MockedAuthUser = &database.User{
	Email:       "user2@example.com",
	Password:    MockedPasswordString,
	FirstName:   "example",
	LastName:    "user",
	Enterprises: []string{},
}
var MockedEnterprise = &graphql.EnterpriseGQL{
	Name: "Example Enterprise",
}

func TestFindUserOK(t *testing.T) {
	dbctx, cancel := utils.CreateDBContext()
	defer cancel()

	result, inserterr := database.UserCol.InsertOne(dbctx, MockedUser)
	if inserterr != nil {
		t.Fatal(inserterr)
	}

	user, finderr := graphql.FindUser(result.InsertedID.(primitive.ObjectID))
	if finderr != nil {
		t.Fatal(finderr)
	}
	if user.Email != MockedUser.Email {
		t.Error("emails are not the same - user:", user.Email, "model:", MockedUser.Email)
	}
}
func TestFindEnterpriseOK(t *testing.T) {
	dbctx, cancel := utils.CreateDBContext()
	defer cancel()

	Menterprise := graphql.EnterpriseGQL{}
	result, inserterr := database.EnterpriseCol.InsertOne(dbctx, MockedUser)
	if inserterr != nil {
		t.Fatal(inserterr)
	}

	enterprise, finderr := graphql.FindEnterprise(result.InsertedID.(primitive.ObjectID))
	if finderr != nil {
		t.Fatal(finderr)
	}
	if enterprise.Name != Menterprise.Name {
		t.Error("emails are not the same - user:", enterprise.Name, "model:", Menterprise.Name)
	}
}
