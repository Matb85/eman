package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Enterprises []int  `json:"enterprises"`
}

func Register(DB *mongo.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		user := User{Enterprises: []int{}}
		// convert json to struct
		err := json.NewDecoder(r.Body).Decode(&user)
		// some validation
		if err != nil || len(user.Password) < 8 || len(user.Email) < 3 || len(user.FirstName) < 3 || len(user.LastName) < 3 {
			SendResponse(w, http.StatusBadRequest, &map[string]string{
				"message": "provided wrong data",
			})
			return
		}
		// get the users collection
		COL := DB.Collection("users")
		// check if email is already used
		if count, _ := COL.CountDocuments(ctx, bson.M{"email": user.Email}); count > 0 {
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
}
