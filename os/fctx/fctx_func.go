package fctx

import (
	"fmt"
	"net/http"

	"github.com/ZYallers/fine/encoding/fjson"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/util/fcast"
	"github.com/gin-gonic/gin"
)

func QueryPostForm(ctx *gin.Context, keys ...string) string {
	if len(keys) == 0 {
		return ""
	}
	if val, ok := ctx.GetPostForm(keys[0]); ok {
		return val
	}
	if val, ok := ctx.GetQuery(keys[0]); ok {
		return val
	}
	if len(keys) == 2 {
		return keys[1]
	}
	return ""
}

func AbortJson(ctx *gin.Context, code int, msg interface{}, data ...interface{}) {
	result := f.MapStrAny{"code": code, "msg": fcast.ToString(msg), "data": nil}
	if len(data) > 0 {
		result["data"] = data[0]
	}
	ctx.Abort()
	ctx.Status(http.StatusOK)
	ctx.Header("Content-Type", "application/json")
	if bte, err := fjson.Marshal(result); err != nil {
		_, _ = ctx.Writer.WriteString(fmt.Sprintf(`{"code":500,"msg":"%v"}`, err))
	} else {
		_, _ = ctx.Writer.Write(bte)
	}
}
