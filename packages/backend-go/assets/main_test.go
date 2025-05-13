package assets_test

import (
	"fmt"
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
	// if folder exits delete it
	if _, staterr := os.Stat(assets.ASSETS_DIR); !os.IsNotExist(staterr) {
		if err := os.RemoveAll(assets.ASSETS_DIR); err != nil {
			fmt.Println(err)
		}
	}
	// create the folder
	if err := os.Mkdir(assets.ASSETS_DIR, 0755); err != nil {
		fmt.Println(err)
	}
	database.Connect()
	code := m.Run()
	os.Exit(code)
}
