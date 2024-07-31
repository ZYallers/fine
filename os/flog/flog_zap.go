package flog

import (
	"fmt"
	"gitlab.sys.hxsapp.net/hxs/fine/net/ftracing"
	"gitlab.sys.hxsapp.net/hxs/fine/os/fgoid"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gitlab.sys.hxsapp.net/hxs/fine/frame/fmsg"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/consts"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/instance"
	"gitlab.sys.hxsapp.net/hxs/fine/os/fcfg"
)

var (
	levelEnabler  zap.LevelEnablerFunc = func(lv zapcore.Level) bool { return lv >= zapcore.DebugLevel }
	encoderConfig                      = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
	}
)

func Use(name ...string) *zap.Logger {
	var filename string
	if len(name) > 0 && name[0] != "" {
		filename = name[0]
	} else {
		filename = fcfg.GetString("app.name")
	}
	if filename == "" {
		return nil
	}
	instanceKey := fmt.Sprintf("%s.%s", consts.FrameComponentLoggerZap, filename)
	result := instance.GetOrSetFunc(instanceKey, func() interface{} {
		path := fcfg.GetString("logger.path")
		if path == "" {
			path, _ = filepath.Abs(filepath.Dir("."))
		}
		suffix := fcfg.GetString("logger.suffix")
		if suffix == "" {
			suffix = DefaultSuffix
		}

		filename, _ := filepath.Abs(filepath.Join(path, filename+suffix))
		lumber := &lumberjack.Logger{Filename: filename, MaxSize: DefaultMaxSize, LocalTime: true}
		if maxSize := fcfg.GetInt("logger.maxSize"); maxSize > 0 {
			lumber.MaxSize = maxSize
		}

		options := make([]zap.Option, 0)
		if fcfg.GetBool("logger.withCaller") {
			options = append(options, zap.AddCaller())
		}
		switch fcfg.GetString("logger.withStacktrace") {
		case "panic":
			options = append(options, zap.AddStacktrace(zapcore.PanicLevel))
		case "error":
			options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
		case "warn":
			options = append(options, zap.AddStacktrace(zapcore.WarnLevel))
		case "info":
			options = append(options, zap.AddStacktrace(zapcore.InfoLevel))
		}

		jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
		fileCore := zapcore.NewCore(jsonEncoder, zapcore.AddSync(lumber), levelEnabler)
		cores := []zapcore.Core{fileCore}
		if fcfg.GetBool("logger.stdoutEnable") {
			consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
			consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), levelEnabler)
			cores = append(cores, consoleCore)
		}

		return zap.New(zapcore.NewTee(cores...), options...)
	})
	switch v := result.(type) {
	case *zap.Logger:
		if goid := fgoid.GetString(); goid != "" {
			if traceId := ftracing.GetTraceID(goid); traceId != "" {
				return v.With(zap.String("trace_id", traceId))
			}
		}
		return v
	default:
		instance.Remove(instanceKey)
		msg := fmt.Sprintf("logger.%s.error: %s", name, "unknown error")
		fmsg.Sender().Simple(msg, true)
		return nil
	}
}
