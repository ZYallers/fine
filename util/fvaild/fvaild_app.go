package fvaild

import (
	"github.com/gin-gonic/gin/binding"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
)

func WithBinding(v binding.StructValidator) fapp.Option {
	return func(app *fapp.App) {
		binding.Validator = v
	}
}
