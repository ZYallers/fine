package handler

import (
	"fmt"
	"time"

	"github.com/ZYallers/fine/frame/fmsg"
	"github.com/ZYallers/fine/internal/util/fsafe"
	"github.com/ZYallers/fine/os/fctx"
	"github.com/ZYallers/fine/os/flog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerWithZap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
func LoggerWithZap(logMaxSec, sendMaxSec int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()
		go logSend(ctx.Copy(), time.Now().Sub(start), logMaxSec, sendMaxSec)
	}
}

func logSend(ctx *gin.Context, runtime time.Duration, logMaxSec, sendMaxSec int) {
	defer fsafe.Defer()
	if len(ctx.Errors) > 0 {
		reqStr := ctx.GetString(fctx.RequestRawKey)
		for i, err := range ctx.Errors.Errors() {
			if ctx.Errors[i].Type != gin.ErrorTypeBind {
				flog.Use().Error(err)
				fmsg.Sender().Context(ctx, err, reqStr, "", true)
			}
		}
	}
	if logMaxSec > 0 && runtime.Seconds() >= float64(logMaxSec) {
		flog.Use("timeout").Info(ctx.Request.URL.Path,
			zap.Duration("runtime", runtime),
			zap.String("host", ctx.Request.Host),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("client_ip", ctx.ClientIP()),
			zap.String("req_raw", ctx.GetString(fctx.RequestRawKey)),
		)
	}
	if sendMaxSec > 0 && runtime.Seconds() >= float64(sendMaxSec) {
		msg := fmt.Sprintf("%s take %s to response, exceeding the maximum %d limit",
			ctx.Request.URL.Path, runtime, sendMaxSec)
		fmsg.Sender().Context(ctx, msg, ctx.GetString(fctx.RequestRawKey), "", false)
	}
}
