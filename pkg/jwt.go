package pkg

import (
	"time"

	"github.com/ThailanTec/jwt/model"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("28@319831%2n&31o!2%m3o2m3o2&%#$%21m30129324#!#@31nf98sehasnd=021m3021m312312-21m82")

func GenerateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expira em 1 hora
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
