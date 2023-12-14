package controller

import (
	"fmt"
	"net/http"

	"github.com/ThailanTec/jwt/model"
	"github.com/ThailanTec/jwt/pkg"
	"github.com/gin-gonic/gin"
)

// @Summary Login on app
// @ID Post
// @Produce json
// @Success 201
// @Failure 500
// @Param model.User
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Verifica as credenciais do usuário
	valid, err := AuthenticateUser(user.Username, user.Password)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Gera token JWT
	token, err := pkg.GenerateToken(user)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	user.Jwt = token

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, "})
}
