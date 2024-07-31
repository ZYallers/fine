package fhandler

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/gin-gonic/gin"
)

func WithSignHandler(handler gin.HandlerFunc) fapp.Option {
	return func(app *fapp.App) {
		app.Sign.SignHandler = handler
	}
}
