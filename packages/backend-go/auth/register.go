package auth

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"matb85.github.io/eman-core/assets"
	"matb85.github.io/eman-core/database"
	"matb85.github.io/eman-core/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := utils.CreateDBContext()
	defer cancel()
	user := database.User{Enterprises: []string{}}
	// convert json to struct
	err := json.NewDecoder(r.Body).Decode(&user)
	// some validation
	if err != nil || len(user.Password) < 8 || len(user.Email) < 6 || len(user.FirstName) < 3 || len(user.LastName) < 3 {
		utils.SendMessage(w, http.StatusBadRequest, "provided wrong data")
		return
	}
	// get the users collection
	COL := database.UserCol
	// check if email is already used
	if count, _ := COL.CountDocuments(ctx, bson.M{"email": user.Email}); count > 0 {
		utils.SendMessage(w, http.StatusBadRequest, "email already in use")
		return
	}
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if hashErr != nil {
		utils.SendMessage(w, http.StatusInternalServerError, "internal error: "+hashErr.Error())
		return
	}
	user.Password = string(hashedPassword)
	user.Id = primitive.NewObjectID()
	folder, foldererr := assets.CreateFolder(assets.USER, user.Id.Hex())
	if foldererr != nil {
		utils.SendMessage(w, http.StatusInternalServerError, "internal error: "+foldererr.Error())
		return
	}
	user.Folder = folder
	// finally, insert the user
	_, insertErr := COL.InsertOne(ctx, user)
	if insertErr != nil {
		utils.SendMessage(w, http.StatusInternalServerError, "internal error: "+insertErr.Error())
	} else {
		utils.SendMessage(w, http.StatusOK, "registration successful")
	}
}
