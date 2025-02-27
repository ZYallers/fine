package handler

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/net/ftracing"
	"github.com/ZYallers/fine/os/fcfg"
	"github.com/ZYallers/fine/os/fctx"
	"github.com/ZYallers/fine/os/fgoid"
	"github.com/ZYallers/fine/os/flog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Tracing(app *fapp.App) {
	app.Server.HttpServer.Handler.(*gin.Engine).Use(handler())
}

func excludeRoutes() map[string]bool {
	if slices := fcfg.GetStringSlice("tracing.exclude_routes"); slices != nil {
		routes := make(map[string]bool, len(slices))
		for _, s := range slices {
			routes[s] = true
		}
		return routes
	}
	return nil
}

func handler() gin.HandlerFunc {
	excludeRoutes := excludeRoutes()
	traceLogger := flog.Use("tracing")
	return func(ctx *gin.Context) {
		if excludeRoutes[ctx.Request.URL.Path[1:]] {
			return
		}

		goid := fgoid.GetString()
		defer ftracing.DelTraceID(goid)

		traceId := ctx.GetHeader(ftracing.TraceIDKey)
		if traceId == "" {
			traceId = ftracing.NewTraceID()
		}
		ftracing.SetTraceID(goid, traceId)
		ctx.Set(ftracing.TraceIDKey, traceId)
		ctx.Header(ftracing.TraceIDKey, traceId)

		traceLogger.Info("",
			zap.String("trace_id", traceId),
			zap.String("host", ctx.Request.Host),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("client_ip", ctx.ClientIP()),
			zap.String("req_raw", ctx.GetString(fctx.RequestRawKey)),
		)

		// Must be added ctx.Next(), otherwise there is an issue with the execution order of the defer func above
		ctx.Next()
	}
}
