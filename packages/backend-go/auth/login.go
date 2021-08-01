package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := utils.CreateDBContext()
	defer cancel()
	user := User{Enterprises: []string{}}
	// convert json to struct
	err := json.NewDecoder(r.Body).Decode(&user)
	// some validation
	if err != nil || len(user.Password) < 8 || len(user.Email) < 6 {
		utils.SendResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "too short password or email",
		})
		return
	}
	// get the users collection
	COL := database.UserCol
	// fetch the user
	fetchuser := User{Enterprises: []string{}}
	fetchErr := COL.FindOne(ctx, bson.M{"email": user.Email}).Decode(&fetchuser)
	// compare passwords
	passworderr := bcrypt.CompareHashAndPassword([]byte(fetchuser.Password), []byte(user.Password))

	if fetchErr != nil || passworderr != nil {
		utils.SendResponse(w, http.StatusBadRequest, &map[string]string{
			"message": "incorrect password or email",
		})
		return
	}
	// create tokens
	access_token, e1 := createToken(fetchuser.Id, time.Minute*30)
	refresh_token, e2 := createToken(fetchuser.Id, time.Hour*24*30)
	if e1 == nil && e2 == nil {
		utils.SendResponse(w, http.StatusOK, &map[string]string{
			"message":       "Welcome to Redinn!",
			"access_token":  access_token,
			"refresh_token": refresh_token,
		})
	} else {
		utils.SendResponse(w, http.StatusInternalServerError, &map[string]string{
			"message": "internal error",
		})
	}
}

type customClaims struct {
	ID primitive.ObjectID `json:"id"`
	jwt.StandardClaims
}

var SECRET = "gdsgfdgdfhfghgfhldsjgdfk"

func createToken(id primitive.ObjectID, duration time.Duration) (string, error) {
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
