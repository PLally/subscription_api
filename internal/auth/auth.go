package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func NewToken(secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"super":      true,
		"identifier": "433348813462175745",
		"type":       "internal",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	fmt.Println(tokenString, err)
	return ""
}
