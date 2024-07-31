package ftracing

import (
	"github.com/ZYallers/fine/internal/tracing"
	"github.com/google/uuid"
)

const (
	TraceIDKey    = "Trace-ID"
	cacheMaxBytes = 32 * 1024 * 1024
)

var intCache tracing.Cache

func init() {
	intCache = tracing.NewFastCache(cacheMaxBytes)
}

func Cache() tracing.Cache {
	return intCache
}

func NewTraceID() string {
	return uuid.NewString()
}

func HasTraceID(key string) bool {
	return intCache.Exist([]byte(key))
}

func GetTraceID(key string) string {
	return string(intCache.Get([]byte(key)))
}

func SetTraceID(key string, value string) {
	intCache.Set([]byte(key), []byte(value))
}

func DelTraceID(key string) {
	intCache.Del([]byte(key))
}
