package fhandler

import (
	"github.com/gin-gonic/gin"
	"gitlab.sys.hxsapp.net/hxs/fine/debug/fdebug"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/util/fmode"

	"gitlab.sys.hxsapp.net/hxs/fine/internal/handler"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/route"
)

func WithSwagger(path ...string) fapp.Option {
	return func(app *fapp.App) {
		if app.Mode != fmode.DevMode {
			return
		}
		relativePath, docDir := "/swag/json", "doc"
		if len(path) > 0 {
			relativePath = path[0]
		}
		if len(path) > 1 {
			docDir = path[1]
		}
		handlerFunc := handler.Swagger(docDir)
		app.Server.HttpServer.Handler.(*gin.Engine).GET(relativePath)
		route.Dumper.Append(relativePath, "get", fdebug.FuncPath(handlerFunc))
	}
}
