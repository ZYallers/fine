package fhandler

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/frame/frouter"
	"github.com/ZYallers/fine/internal/handler"
)

func WithRouter(router frouter.Router) fapp.Option {
	return func(app *fapp.App) {
		app.Router = router
		handler.ParseRouter(app)
	}
}
