package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"matb85.github.io/eman-core/auth"
	"matb85.github.io/eman-core/database"
	"matb85.github.io/eman-core/utils"
)

// a mocked user
var MockedUser = &database.User{
	Email:     "user@example.com",
	Password:  "example-password",
	FirstName: "example",
	LastName:  "user",
}

func TestLogin200(t *testing.T) {

	ctx, cancel := utils.CreateDBContext()
	defer cancel()

	// insert a mock user to the db
	mockUser := *MockedUser
	// hash the password
	hashedPassword, hasherr := bcrypt.GenerateFromPassword([]byte(mockUser.Password), 10)
	if hasherr != nil {
		t.Fatal(hasherr)
	}
	mockUser.Password = string(hashedPassword)
	// insert
	insertResult, inserterr := database.UserCol.InsertOne(ctx, mockUser)
	if inserterr != nil {
		t.Fatal(inserterr)
	}

	// test the controller
	w := httptest.NewRecorder()

	payload, _ := json.Marshal(MockedUser)
	auth.Login(w, httptest.NewRequest("POST", "/login", bytes.NewReader(payload)))
	res := w.Result()

	// read the response
	mes := &utils.Message{}
	if jsonerr := json.NewDecoder(res.Body).Decode(&mes); jsonerr != nil {
		t.Fatal(jsonerr)
	}
	// log possible errors
	if res.StatusCode != 200 {
		t.Error("wrong status:", res.StatusCode, "instead of 200")
	}
	if mes.Message != "Welcome to Eman!" {
		t.Error("wrong message:", mes.Message)
	}
	// delete the db entry
	if _, deleteerr := database.UserCol.DeleteOne(ctx, bson.M{"_id": insertResult.InsertedID}); deleteerr != nil {
		t.Fatal(deleteerr)
	}
}

func TestLogin400(t *testing.T) {
	// setup scenarios
	scenarios := []database.User{
		// too short email - empty string
		{Password: "example-password"},
		// too short password - empty string
		{Email: "example@matb85.github.io"},
		// incorrect email
		{Email: "wrong@matb85.github.io", Password: "example-password"},
		// incorrect password
		{Email: "example@matb85.github.io", Password: "wrong-password"},
	}
	// insert users to the db
	ctx, cancel := utils.CreateDBContext()
	defer cancel()

	for i, scenario := range scenarios {
		scenarioResult, _ := database.UserCol.InsertOne(ctx, scenario)

		w := httptest.NewRecorder()
		payload, _ := json.Marshal(scenario)

		auth.Login(w, httptest.NewRequest("POST", "/login", bytes.NewReader(payload)))
		res := w.Result()

		// delete the scenario in db
		database.UserCol.DeleteOne(ctx, bson.M{"_id": scenarioResult.InsertedID})

		// read the response
		mes := &utils.Message{}
		if jsonerr := json.NewDecoder(res.Body).Decode(&mes); jsonerr != nil {
			t.Fatal(jsonerr)
		}

		// log possible errors
		if res.StatusCode != 400 {
			t.Fatal("scenario", i, "wrong status:", res.StatusCode, "instead of 400")
		}
		if mes.Message != "incorrect password or email" && mes.Message != "too short password or email" {
			t.Fatal("scenario", i, "wrong message:", mes.Message)
		}
	}

}
