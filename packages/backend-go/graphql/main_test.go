package graphql_test

import (
	"log"
	"os"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"redinnlabs.com/redinn-core/assets"
	"redinnlabs.com/redinn-core/database"
)

func TestMain(m *testing.M) {
	// set up ASSETS_DIR - normally it should be set in the Setup func
	assets.ASSETS_DIR = os.Getenv("ASSETS_DIR")
	if len(assets.ASSETS_DIR) == 0 {
		log.Fatal("ASSETS_DIR not specified")
	}
	database.Connect()
	// hash the password of the MockAuthUser
	hash, _ := bcrypt.GenerateFromPassword([]byte(MockedAuthUser.Password), 10)
	MockedAuthUser.Password = string(hash)
	code := m.Run()

	os.Exit(code)
}
