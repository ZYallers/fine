package handler

import (
	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
)

func Stats(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Param("filepath") == "/ws" {
			statsviz.Ws(ctx.Writer, ctx.Request)
			return
		}
		statsviz.IndexAtRoot(root).ServeHTTP(ctx.Writer, ctx.Request)
	}
}
