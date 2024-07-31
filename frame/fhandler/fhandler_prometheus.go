package fhandler

import (
	"github.com/gin-gonic/gin"
	"gitlab.sys.hxsapp.net/hxs/fine/debug/fdebug"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/handler"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/route"
)

func WithPrometheus() fapp.Option {
	return func(app *fapp.App) {
		handlerFunc := handler.Prometheus(app.Name)
		app.Server.HttpServer.Handler.(*gin.Engine).GET("/metrics", handlerFunc)
		route.Dumper.Append("/metrics", "get", fdebug.FuncPath(handlerFunc))
	}
}
