package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

func verify(tokenString string, secret interface{}) bool {
	token, err := jwt.Parse(tokenString, func(_ *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if token != nil && token.Valid && err == nil {
		return true
	}
	return false
}
