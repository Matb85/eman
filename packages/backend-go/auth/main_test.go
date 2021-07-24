package auth_test

import (
	"os"
	"testing"

	"redinnlabs.com/redinn-core/database"
)

// struct for decoding responses
type Message struct {
	Message string `json:"message"`
}

func TestMain(m *testing.M) {
	database.Connect()
	code := m.Run()
	os.Exit(code)
}
