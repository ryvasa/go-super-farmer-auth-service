package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	GoogleLogin(ctx *gin.Context)
}
