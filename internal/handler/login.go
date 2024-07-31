package handler

import (
	"gitlab.sys.hxsapp.net/hxs/fine/os/fctx"
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/os/fsession"
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
