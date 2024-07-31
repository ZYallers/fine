package fsession

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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
