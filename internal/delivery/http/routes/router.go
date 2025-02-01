package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ryvasa/go-super-farmer-auth-service/internal/delivery/http/handler"
)

type Router interface {
}

func NewRoutes(handlers *handler.Handlers) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	NewAuthRoutes(r, handlers.AuthHandler)

	return r
}
