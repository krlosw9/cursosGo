package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	user := "Carlos Waldo"
	email := "krlosw9@gmail.com"
	now := time.Now()
	iat := now.Unix()
	eat := now.AddDate(0, 0, 1).Unix()
	secret := "this-is-a-very-long-and-secure-secret"

	payload := jwt.MapClaims{
		"name":  user,
		"email": email,
		"iat":   iat,
		"exp":   eat,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tokenString)
}
