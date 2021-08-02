package assets_test

import (
	"fmt"
	"os"
	"testing"

	"redinnlabs.com/redinn-core/assets"
)

func TestMain(m *testing.M) {
	// set up ASSETS_DIR - normally it should be set in the Setup func
	assets.ASSETS_DIR = os.Getenv("ASSETS_DIR")
	if len(assets.ASSETS_DIR) == 0 {
		CWD, cwderr := os.Getwd()
		if cwderr != nil {
			fmt.Println(cwderr)
		}
		assets.ASSETS_DIR = CWD + "/uploads/"
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
	code := m.Run()
	os.Exit(code)
}
