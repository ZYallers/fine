package flog

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ZYallers/fine/frame/fmsg"
	"github.com/ZYallers/fine/internal/consts"
	"github.com/ZYallers/fine/internal/instance"
	"github.com/ZYallers/fine/net/ftracing"
	"github.com/ZYallers/fine/os/fcfg"
	"github.com/ZYallers/fine/os/fgoid"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

var (
	loggerPath           string
	loggerSuffix         string
	loggerMaxSize        int
	loggerWithCaller     bool
	loggerStdoutEnable   bool
	loggerWithStacktrace string
	configSingleton      sync.Once
)

func loadConfig() {
	configSingleton.Do(func() {
		if path := fcfg.GetString("logger.path"); path != "" {
			loggerPath = path
		} else {
			loggerPath, _ = filepath.Abs(filepath.Dir("."))
		}
		if suffix := fcfg.GetString("logger.suffix"); suffix != "" {
			loggerSuffix = suffix
		} else {
			loggerSuffix = DefaultSuffix
		}
		if maxSize := fcfg.GetInt("logger.maxSize"); maxSize > 0 {
			loggerMaxSize = maxSize
		} else {
			loggerMaxSize = DefaultMaxSize
		}
		loggerWithCaller = fcfg.GetBool("logger.withCaller")
		loggerWithStacktrace = fcfg.GetString("logger.withStacktrace")
		loggerStdoutEnable = fcfg.GetBool("logger.stdoutEnable")
	})
}

func newLogger(filename string) *zap.Logger {
	loadConfig()
	loggerFilename, _ := filepath.Abs(filepath.Join(loggerPath, filename+loggerSuffix))
	lumber := &lumberjack.Logger{Filename: loggerFilename, MaxSize: loggerMaxSize, LocalTime: true}

	options := make([]zap.Option, 0)
	if loggerWithCaller {
		options = append(options, zap.AddCaller())
	}
	switch loggerWithStacktrace {
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
	if loggerStdoutEnable {
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), levelEnabler)
		cores = append(cores, consoleCore)
	}

	return zap.New(zapcore.NewTee(cores...), options...)
}

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
	result := instance.GetOrSetFunc(instanceKey, func() interface{} { return newLogger(filename) })
	switch v := result.(type) {
	case *zap.Logger:
		if id := fgoid.GetString(); id != "" {
			if traceId := ftracing.GetTraceID(id); traceId != "" {
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
