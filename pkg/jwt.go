package pkg

import (
	"os"
	"time"

	"github.com/ThailanTec/jwt/model"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

// var jwtSecret = []byte("28@319831%2n&31o!2%m3o2m3o2&%#$%21m30129324#!#@31nf98sehasnd=021m3021m312312-21m82")
var jwtSecret []byte

func GenerateToken(user model.User) (string, error) {

	if err := godotenv.Load(); err != nil {
		panic("Erro ao carregar o arquivo .env")
	}
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 120).Unix(), // Token expira em 5 dias
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
