package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Login(DB *mongo.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		user := User{Enterprises: []int{}}
		// convert json to struct
		err := json.NewDecoder(r.Body).Decode(&user)
		// some validation
		if err != nil || len(user.Password) < 8 || len(user.Email) < 3 {
			SendResponse(w, http.StatusBadRequest, &map[string]string{
				"message": "provided wrong data",
			})
			return
		}
		// get the users collection
		COL := DB.Collection("users")
		// fetch the user
		fetchuser := User{Enterprises: []int{}}
		fetchErr := COL.FindOne(ctx, bson.M{"email": user.Email}).Decode(&fetchuser)
		if fetchErr != nil {
			SendResponse(w, http.StatusBadRequest, &map[string]string{
				"message": "provided wrong data1",
			})
			return
		}
		// compare passwords
		if bcrypt.CompareHashAndPassword([]byte(fetchuser.Password), []byte(user.Password)) != nil {
			SendResponse(w, http.StatusBadRequest, &map[string]string{
				"message": "provided wrong data2",
			})
			return
		}
		access_token, e1 := createToken(fetchuser.Id.Hex(), time.Minute*15)
		refresh_token, e2 := createToken(fetchuser.Id.Hex(), time.Hour*24*30)
		if e1 == nil && e2 == nil {
			SendResponse(w, http.StatusOK, &map[string]string{
				"message":       "Welcome to Redinn!",
				"access_token":  access_token,
				"refresh_token": refresh_token,
			})
		} else {
			SendResponse(w, http.StatusInternalServerError, &map[string]string{
				"message": "internal error",
			})
		}
	}
}

type customClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var SECRET = "gdsgfdgdfhfghgfhldsjgdfk"

func createToken(id string, duration time.Duration) (string, error) {
	claims := customClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			Subject:   "1",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET))
}

func VerifyToken(tokenstring string) (*customClaims, error) {
	claims := &customClaims{}
	_, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	return claims, err

}
