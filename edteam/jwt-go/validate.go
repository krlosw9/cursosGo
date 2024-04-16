package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtybG9zdzlAZ21haWwuY29tIiwiZXhwIjoxNzEzMzkxMTAzLCJpYXQiOjE3MTMzMDQ3MDMsIm5hbWUiOiJDYXJsb3MgV2FsZG8ifQ.PcvtntJbWzN0U3djzXVo_f4FBSYBM6KUdowZqAjs37k"
	secret      = []byte("this-is-a-very-long-and-secure-secret")
)

func validateMethodAndGetSecret(token *jwt.Token) (any, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("Metodo de la firma no es valido")
	}

	return secret, nil
}

func main() {
	token, err := jwt.Parse(tokenString, validateMethodAndGetSecret)
	if err != nil {
		fmt.Printf("Token no valido, razon %v", err)
		return
	}

	userData, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Printf("no se pudo obtener la información del payload")
		return
	}

	fmt.Println("User", userData["name"])
	fmt.Println("email", userData["email"])
}
