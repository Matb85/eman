package assets

import (
	"os"
	"os/exec"
	"strings"
)

const ENTERPRISE = true
const USER = false

func GenUUIDv4() (string, error) {
	out, err := exec.Command("uuidgen").CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

func GenUUIDv5(name, namespace string) (string, error) {
	out, err := exec.Command("uuidgen", "--sha1", "--name", name, "--namespace", namespace).CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

// category - ENTERPRISE or USER; id - uuidv4 or mongodb document id
func CreateFolder(category bool, id string) (string, error) {
	var path string
	// generate uuids
	uuidv4, errv4 := GenUUIDv4()
	if errv4 != nil {
		return "uuidv4 error" + uuidv4, errv4
	}
	uuidv5, errv5 := GenUUIDv5(id, uuidv4)
	if errv5 != nil {
		return "uuidv5 error" + uuidv5, errv5
	}
	// add a prefix depending on the given category
	if category {
		path = "/e-" + uuidv5
	} else {
		path = "/u-" + uuidv5
	}
	// create the folder
	if err := os.Mkdir(ASSETS_DIR+path, 0755); err != nil {
		return "", err
	}
	return path, nil
}

func RemoveFolder(path string) error {
	return os.RemoveAll(ASSETS_DIR + path)
}
