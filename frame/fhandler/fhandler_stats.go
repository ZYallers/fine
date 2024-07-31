package fhandler

import (
	"github.com/gin-gonic/gin"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/handler"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/route"
)

func WithStats(root string, acts gin.Accounts) fapp.Option {
	return func(app *fapp.App) {
		if root == "" {
			root = "/stats"
		}
		engine := app.Server.HttpServer.Handler.(*gin.Engine)
		if acts == nil {
			engine.GET(root+"/*filepath", handler.Stats(root))
		} else {
			engine.Group(root, gin.BasicAuth(acts)).GET("/*filepath", handler.Stats(root))
		}
		route.Dumper.Append(root+"/*", "get", "")
	}
}
