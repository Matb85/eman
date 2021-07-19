package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"redinnlabs.com/redinn-core/auth"
)

// struct for decoding responses
type message struct {
	Message string `json:"message"`
}

func TestRegister200(t *testing.T) {
	w := httptest.NewRecorder()

	// setup payload
	payload, _ := json.Marshal(&auth.User{
		Email:     "example@redinnlabs.com",
		Password:  "example-password",
		FirstName: "example",
		LastName:  "user",
	})

	auth.Register(w, httptest.NewRequest("POST", "/register", bytes.NewReader(payload)))
	res := w.Result()

	// read the response
	mes := &message{}
	if jsonerr := json.NewDecoder(res.Body).Decode(&mes); jsonerr != nil {
		t.Fatal(jsonerr)
	}

	// log possible errors
	if res.StatusCode != 200 {
		t.Error("wrong status:", res.StatusCode, "instead of 200")
	}
	if mes.Message != "registration successsful" {
		t.Error("wrong message:", mes.Message)
	}
}
