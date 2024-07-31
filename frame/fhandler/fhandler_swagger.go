package fhandler

import (
	"github.com/ZYallers/fine/debug/fdebug"
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/internal/handler"
	"github.com/ZYallers/fine/internal/route"
	"github.com/ZYallers/fine/util/fmode"
	"github.com/gin-gonic/gin"
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
