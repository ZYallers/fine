package fsession

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
)

func WithClient(client func() *redis.Client) fapp.Option {
	return func(app *fapp.App) {
		app.Session.Client = client
	}
}

func WithLoginHandler(handler gin.HandlerFunc) fapp.Option {
	return func(app *fapp.App) {
		app.Session.LoginHandler = handler
	}
}
