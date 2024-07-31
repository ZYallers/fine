package handler

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/os/fctx"
	"github.com/gin-gonic/gin"
)

func signCheck(app *fapp.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !doSignCheck(ctx, app) {
			fctx.AbortJson(ctx, http.StatusForbidden, "sign verification failed")
		}
	}
}

func doSignCheck(ctx *gin.Context, app *fapp.App) bool {
	key, timeKey, secret, expiration := app.Sign.Key, app.Sign.TimeKey, app.Sign.Secret, app.Sign.Expiration
	sign := fctx.QueryPostForm(ctx, key)
	if sign == "" {
		return false
	}
	sign, _ = url.QueryUnescape(sign)
	times := fctx.QueryPostForm(ctx, timeKey)
	if times == "" {
		return false
	}
	timestamp, err := strconv.ParseInt(times, 10, 0)
	if err != nil {
		return false
	}
	if time.Now().Unix()-timestamp > int64(expiration) {
		return false
	}
	hash := md5.New()
	hash.Write([]byte(times + secret))
	md5str := hex.EncodeToString(hash.Sum(nil))
	if sign == base64.StdEncoding.EncodeToString([]byte(md5str)) {
		return true
	}
	return false
}
