package fsession

import (
	"github.com/gin-gonic/gin"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/os/fctx"
)

func Data(ctx *gin.Context, app *fapp.App) string {
	if app == nil {
		app = fapp.Now()
	}
	if app == nil {
		return ""
	}
	key, prefix, clientFunc := app.Session.Key, app.Session.KeyPrefix, app.Session.Client
	if clientFunc == nil {
		return ""
	}
	client := clientFunc()
	if client == nil {
		return ""
	}
	token := fctx.QueryPostForm(ctx, key)
	if token == "" {
		return ""
	}
	val := client.Get(prefix + token).Val()
	return val
}
