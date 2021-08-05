package assets

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	router.HandleFunc("/assets", upload).Methods("POST")
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
