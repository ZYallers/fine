package fhandler

import (
	"github.com/gin-gonic/gin"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/handler"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/route"
)

func WithPProf() fapp.Option {
	return func(app *fapp.App) {
		handler.PProfRegister(app.Server.HttpServer.Handler.(*gin.Engine))
		route.Dumper.Append("/debug/pprof/*", "get", "")
	}
}
