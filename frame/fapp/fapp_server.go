package fapp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ZYallers/fine/frame/fmsg"
	"github.com/ZYallers/fine/os/flog"
)

func record(msg string, args ...string) {
	level := "INFO"
	if len(args) > 0 && args[0] != "" {
		level = strings.ToUpper(args[0])
	}
	s := fmt.Sprintf("[%d] %s", syscall.Getpid(), msg)
	switch level {
	case "ERROR":
		flog.Use().Error(s)
	default:
		flog.Use().Info(s)
	}
	fmsg.Sender().Simple(s, true)
}

func (app *App) startServer() {
	go func() {
		record(fmt.Sprintf("http server started listening on %s", app.Server.Addr))
		if err := app.Server.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			record(fmt.Sprintf("http server start listening error: %v", err), "ERROR")
			os.Exit(1)
		}
	}()

	quitChan := make(chan os.Signal, 1)
	// SIGTERM 结束程序(kill pid)(可以被捕获、阻塞或忽略)
	// SIGHUP 终端控制进程结束(终端连接断开)
	// SIGINT 用户发送INTR字符(Ctrl+C)触发
	// SIGQUIT 用户发送QUIT字符(Ctrl+/)触发
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	sign := <-quitChan

	// 保证quitChan将不再接收信号
	signal.Stop(quitChan)

	// 控制是否启用HTTP保持活动，默认情况下始终启用保持活动，只有资源受限的环境或服务器在关闭过程中才应禁用它们
	app.Server.HttpServer.SetKeepAlivesEnabled(false)

	shutdownTimeout := time.Duration(app.Server.ShutdownTimeout)
	if shutdownTimeout < 15*time.Second {
		shutdownTimeout = 15 * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	record(fmt.Sprintf("http server shutting down by %v", sign))
	if err := app.Server.HttpServer.Shutdown(ctx); err != nil {
		record(fmt.Sprintf("http server shutting down error: %v", err), "ERROR")
	} else {
		record("http server shutdown")
	}
}
