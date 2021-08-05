package assets

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/utils"
)

func upload(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := utils.CreateDBContext()
	defer cancel()
	// auth - validate the upload token
	uploadClaims, autherr := VerifyUploadToken(r.Header.Get("Upload"))
	if autherr != nil {
		w.Header().Set("Content-Type", "application/json")
		utils.SendMessage(w, http.StatusUnauthorized, autherr.Error())
		return
	}
	// get user's folder path
	user := &database.User{}
	finderr := database.UserCol.FindOne(ctx, bson.M{"_id": r.Context().Value(utils.User_id)}).Decode(user)
	if finderr != nil {
		utils.SendMessage(w, http.StatusBadRequest, "could not find the user: "+finderr.Error())
	}
	// Parse request body as multipart form data with 32MB maxrmemory
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		utils.SendMessage(w, http.StatusBadRequest, "could not parse the body: "+err.Error())
	}
	// Get the file from Form
	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.SendMessage(w, http.StatusBadRequest, "could not read from the \"file\" key: "+err.Error())
		return
	}
	defer file.Close()
	// Create file locally
	dst, err := os.Create(ASSETS_DIR + user.Folder + uploadClaims.Hash + filepath.Ext(handler.Filename))
	if err != nil {
		utils.SendMessage(w, http.StatusInternalServerError, "could not create the file: "+err.Error())
		return
	}
	defer dst.Close()
	// Copy the uploaded file data to the newly created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		utils.SendMessage(w, http.StatusInternalServerError, "could not copy the file data: "+err.Error())
		return
	}

	utils.SendMessage(w, http.StatusOK, "successfully uploaded a file! ")
}
