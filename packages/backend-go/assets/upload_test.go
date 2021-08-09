package assets_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/assets"
	"redinnlabs.com/redinn-core/auth"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/utils"
)

// test on a user
func TestUpload200(t *testing.T) {
	ctx, cancel := utils.CreateDBContext()
	defer cancel()

	MockedUser := database.User{
		Email:       "user2@example.com",
		Password:    "example-password",
		FirstName:   "example",
		LastName:    "user",
		Enterprises: []string{},
	}
	MockedUser.Id = primitive.NewObjectID()
	// create a folder for the user
	folder, folderErr := assets.CreateFolder(assets.USER, MockedUser.Id.Hex())
	if folderErr != nil {
		t.Fatal(folderErr)
	}
	MockedUser.Folder = folder
	// insert a user to the db
	_, insertErr := database.UserCol.InsertOne(ctx, MockedUser)
	if insertErr != nil {
		t.Fatal(insertErr)
	}
	// generate a hash for the file
	hash, uuiderr := assets.GenUUIDv4()
	if uuiderr != nil {
		t.Fatal(uuiderr)
	}
	// get an auth token
	authToken, authErr := auth.CreateToken(MockedUser.Id, time.Minute)
	if authErr != nil {
		t.Fatal(authErr)
	}
	// get an upload token
	uploadToken, uploadTokenErr := assets.CreateUploadToken(folder+"/"+hash+".jpg", assets.PHOTO_UPLOAD_DUR)
	if uploadTokenErr != nil {
		t.Fatal(uploadTokenErr)
	}

	// attach a file
	// New multipart writer
	file, openErr := os.Open(os.Getenv("BASE_DIR") + "/assets/image.jpg")
	if openErr != nil {
		t.Fatal(openErr)
	}
	fileContents, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		t.Fatal(readErr)
	}
	fi, statErr := file.Stat()
	if statErr != nil {
		t.Fatal(statErr)
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, formErr := writer.CreateFormFile("file", fi.Name())
	if formErr != nil {
		t.Fatal(formErr)
	}
	part.Write(fileContents)
	// close
	closeErr := writer.Close()
	if closeErr != nil {
		t.Fatal(closeErr)
	}

	// setup the WITH CONTEXT request & the recorder
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/upload", bytes.NewReader(body.Bytes())).WithContext(context.WithValue(context.Background(), utils.User_id, MockedUser.Id))
	// set headers
	req.Header.Set("Authorization", authToken)
	req.Header.Set("Upload", uploadToken)
	// Request Content-Type with boundary parameter.
	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary())
	req.Header.Set("Content-Type", contentType)
	// finally, test the controller
	assets.Upload(w, req)
	res := w.Result()

	// read the response
	mes := &utils.Message{}
	if jsonerr := json.NewDecoder(res.Body).Decode(&mes); jsonerr != nil {
		t.Fatal(jsonerr)
	}

	if mes.Message != "successfully uploaded a file! " {
		t.Fatal("wrong message: ", mes.Message)
	}
}
