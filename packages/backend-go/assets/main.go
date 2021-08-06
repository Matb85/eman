package assets

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/assets", Upload).Methods("POST")
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/api/assets/", http.FileServer(http.Dir(ASSETS_DIR)))).Methods("GET")
}
