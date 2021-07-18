package auth_test

import (
	"os"
	"testing"

	"redinnlabs.com/redinn-core/database"
)

func TestMain(m *testing.M) {
	// setup
	database.Connect()
	// run tests
	code := m.Run()
	// shutdown
	os.Exit(code)
}
