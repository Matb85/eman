package main

import (
	"fmt"
	"log"
	"os"
)

var RESET = "\033[0m"
var RED = "\033[31m"
var GREEN = "\033[32m"

func main() {
	programName := os.Args[1]
	switch programName {
	case "proxy":
		fmt.Println("creating a proxy...")
		setupProxy()
	case "unit":
		fmt.Println("preparing to run unit tests..")
		runTests(false)
	case "e2e":
		fmt.Println("preparing to run e2e tests..")
		runTests(true)
	default:
		log.Fatal("no arguments")
	}
}
