package authorization

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krlosw9/cursosGo/api-go/class-3/model"
)

func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    "krlosw9",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
