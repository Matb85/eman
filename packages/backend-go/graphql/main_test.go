package graphql_test

import (
	"os"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"redinnlabs.com/redinn-core/tests"
)

// struct for decoding responses
type Message struct {
	Message string `json:"message"`
}

func TestMain(m *testing.M) {
	shutdown := tests.GlobalSetup()
	// hash the password of the MockAuthUser
	hash, _ := bcrypt.GenerateFromPassword([]byte(MockedAuthUser.Password), 10)
	MockedAuthUser.Password = string(hash)
	code := m.Run()
	// shut down
	shutdown()
	os.Exit(code)
}
