package server

import (
	"github.com/ThailanTec/jwt/controller"
	"github.com/ThailanTec/jwt/middleware"
	"github.com/gin-gonic/gin"
)

func Ginrouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", controller.LoginHandler)
	router.GET("/protected", middleware.AuthMiddleware(), controller.ProtectedHandler)
	router.Run(":8080")

	return router
}
