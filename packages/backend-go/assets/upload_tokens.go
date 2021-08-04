package assets

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type uploadClaims struct {
	Hash string `json:"hash"`
	jwt.StandardClaims
}

var PHOTO_UPLOAD_DUR = time.Second * 30

var SECRET = "gdsgfdgdfhfghgfhldsjgdfk"

func CreateUploadToken(hash string, duration time.Duration) (string, error) {
	claims := uploadClaims{
		Hash: hash,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			Subject:   "1",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET))
}

func VerifyUploadToken(tokenstring string) (*uploadClaims, error) {
	claims := &uploadClaims{}
	_, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	return claims, err
}
