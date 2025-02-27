package fhandler

import (
	"github.com/ZYallers/fine/debug/fdebug"
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/internal/handler"
	"github.com/ZYallers/fine/internal/route"
	"github.com/gin-gonic/gin"
)

func WithNoRoute(h ...gin.HandlerFunc) fapp.Option {
	return func(app *fapp.App) {
		if len(h) > 0 {
			app.Server.HttpServer.Handler.(*gin.Engine).NoRoute(h[0])
		} else {
			app.Server.HttpServer.Handler.(*gin.Engine).NoRoute(handler.NoRoute())
		}
	}
}

func WithRootPath(handler gin.HandlerFunc) fapp.Option {
	return func(app *fapp.App) {
		app.Server.HttpServer.Handler.(*gin.Engine).Any("/", handler)
		route.Dumper.Append("/", "any", fdebug.FuncPath(handler))
	}
}

func WithPing(h ...gin.HandlerFunc) fapp.Option {
	return func(app *fapp.App) {
		if len(h) > 0 {
			app.Server.HttpServer.Handler.(*gin.Engine).GET("/ping", h[0])
			route.Dumper.Append("/ping", "get", fdebug.FuncPath(h[0]))
		} else {
			handlerFunc := handler.Pong()
			app.Server.HttpServer.Handler.(*gin.Engine).GET("/ping", handlerFunc)
			route.Dumper.Append("/ping", "get", fdebug.FuncPath(handlerFunc))
		}
	}
}

func WithHealth(h ...gin.HandlerFunc) fapp.Option {
	return func(app *fapp.App) {
		if len(h) > 0 {
			app.Server.HttpServer.Handler.(*gin.Engine).GET("/health", h[0])
			route.Dumper.Append("/health", "get", fdebug.FuncPath(h[0]))
		} else {
			handlerFunc := handler.Health()
			app.Server.HttpServer.Handler.(*gin.Engine).GET("/health", handler.Health())
			route.Dumper.Append("/health", "get", fdebug.FuncPath(handlerFunc))
		}
	}
}
