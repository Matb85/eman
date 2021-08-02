package assets_test

import (
	"fmt"
	"testing"

	"redinnlabs.com/redinn-core/assets"
)

func TestCreateFolderOK(t *testing.T) {
	uuidv4, uuiderr := assets.GenUUIDv4()
	if uuiderr != nil {
		t.Error("unexpected error: ", uuiderr)
	}
	fmt.Println(uuidv4)
	path, createerr := assets.CreateFolder(true, uuidv4)
	if createerr != nil {
		t.Error("unexpected error: ", createerr, path)
	}
}
