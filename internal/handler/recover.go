package handler

import (
	"gitlab.sys.hxsapp.net/hxs/fine/os/fctx"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"gitlab.sys.hxsapp.net/hxs/fine/frame/fmsg"
	"gitlab.sys.hxsapp.net/hxs/fine/os/flog"
	"gitlab.sys.hxsapp.net/hxs/fine/util/fcast"
)

// RecoverWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
func RecoverWithZap(isDebug bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// SetPanicOnFault 控制程序在意外（非零）地址出现故障时运行时的行为
		// 此类故障通常是由运行时内存损坏等错误引起的，因此默认响应是使程序崩溃
		// 在不太严重的情况下，使用内存映射文件或不安全的内存操作的程序可能会导致非空地址的错误
		// 允许此类程序请求运行时仅触发恐慌，而不是崩溃
		// 仅适用于当前 goroutine。它返回之前的设置。
		defer debug.SetPanicOnFault(debug.SetPanicOnFault(true))

		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				errStr := fcast.ToString(err)
				msg := "recover from panic: " + errStr
				reqRaw := ctx.GetString(fctx.RequestRawKey)
				stacks := string(debug.Stack())

				fmsg.Sender().Context(ctx, msg, reqRaw, stacks, true)
				flog.Use().Error(msg, zap.String("req_raw", reqRaw), zap.String("stack", stacks))

				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					_ = ctx.Error(err.(error)) // nolint: errorcheck
					ctx.Abort()
					return
				}

				data := map[string]interface{}{"error": errStr}
				if isDebug {
					data["request"] = reqRaw
					data["stack"] = strings.Split(stacks, "\n")
				}
				fctx.AbortJson(ctx, http.StatusInternalServerError, "internal server error", data)
			}
		}()

		bytes, _ := httputil.DumpRequest(ctx.Request, true)
		ctx.Set(fctx.RequestRawKey, string(bytes))

		ctx.Next()
	}
}
