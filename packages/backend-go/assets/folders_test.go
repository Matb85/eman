package assets_test

import (
	"testing"

	"redinnlabs.com/redinn-core/assets"
)

func TestGenUUIDv4(t *testing.T) {
	uuidv4, uuiderr := assets.GenUUIDv4()
	if uuiderr != nil {
		t.Fatal(uuiderr, uuidv4)
	}
}
func TestGenUUIDv5(t *testing.T) {
	uuidv4, _ := assets.GenUUIDv4()
	uuidv5, uuiderr := assets.GenUUIDv5(uuidv4, uuidv4)
	if uuiderr != nil {
		t.Fatal(uuiderr, uuidv5)
	}
}

func TestCreateFolderOK(t *testing.T) {
	uuidv4, uuiderr := assets.GenUUIDv4()
	if uuiderr != nil {
		t.Fatal("unexpected error: ", uuiderr)
	}
	if path, createerr := assets.CreateFolder(assets.ENTERPRISE, uuidv4); createerr != nil {
		t.Fatal("unexpected error when creating an enterprise folder: ", createerr, path)
	}
	if path, createerr := assets.CreateFolder(assets.USER, uuidv4); createerr != nil {
		t.Fatal("unexpected error when creating an user folder: ", createerr, path)
	}
}
func TestRemoveFolderOK(t *testing.T) {
	uuidv4, uuiderr := assets.GenUUIDv4()
	if uuiderr != nil {
		t.Fatal("unexpected error: ", uuiderr)
	}
	path, createerr := assets.CreateFolder(assets.ENTERPRISE, uuidv4)
	if createerr != nil {
		t.Fatal("unexpected error when creating an enterprise folder: ", createerr, path)
	}
	// remove the folder
	if err := assets.RemoveFolder(path); err != nil {
		t.Fatal(err)
	}
}
