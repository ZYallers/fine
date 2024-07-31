package fhandler

import (
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/frouter"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/handler"
)

func WithRouter(router frouter.Router) fapp.Option {
	return func(app *fapp.App) {
		app.Router = router
		handler.ParseRouter(app)
	}
}
