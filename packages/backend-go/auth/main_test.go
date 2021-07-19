package auth_test

import (
	"log"
	"os"
	"os/exec"
	"testing"

	"redinnlabs.com/redinn-core/database"
)

func TestMain(m *testing.M) {
	// setup
	// get current working directory
	CWD, cwderr := os.Getwd()
	if cwderr != nil {
		log.Fatal(cwderr)
	}
	DBSTORAGE := CWD + "/test_db"
	// create a temporary folder for mongodb
	// if already exists, clear content
	if _, staterr := os.Stat(DBSTORAGE); !os.IsNotExist(staterr) {
		if clearerr := os.RemoveAll(DBSTORAGE); clearerr != nil {
			log.Fatal(clearerr)
		}
	}

	if Mkdirerr := os.Mkdir(DBSTORAGE, 0755); Mkdirerr != nil {
		log.Fatal(Mkdirerr)
	}

	// start a mongo instance
	args := []string{
		"--bind_ip", "127.0.0.1",
		"--port", os.Getenv("DB_PORT"),
		"--dbpath", DBSTORAGE,
	}
	server := exec.Command(os.Getenv("TEST_DB_BIN_PATH"), args...)
	if mongoerr := server.Start(); mongoerr != nil {
		log.Fatal(mongoerr)
	}
	// connect to the instance
	database.Connect()
	// run tests
	code := m.Run()
	// shut down
	server.Process.Kill()
	os.Exit(code)
}
