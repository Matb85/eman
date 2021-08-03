package auth_test

import (
	"log"
	"os"
	"testing"

	"redinnlabs.com/redinn-core/assets"
	"redinnlabs.com/redinn-core/database"
)

// struct for decoding responses
type Message struct {
	Message string `json:"message"`
}

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
