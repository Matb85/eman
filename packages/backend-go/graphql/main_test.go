package graphql_test

import (
	"os"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"redinnlabs.com/redinn-core/database"
)

func TestMain(m *testing.M) {
	// hash the password of the MockAuthUser
	database.Connect()
	hash, _ := bcrypt.GenerateFromPassword([]byte(MockedAuthUser.Password), 10)
	MockedAuthUser.Password = string(hash)
	code := m.Run()

	os.Exit(code)
}
