package lib

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("demojwt")

type jwtCustomClaims struct {
	ID       uint   `json: id`
	Username string `json: username`
	jwt.StandardClaims
}

func GenerateJWT(ID uint, Username string) string {
	t := time.Now()
	claim := &jwtCustomClaims{
		ID:       ID,
		Username: Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t.Add(48 * time.Hour).Unix(),
			Subject:   "access_token",
			IssuedAt:  t.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}
