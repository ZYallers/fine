package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Abort()
		ctx.Header("Content-Type", "application/json")
		ctx.Status(http.StatusOK)
		_, _ = ctx.Writer.WriteString(`{"code":404,"msg":"page not found"}`)
	}
}

func RootPath() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	}
}

func Health() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	}
}

func Pong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	}
}
