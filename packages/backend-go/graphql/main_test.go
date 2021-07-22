package graphql_test

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
	shutdown := tests.GlobalSetup()
	code := m.Run()
	// shut down
	shutdown()
	os.Exit(code)
}
