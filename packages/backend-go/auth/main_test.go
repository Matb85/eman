package auth_test

import (
	"os"
	"testing"

	"redinnlabs.com/redinn-core/tests"
)

// struct for decoding responses
type Message struct {
	Message string `json:"message"`
}

func TestMain(m *testing.M) {
	server := tests.GlobalSetup()
	code := m.Run()
	// shut down
	server.Process.Kill()
	os.Exit(code)
}
