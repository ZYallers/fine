package intlog

import (
	"log"

	"go.uber.org/zap"
)

type goRedisLogger struct {
	*zap.Logger
}

func NewGoRedisLogger(logger *zap.Logger) *log.Logger {
	return log.New(&goRedisLogger{logger}, "", log.LstdFlags|log.Lshortfile)
}

func (l *goRedisLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}
