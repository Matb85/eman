package assets

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type uploadClaims struct {
	Path string `json:"path"`
	jwt.StandardClaims
}

var PHOTO_UPLOAD_DUR = time.Second * 30

var SECRET = "gdsgfdgdfhfghgfhldsjgdfk"

// path = /folder/file.ext - starts with a slash!
func CreateUploadToken(path string, duration time.Duration) (string, error) {
	claims := uploadClaims{
		Path: path,
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
