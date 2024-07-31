package fsession

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/os/fctx"
	"github.com/gin-gonic/gin"
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
