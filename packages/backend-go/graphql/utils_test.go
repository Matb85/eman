package graphql_test

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/auth"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/graphql"
)

func TestFindUser200(t *testing.T) {
	dbctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	model := auth.User{
		Email:     "user@example.com",
		Password:  "example-password",
		FirstName: "example",
		LastName:  "user",
	}
	result, inserterr := database.UserCol.InsertOne(dbctx, model)
	if inserterr != nil {
		t.Fatal(inserterr)
	}

	user, finderr := graphql.FindUser(result.InsertedID.(primitive.ObjectID))
	if finderr != nil {
		t.Fatal(finderr)
	}
	if user.Email != model.Email {
		t.Error("emails are not the same - user:", user.Email, "model:", model.Email)
	}
}
