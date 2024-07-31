package fapp

import (
	"net/http"

	"github.com/ZYallers/fine/frame/frouter"
	"github.com/ZYallers/fine/internal/route"
	"github.com/ZYallers/fine/os/fcfg"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type AppSession struct {
	Key          string
	KeyPrefix    string
	Expiration   int
	Client       func() *redis.Client
	LoginHandler gin.HandlerFunc
}

type AppVersion struct {
	Latest string
	Key    string
}

type AppLogger struct {
	LogTimeout  int
	SendTimeout int
}

type AppSign struct {
	Key         string
	TimeKey     string
	Secret      string
	Expiration  int
	SignHandler gin.HandlerFunc
}

type AppServer struct {
	Addr            string
	ReadTimeout     int
	WriteTimeout    int
	ShutdownTimeout int
	HttpServer      *http.Server
}

type App struct {
	Name    string
	Debug   bool
	Mode    string
	Session AppSession
	Version AppVersion
	Logger  AppLogger
	Sign    AppSign
	Server  AppServer
	Router  frouter.Router
}

var intApp *App

const (
	configPath = "manifest/config"
	configName = "config"
	configType = "yaml"
)

func (app *App) Run(options ...Option) {
	for _, option := range options {
		option(app)
	}
	app.info()
	if fcfg.GetBool("app.router.dumpRoutes") {
		route.Dumper.Dump()
	}
	app.startServer()
}
