package handler

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/ZYallers/fine/frame/fmsg"
	"github.com/ZYallers/fine/os/fctx"
	"github.com/ZYallers/fine/os/flog"
	"github.com/ZYallers/fine/util/fcast"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RecoverWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
func RecoverWithZap(isDebug bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		// SetPanicOnFault controls the behavior of the program when it runs in the event of
		// an unexpected (non-zero) address failure.
		// This type of failure is usually caused by errors such as runtime memory corruption,
		// so the default response is to cause the program to crash.
		// In less severe cases, programs that use memory mapped files or unsafe memory operations
		// may result in not null address errors.
		// Allow such programs to only trigger panic during runtime, rather than crashing.
		// Only applicable to the current goroutine. It returns to the previous settings.
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
