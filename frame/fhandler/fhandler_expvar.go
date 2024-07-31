package fhandler

import (
	"github.com/ZYallers/fine/debug/fdebug"
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/internal/handler"
	"github.com/ZYallers/fine/internal/route"
	"github.com/gin-gonic/gin"
)

func WithExpVar() fapp.Option {
	return func(app *fapp.App) {
		handlerFunc := handler.ExpVar()
		app.Server.HttpServer.Handler.(*gin.Engine).GET("/expvar", handlerFunc)
		route.Dumper.Append("/expvar", "get", fdebug.FuncPath(handlerFunc))
	}
}
