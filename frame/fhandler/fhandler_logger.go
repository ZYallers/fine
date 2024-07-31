package fhandler

import (
	"github.com/gin-gonic/gin"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/handler"
)

func WithLogger() fapp.Option {
	return func(app *fapp.App) {
		app.Server.HttpServer.Handler.(*gin.Engine).Use(
			handler.LoggerWithZap(app.Logger.LogTimeout, app.Logger.SendTimeout),
		)
	}
}
