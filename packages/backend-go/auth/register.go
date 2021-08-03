package auth

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"redinnlabs.com/redinn-core/assets"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/utils"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	FirstName   string             `json:"firstname"`
	LastName    string             `json:"lastname"`
	Enterprises []string           `json:"enterprises"`
	Folder      string             `json:"folder"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := utils.CreateDBContext()
	defer cancel()
	user := User{Enterprises: []string{}}
	// convert json to struct
	err := json.NewDecoder(r.Body).Decode(&user)
	// some validation
	if err != nil || len(user.Password) < 8 || len(user.Email) < 6 || len(user.FirstName) < 3 || len(user.LastName) < 3 {
		utils.SendResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "provided wrong data",
		})
		return
	}
	// get the users collection
	COL := database.UserCol
	// check if email is already used
	if count, _ := COL.CountDocuments(ctx, bson.M{"email": user.Email}); count > 0 {
		utils.SendResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "email already in use",
		})
		return
	}
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if hashErr != nil {
		utils.SendResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "internal error: " + hashErr.Error(),
		})
		return
	}
	user.Password = string(hashedPassword)
	user.Id = primitive.NewObjectID()
	folder, foldererr := assets.CreateFolder(assets.USER, user.Id.Hex())
	if foldererr != nil {
		utils.SendResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "internal error: " + foldererr.Error(),
		})
		return
	}
	user.Folder = folder
	// finally, insert the user
	_, insertErr := COL.InsertOne(ctx, user)
	if insertErr != nil {
		utils.SendResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "internal error: " + insertErr.Error(),
		})
	} else {
		utils.SendResponse(w, http.StatusOK, &map[string]string{
			"message": "registration successful",
		})
	}
}
