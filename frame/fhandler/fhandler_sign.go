package fhandler

import (
	"github.com/gin-gonic/gin"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
)

func WithSignHandler(handler gin.HandlerFunc) fapp.Option {
	return func(app *fapp.App) {
		app.Sign.SignHandler = handler
	}
}
