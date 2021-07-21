package graphql_test

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/graphql"
	"redinnlabs.com/redinn-core/tests"
)

func TestFindUser200(t *testing.T) {
	dbctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, inserterr := database.UserCol.InsertOne(dbctx, tests.Muser)
	if inserterr != nil {
		t.Fatal(inserterr)
	}

	user, finderr := graphql.FindUser(result.InsertedID.(primitive.ObjectID))
	if finderr != nil {
		t.Fatal(finderr)
	}
	if user.Email != tests.Muser.Email {
		t.Error("emails are not the same - user:", user.Email, "model:", tests.Muser.Email)
	}
}
