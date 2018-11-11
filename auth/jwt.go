package auth

import (
	"time"

	"github.com/Asprilla24/vermouth/config"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Name string
	jwt.StandardClaims
}

func CreateJWTToken(username string) (string, error) {
	tokenCode := config.GetConfig().TokenCode

	claims := JWTClaims{
		username,
		jwt.StandardClaims{
			Id:        username,
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and send it as response
	t, err := rawToken.SignedString([]byte(tokenCode))
	if err != nil {
		return "", err
	}

	return t, nil
}
