package fhandler

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/internal/handler"
)

func WithTracing() fapp.Option {
	return func(app *fapp.App) {
		handler.Tracing(app)
	}
}
