package fhandler

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/internal/handler"
	"github.com/ZYallers/fine/internal/route"
	"github.com/gin-gonic/gin"
)

func WithPProf() fapp.Option {
	return func(app *fapp.App) {
		handler.PProfRegister(app.Server.HttpServer.Handler.(*gin.Engine))
		route.Dumper.Append("/debug/pprof/*", "get", "")
	}
}
