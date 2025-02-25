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

	// SIGTERM terminates program (kill pid) (can be captured, blocked, or ignored)
	// SIGHUP terminal control process ends (terminal connection disconnected)
	// SIGINT user sends INTR character (Ctrl+C) trigger
	// SIGQUIT user sends QUIT character (Ctrl+/) trigger
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	sign := <-quitChan

	// Ensure that quitChan will no longer receive signals
	signal.Stop(quitChan)

	// Control whether to enable HTTP keep alive.
	// Keep alive is always enabled by default,
	// and should only be disabled in resource constrained environments
	// or servers during shutdown
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
