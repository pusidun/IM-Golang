package util

import (
	"im-golang/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secreate_key")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

func GenerateToken(user model.User) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tect",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
