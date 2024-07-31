package fhandler

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/internal/handler"
	"github.com/gin-gonic/gin"
)

func WithLogger() fapp.Option {
	return func(app *fapp.App) {
		app.Server.HttpServer.Handler.(*gin.Engine).Use(
			handler.LoggerWithZap(app.Logger.LogTimeout, app.Logger.SendTimeout),
		)
	}
}
