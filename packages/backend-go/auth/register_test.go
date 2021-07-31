package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"redinnlabs.com/redinn-core/auth"
)

func TestRegister200(t *testing.T) {
	// setup payload
	payload, _ := json.Marshal(MockedUser)

	w := httptest.NewRecorder()
	auth.Register(w, httptest.NewRequest("POST", "/register", bytes.NewReader(payload)))
	res := w.Result()

	// read the response
	mes := &Message{}
	if jsonerr := json.NewDecoder(res.Body).Decode(&mes); jsonerr != nil {
		t.Fatal(jsonerr)
	}

	// log possible errors
	if res.StatusCode != 200 {
		t.Error("wrong status:", res.StatusCode, "instead of 200")
	}
	if mes.Message != "registration successful" {
		t.Error("wrong message:", mes.Message)
	}
}

func TestRegister400(t *testing.T) {
	// setup scenarios
	scenarios := []auth.User{
		// too short email - empty string
		{Password: "example-password", FirstName: "example", LastName: "user"},
		// too short password - empty string
		{Email: "example@redinnlabs.com", FirstName: "example", LastName: "user"},
		// too short first name - empty string
		{Email: "example@redinnlabs.com", Password: "example-password", LastName: "user"},
		// too short last name - empty string
		{Email: "example@redinnlabs.com", Password: "example-password", FirstName: "example"},
	}

	for i, scenario := range scenarios {
		w := httptest.NewRecorder()
		payload, _ := json.Marshal(scenario)

		auth.Register(w, httptest.NewRequest("POST", "/register", bytes.NewReader(payload)))
		res := w.Result()

		// read the response
		mes := &Message{}
		if jsonerr := json.NewDecoder(res.Body).Decode(&mes); jsonerr != nil {
			t.Fatal(jsonerr)
		}

		// log possible errors
		if res.StatusCode != 400 {
			t.Fatal("scenario", i, "wrong status:", res.StatusCode, "instead of 400")
		}
		if mes.Message != "provided wrong data" {
			t.Fatal("scenario", i, "wrong message:", mes.Message)
		}
	}

}
