package fhandler

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/internal/handler"
	"github.com/ZYallers/fine/internal/route"
	"github.com/gin-gonic/gin"
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
