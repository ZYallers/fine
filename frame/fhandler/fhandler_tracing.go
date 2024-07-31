package fhandler

import (
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/handler"
)

func WithTracing() fapp.Option {
	return func(app *fapp.App) {
		handler.Tracing(app)
	}
}
