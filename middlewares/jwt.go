package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(id uint, name string) string {
	claims := &jwtCustomClaims{
		id,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t
}
