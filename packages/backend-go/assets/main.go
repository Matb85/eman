package assets

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/utils"
)

var ASSETS_DIR string

func Setup(router *mux.Router) {
	ASSETS_DIR = os.Getenv("ASSETS_DIR")
	if len(ASSETS_DIR) == 0 {
		CWD, cwderr := os.Getwd()
		if cwderr != nil {
			log.Fatal(cwderr)
		}
		ASSETS_DIR = CWD + "/uploads"
	}
	router.HandleFunc("/assets", handlePost).Methods("POST")
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/api/assets/", http.FileServer(http.Dir(ASSETS_DIR)))).Methods("GET")
}

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	FirstName   string             `json:"firstname"`
	LastName    string             `json:"lastname"`
	Enterprises []string           `json:"enterprises"`
	Folder      string             `json:"folder"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := utils.CreateDBContext()
	defer cancel()
	// auth - validate the upload token
	uploadClaims, autherr := VerifyUploadToken(r.Header.Get("Upload"))
	if autherr != nil {
		w.Header().Set("Content-Type", "application/json")
		utils.SendResponse(w, http.StatusUnauthorized, &map[string]string{
			"message": autherr.Error(),
		})
		return
	}
	// get user's folder path
	user := &User{}
	finderr := database.UserCol.FindOne(ctx, bson.M{"_id": r.Context().Value(utils.User_id)}).Decode(user)
	if finderr != nil {
		utils.SendResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "could not find the user: " + finderr.Error(),
		})
	}
	// Parse request body as multipart form data with 32MB maxrmemory
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "could not parse the body: " + err.Error(),
		})
	}
	// Get the file from Form
	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "could not read from the \"file\" key: " + err.Error(),
		})
		return
	}
	defer file.Close()
	// Create file locally
	dst, err := os.Create(ASSETS_DIR + user.Folder + uploadClaims.Hash + filepath.Ext(handler.Filename))
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "could not create the file: " + err.Error(),
		})
		return
	}
	defer dst.Close()
	// Copy the uploaded file data to the newly created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "could not copy the file data: " + err.Error(),
		})
		return
	}

	utils.SendResponse(w, http.StatusOK, &map[string]string{
		"message": "successfully uploaded a file! ",
	})
}
