package auth

import (
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Asprilla24/vermouth/config"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Name    string
	IsAdmin bool
	jwt.StandardClaims
}

func CreateJWTToken(username string, isAdmin bool) (string, error) {
	tokenCode := config.GetConfig().TokenCode

	claims := &JWTClaims{
		username,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	t, err := rawToken.SignedString([]byte(tokenCode))
	if err != nil {
		return "", err
	}

	return t, nil
}

func AuthenticateJWT() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &JWTClaims{},
		SigningKey: []byte(config.GetConfig().TokenCode),
	}
	return middleware.JWTWithConfig(config)
}
