package message

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.sys.hxsapp.net/hxs/fine/net/fclient"
	"gitlab.sys.hxsapp.net/hxs/fine/net/fipv4"
	"gitlab.sys.hxsapp.net/hxs/fine/os/fcfg"
)

type DingTalk struct {
	enable  bool
	simple  string
	context string
}

var singleton sync.Once

func (dt *DingTalk) Simple(msg interface{}, atAll bool) {
	dt.lazyInit()
	if dt.enable {
		push(dt.simple, msg, nil, "", "", atAll)
	}
}

func (dt *DingTalk) Context(ctx *gin.Context, msg interface{}, reqStr string, stack string, atAll bool) {
	dt.lazyInit()
	if dt.enable {
		push(dt.context, msg, ctx, reqStr, stack, atAll)
	}
}

func (dt *DingTalk) lazyInit() {
	singleton.Do(func() {
		dt.enable = fcfg.GetBool("message.dingtalk.enable")
		dt.simple = fcfg.GetString("message.dingtalk.token.simple")
		dt.context = fcfg.GetString("message.dingtalk.token.context")
	})
}

func push(token string, msg interface{}, ctx *gin.Context, reqStr string, stack string, atAll bool) {
	title := fmt.Sprint(msg)
	if token == "" || title == "" {
		return
	}
	hostname, _ := os.Hostname()
	text := []string{title + "\n---------------------------",
		"App: " + fcfg.GetString("app.name"),
		"Mode: " + fcfg.GetString("app.mode"),
		"Listen: " + fcfg.GetString("app.server.addr"),
		"HostName: " + hostname,
		"Time: " + time.Now().Format("2006/01/02 15:04:05.000"),
		"SystemIP: " + fipv4.SystemIP(),
	}
	if fcfg.GetBool("app.debug") {
		atAll = false
	}
	if ctx != nil {
		text = append(text, "ClientIP: "+ctx.ClientIP(), "Url: "+"https://"+ctx.Request.Host+ctx.Request.URL.String())
	}
	if reqStr != "" {
		text = append(text, "\nRequest:\n"+strings.ReplaceAll(reqStr, "\n", ""))
	}
	if stack != "" {
		text = append(text, "\nStack:\n"+stack)
	}
	data := map[string]interface{}{
		"msgtype": "text",
		"at":      map[string]interface{}{"isAtAll": atAll},
		"text":    map[string]string{"content": strings.Join(text, "\n") + "\n"},
	}
	_, _ = fclient.NewRequest("https://oapi.dingtalk.com/robot/send?access_token=" + token).
		SetContentType("application/json").
		SetTimeOut(3 * time.Second).
		SetPostData(data).
		Post()
}
