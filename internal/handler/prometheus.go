package handler

import (
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	promSingleton sync.Once
	prom          = promhttp.Handler()
	appInfo       = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "go_app_info",
			Help: "now running go app information.",
		},
		[]string{"name", "cmdline"},
	)
)

func Prometheus(name string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		promSingleton.Do(func() {
			prometheus.MustRegister(appInfo)
			appInfo.WithLabelValues(name, strings.Join(os.Args, " ")).Inc()
		})
		prom.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
