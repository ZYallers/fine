package fvaild

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/gin-gonic/gin/binding"
)

func WithBinding(v binding.StructValidator) fapp.Option {
	return func(app *fapp.App) {
		binding.Validator = v
	}
}
