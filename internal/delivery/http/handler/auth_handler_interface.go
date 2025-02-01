package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Home(ctx *gin.Context)
	GoogleLogin(ctx *gin.Context)
	Redirect(ctx *gin.Context)
}
