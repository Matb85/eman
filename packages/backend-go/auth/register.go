package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"redinnlabs.com/redinn-core/database"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	FirstName   string             `json:"firstname"`
	LastName    string             `json:"lastname"`
	Enterprises []string           `json:"enterprises"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := User{Enterprises: []string{}}
	// convert json to struct
	err := json.NewDecoder(r.Body).Decode(&user)
	// some validation
	if err != nil || len(user.Password) < 8 || len(user.Email) < 6 || len(user.FirstName) < 3 || len(user.LastName) < 3 {
		SendResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "provided wrong data",
		})
		return
	}
	// get the users collection
	COL := database.UserCol
	// check if email is already used
	if err := COL.FindOne(ctx, bson.M{"email": user.Email}).Decode(User{Enterprises: []string{}}); err == mongo.ErrNoDocuments {
		SendResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "email already in use",
		})
		return
	}
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if hashErr != nil {
		SendResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "internal error",
		})
		return
	}
	user.Password = string(hashedPassword)

	// finally, insert the user
	_, insertErr := COL.InsertOne(ctx, user)
	if insertErr != nil {
		SendResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "internal error",
		})
	} else {
		SendResponse(w, http.StatusOK, &map[string]string{
			"message": "registration successsful",
		})
	}
}
