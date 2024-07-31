package intlog

import (
	"context"
	"errors"
	"fmt"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormless "gorm.io/gorm/logger"

	"gitlab.sys.hxsapp.net/hxs/fine/net/fclient"
	"gitlab.sys.hxsapp.net/hxs/fine/net/fipv4"
	"gitlab.sys.hxsapp.net/hxs/fine/os/fcfg"
	"gitlab.sys.hxsapp.net/hxs/fine/os/flog"
)

const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

const (
	infoStr      = "%s\n[info] "
	warnStr      = "%s\n[warn] "
	errStr       = "%s\n[error] "
	traceStr     = "%s\n[%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"

	infoColorStr      = Green + "%s\n" + Reset + Green + "[info] " + Reset
	warnColorStr      = BlueBold + "%s\n" + Reset + Magenta + "[warn] " + Reset
	errColorStr       = Magenta + "%s\n" + Reset + Red + "[error] " + Reset
	traceColorStr     = Green + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
	traceWarnColorStr = Green + "%s " + Yellow + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Magenta + " %s" + Reset
	traceErrColorStr  = RedBold + "%s " + MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
)

const (
	fileCompileRegular   = "/(service|model|logic|dao)/"
	defaultSlowThreshold = 500 * time.Millisecond
)

type gormLoggerSender struct {
	Enable bool
	Token  string
}

type gormLogger struct {
	gormless.Config
	infoStr      string
	warnStr      string
	errStr       string
	traceStr     string
	traceErrStr  string
	traceWarnStr string
	Writer       *zap.Logger
	Sender       gormLoggerSender
}

func getLogLevel(name string) gormless.LogLevel {
	switch fstr.ToLower(fcfg.GetString(fmt.Sprintf("database.mysql.%s.logger.level", name))) {
	case "info":
		return gormless.Info
	case "warn":
		return gormless.Warn
	case "error":
		return gormless.Error
	default:
		return gormless.Silent
	}
}

func NewGormLogger(name string) gormless.Interface {
	slowThreshold := fcfg.GetDuration(fmt.Sprintf("database.mysql.%s.logger.slowThreshold", name)) * time.Millisecond
	if slowThreshold <= 0 {
		slowThreshold = defaultSlowThreshold
	}
	colorful := fcfg.GetBool(fmt.Sprintf("database.mysql.%s.logger.colorful", name))
	sender := gormLoggerSender{
		Enable: fcfg.GetBool(fmt.Sprintf("database.mysql.%s.logger.sender.enable", name)),
		Token:  fcfg.GetString(fmt.Sprintf("database.mysql.%s.logger.sender.token", name)),
	}
	gl := &gormLogger{
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
		Writer:       flog.Use(name),
		Sender:       sender,
		Config:       gormless.Config{SlowThreshold: slowThreshold, LogLevel: getLogLevel(name), Colorful: colorful},
	}
	if colorful {
		gl.infoStr, gl.warnStr, gl.errStr, gl.traceStr, gl.traceWarnStr, gl.traceErrStr =
			infoColorStr, warnColorStr, errColorStr, traceColorStr, traceWarnColorStr, traceErrColorStr
	}
	return gl
}

func (l *gormLogger) LogMode(level gormless.LogLevel) gormless.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *gormLogger) push(msg string) {
	if !l.Sender.Enable || l.Sender.Token == "" {
		return
	}
	title := fmt.Sprint(msg)
	hostname, _ := os.Hostname()
	text := []string{title + "\n---------------------------",
		"App: " + fcfg.GetString("app.name"),
		"Mode: " + fcfg.GetString("app.mode"),
		"Listen: " + fcfg.GetString("app.server.addr"),
		"HostName: " + hostname,
		"Time: " + time.Now().Format("2006/01/02 15:04:05.000"),
		"SystemIP: " + fipv4.SystemIP(),
	}
	data := map[string]interface{}{
		"msgtype": "text",
		"at":      map[string]interface{}{"isAtAll": false},
		"text":    map[string]string{"content": strings.Join(text, "\n") + "\n"},
	}
	_, _ = fclient.NewRequest("https://oapi.dingtalk.com/robot/send?access_token=" + l.Sender.Token).
		SetContentType("application/json").
		SetTimeOut(3 * time.Second).
		SetPostData(data).
		Post()
}

func (l *gormLogger) Printf(level gormless.LogLevel, format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	switch level {
	case gormless.Error:
		if l.Writer != nil {
			l.Writer.Error(str)
		}
	case gormless.Warn:
		if l.Writer != nil {
			l.Writer.Warn(str)
		}
	default:
		if l.Writer != nil {
			l.Writer.Info(str)
		}
	}
	l.push(str)
}

func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormless.Info {
		l.Printf(gormless.Info, l.infoStr+msg, append([]interface{}{fileWithLineNum()}, data...)...)
	}
}

func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormless.Warn {
		l.Printf(gormless.Warn, l.warnStr+msg, append([]interface{}{fileWithLineNum()}, data...)...)
	}
}

func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormless.Error {
		l.Printf(gormless.Error, l.errStr+msg, append([]interface{}{fileWithLineNum()}, data...)...)
	}
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= gormless.Error:
			sql, rows := fc()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				l.Printf(gormless.Info, l.traceStr, fileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			} else {
				l.Printf(gormless.Error, l.traceErrStr, fileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormless.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			l.Printf(gormless.Warn, l.traceWarnStr, slowLog, fileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		case l.LogLevel >= gormless.Info:
			sql, rows := fc()
			l.Printf(gormless.Info, l.traceStr, fileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

func fileWithLineNum() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && strings.HasSuffix(file, ".go") && regexp.MustCompile(fileCompileRegular).MatchString(file) {
			return fmt.Sprintf("%s:%d", file, line)
		}
	}
	return ""
}
