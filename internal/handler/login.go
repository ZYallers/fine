package handler

import (
	"net/http"

	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/os/fctx"
	"github.com/ZYallers/fine/os/fsession"
	"github.com/gin-gonic/gin"
)

func loginCheck(app *fapp.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if data := fsession.Data(ctx, app); data != "" {
			ctx.Set(fctx.SessionDataKey, data)
		} else {
			fctx.AbortJson(ctx, http.StatusUnauthorized, "please login first")
		}
	}
}
