package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"

	"redinnlabs.com/redinn-core/auth"
	"redinnlabs.com/redinn-core/database"
)

// struct for decoding responses
type message struct {
	Message string `json:"message"`
}

func TestMain(m *testing.M) {
	// setup
	database.Connect()
	// run tests
	code := m.Run()
	// shutdown
	os.Exit(code)
}

func TestRegister200(t *testing.T) {
	w := httptest.NewRecorder()

	// setup payload
	payload, _ := json.Marshal(&map[string]string{
		"email":     "example@redinnlabs.com",
		"password":  "example-password",
		"firstname": "example",
		"lastname":  "user",
	})

	auth.Register(w, httptest.NewRequest("GET", "/register", bytes.NewReader(payload)))
	res := w.Result()

	// read the response
	mes := &message{}
	jsonerr := json.NewDecoder(res.Body).Decode(&mes)

	// log possible errors
	if jsonerr != nil {
		t.Fatal(jsonerr)
	}
	if res.StatusCode != 200 {
		t.Error("wrong status:", res.StatusCode, "instead of 200")
	}
	if mes.Message != "registration successsful" {
		t.Error("wrong message:", mes.Message)
	}
}
