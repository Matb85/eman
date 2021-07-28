package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"time"
)

var RESET = "\033[0m"
var RED = "\033[31m"
var GREEN = "\033[32m"

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

	serversShutdown := setupServers()
	command := exec.Command("lerna", []string{"run", "test"}...)
	output, err := command.Output()
	fmt.Println(GREEN+"output:\n"+RESET, string(output))
	if err != nil {
		fmt.Println(RED+"errors:\n"+RESET, err.Error())
	}
	// shutdown
	fmt.Println(GREEN + "shutting servers down..." + RESET)
	serversShutdown()
	server.Process.Kill()
}

func generateProxy(proxyurl string) *handler {
	url, err := url.Parse(proxyurl)
	if err != nil {
		panic(err)
	}

	director := func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
	}

	reverseProxy := &httputil.ReverseProxy{Director: director}
	return &handler{proxy: reverseProxy}
}

func setupServers() func() {
	fmt.Println(GREEN + "preparing setup frontend & backend..." + RESET)
	// setup backend
	dev := exec.Command("lerna", []string{"run", "dev"}...)

	// start servers in packages
	if err := dev.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(GREEN + "frontend & backend are running!" + RESET)

	fmt.Println(GREEN + "waiting 6 seconds for both ends to warm up..." + RESET)
	time.Sleep(time.Second * 6)

	// setup the proxies
	fmt.Println(GREEN + "setting up the reverse proxy..." + RESET)
	http.Handle("/app/", generateProxy("http://localhost:3000"))
	http.Handle("/api/", generateProxy("http://localhost:3001"))
	go http.ListenAndServe("localhost:3002", nil)
	fmt.Println(GREEN + "reverse proxy is running!" + RESET)

	return func() {
		dev.Process.Kill()
	}
}

type handler struct {
	proxy *httputil.ReverseProxy
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.proxy.ServeHTTP(w, r)
}
