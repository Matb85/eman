package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// get the current working directory
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
	// start a mongodb instance
	args := []string{
		"--bind_ip", "127.0.0.1",
		"--port", os.Getenv("DB_PORT"),
		"--dbpath", DBSTORAGE,
	}
	server := exec.Command(os.Getenv("TEST_DB_BIN_PATH"), args...)
	if mongoerr := server.Start(); mongoerr != nil {
		log.Fatal(mongoerr)
	}

	command := exec.Command("lerna", []string{"run", "test"}...)
	output, err := command.Output()
	fmt.Println("output:\n", string(output))
	if err != nil {
		fmt.Println("errors:\n", err.Error())
	}
	// shutdown
	server.Process.Kill()
}
