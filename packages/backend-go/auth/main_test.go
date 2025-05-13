package auth_test

import (
	"log"
	"os"
	"testing"

	"matb85.github.io/eman-core/assets"
	"matb85.github.io/eman-core/database"
)

func TestMain(m *testing.M) {
	// set up ASSETS_DIR - normally it should be set in the Setup func
	assets.ASSETS_DIR = os.Getenv("ASSETS_DIR")
	if len(assets.ASSETS_DIR) == 0 {
		log.Fatal("ASSETS_DIR not specified")
	}
	database.Connect()
	code := m.Run()
	os.Exit(code)
}
