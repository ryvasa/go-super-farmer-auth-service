package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ryvasa/go-super-farmer-auth-service/internal/delivery/http/handler"
)

type AuthRoutes struct {
	router  *gin.Engine
	handler handler.AuthHandler
}

func NewAuthRoutes(router *gin.Engine, handler handler.AuthHandler) {
	router.GET("/login", handler.GoogleLogin)
}
