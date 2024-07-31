package fhandler

import (
	"github.com/ZYallers/fine/debug/fdebug"
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/internal/handler"
	"github.com/ZYallers/fine/internal/route"
	"github.com/gin-gonic/gin"
)

func WithPrometheus() fapp.Option {
	return func(app *fapp.App) {
		handlerFunc := handler.Prometheus(app.Name)
		app.Server.HttpServer.Handler.(*gin.Engine).GET("/metrics", handlerFunc)
		route.Dumper.Append("/metrics", "get", fdebug.FuncPath(handlerFunc))
	}
}
