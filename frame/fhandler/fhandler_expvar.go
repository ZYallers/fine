package fhandler

import (
	"github.com/gin-gonic/gin"
	"gitlab.sys.hxsapp.net/hxs/fine/debug/fdebug"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/handler"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/route"
)

func WithExpVar() fapp.Option {
	return func(app *fapp.App) {
		handlerFunc := handler.ExpVar()
		app.Server.HttpServer.Handler.(*gin.Engine).GET("/expvar", handlerFunc)
		route.Dumper.Append("/expvar", "get", fdebug.FuncPath(handlerFunc))
	}
}
