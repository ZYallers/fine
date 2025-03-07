package flog_test

import (
	"testing"

	"github.com/ZYallers/fine/net/ftracing"
	"github.com/ZYallers/fine/os/fcfg"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/os/fgoid"
	"github.com/ZYallers/fine/os/flog"
	"github.com/ZYallers/fine/test/ftest"
	"go.uber.org/zap"
)

func Test_Use(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		err := fcfg.ReadConfig(ftest.DataPath("config.yaml"), false)
		t.AssertNil(err)
		loggerPath := fcfg.GetString("logger.path")
		t.Assert(loggerPath, "./testdata/log")

		lg := flog.Use()
		t.AssertNil(lg)

		lg2 := flog.Use("test")
		t.Assert(lg2 != nil, true)

		lg2.Info("test msg")
		t.Assert(ffile.Exists(loggerPath+"/test.log"), true)

		ftracing.SetTraceID(fgoid.Get(), ftracing.NewTraceID())
		lg3 := flog.Use("test2")
		t.Assert(lg3 != nil, true)
		lg3.Debug("test msg", zap.String("a", "b"))
		t.Assert(ffile.Exists(loggerPath+"/test2.log"), true)
	})
}
