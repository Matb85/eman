package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func runTests(e2e bool) {
	testType := "unit"

	callbackDB := setupDB()
	defer callbackDB()
	// if running e2e test run servers on localhost
	if e2e {
		testType = "e2e"
		serversShutdown := setupServers()
		defer serversShutdown()
	}

	command := exec.Command("lerna", []string{"run", "test:" + testType}...)

	fmt.Println(GREEN + "output:\n" + RESET)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Println(RED+"errors:\n"+RESET, err.Error())
	}
	// shutdown
	fmt.Println(GREEN + "shutting servers down..." + RESET)
}

func setupServers() func() {
	// start servers in packages
	start := exec.Command("lerna", []string{"run", "start"}...)
	start.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	start.Stdout = os.Stdout
	start.Stderr = os.Stderr
	if err := start.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(GREEN + "frontend & backend are running!" + RESET)
	fmt.Println(GREEN + "waiting 3 seconds for both ends to warm up..." + RESET)
	time.Sleep(time.Second * 3)

	// setup the proxies
	fmt.Println(GREEN + "setting up the reverse proxy..." + RESET)
	http.Handle("/app/", generateProxy("http://localhost:3000"))
	http.Handle("/api/", generateProxy("http://localhost:3001"))
	go http.ListenAndServe("localhost:3002", nil)
	fmt.Println(GREEN + "reverse proxy is running!" + RESET)

	return func() {
		pgid, err := syscall.Getpgid(start.Process.Pid)
		if err == nil {
			syscall.Kill(-pgid, 15) // note the minus sign
		}
		start.Wait()
	}
}

func setupDB() func() {
	// get the current working directory
	CWD, cwderr := os.Getwd()
	if cwderr != nil {
		log.Fatal(cwderr)
	}
	DBSTORAGE := CWD + "/test_db"
	ASSETS_DIR := CWD + "/uploads"
	// get the BASE_DIR
	setErr := os.Setenv("BASE_DIR", CWD)
	if setErr != nil {
		log.Fatal(setErr)
	}
	setErr2 := os.Setenv("ASSETS_DIR", ASSETS_DIR)
	if setErr2 != nil {
		log.Fatal(setErr2)
	}
	// create a temporary folders for mongodb & assets
	for _, value := range []string{DBSTORAGE, ASSETS_DIR} {
		// if already exists, clear content
		if _, staterr := os.Stat(value); !os.IsNotExist(staterr) {
			if clearerr := os.RemoveAll(value); clearerr != nil {
				log.Fatal(clearerr)
			}
		}
		if Mkdirerr := os.Mkdir(value, 0755); Mkdirerr != nil {
			log.Fatal(Mkdirerr)
		}
	}
	// start a mongodb instance
	fmt.Println(GREEN + "preparing to run mongodb..." + RESET)

	args := []string{
		"--bind_ip", "127.0.0.1",
		"--port", os.Getenv("DB_PORT"),
		"--dbpath", DBSTORAGE,
	}
	server := exec.Command(os.Getenv("TEST_DB_BIN_PATH"), args...)
	if mongoerr := server.Start(); mongoerr != nil {
		log.Fatal(mongoerr)
	}
	fmt.Println(GREEN + "mongodb is running!" + RESET)
	return func() {
		server.Process.Kill()
		server.Process.Wait()
		fmt.Println("removing storage")
		os.RemoveAll(DBSTORAGE)
		os.RemoveAll(ASSETS_DIR)
	}
}
