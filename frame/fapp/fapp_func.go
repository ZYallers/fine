package fapp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ZYallers/fine/frame/frouter"
	"github.com/ZYallers/fine/os/fcfg"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/gin-gonic/gin"
)

func Now() *App {
	return intApp
}

func ReadConfig() error {
	configFile := ffile.Join(configPath, configName+"."+configType)
	if err := fcfg.ReadConfig(configFile, false); err != nil {
		return err
	}
	if mode := fcfg.GetString("app.mode"); mode != "" {
		modeConfigFile := fmt.Sprintf("%s/%s.%s.%s", configPath, configName, mode, configType)
		_ = fcfg.ReadConfig(modeConfigFile, true)
	}
	return nil
}

func New(options ...Option) *App {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {}
	if err := ReadConfig(); err != nil {
		panic(err)
	}
	intApp = &App{
		Name:  fcfg.GetString("app.name"),
		Debug: fcfg.GetBool("app.debug"),
		Mode:  fcfg.GetString("app.mode"),
		Session: AppSession{
			Key:        fcfg.GetString("app.session.key"),
			KeyPrefix:  fcfg.GetString("app.session.keyPrefix"),
			Expiration: fcfg.GetInt("app.session.expiration"),
		},
		Version: AppVersion{
			Latest: fcfg.GetString("app.version.latest"),
			Key:    fcfg.GetString("app.version.key"),
		},
		Logger: AppLogger{
			LogTimeout:  fcfg.GetInt("app.logger.logTimeout"),
			SendTimeout: fcfg.GetInt("app.logger.sendTimeout"),
		},
		Sign: AppSign{
			Key:        fcfg.GetString("app.sign.key"),
			TimeKey:    fcfg.GetString("app.sign.timeKey"),
			Secret:     fcfg.GetString("app.sign.secret"),
			Expiration: fcfg.GetInt("app.sign.expiration"),
		},
		Server: AppServer{
			Addr:            fcfg.GetString("app.server.addr"),
			ReadTimeout:     fcfg.GetInt("app.server.readTimeout"),
			WriteTimeout:    fcfg.GetInt("app.server.writeTimeout"),
			ShutdownTimeout: fcfg.GetInt("app.server.shutdownTimeout"),
		},
		Router: frouter.Router{},
	}
	for _, option := range options {
		option(intApp)
	}
	intApp.Server.HttpServer = &http.Server{
		Handler:      gin.New(),
		Addr:         intApp.Server.Addr,
		ReadTimeout:  time.Duration(intApp.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(intApp.Server.WriteTimeout) * time.Second,
	}
	return intApp
}
