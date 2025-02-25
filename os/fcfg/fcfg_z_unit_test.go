package fcfg_test

import (
	"testing"

	"github.com/ZYallers/fine/os/fcfg"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_ReadConfig(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		configFile := ftest.DataPath("config.yaml")
		err := fcfg.ReadConfig(configFile, false)
		t.AssertNil(err)
		t.AssertEQ(fcfg.GetString("app.name"), "fine-example")
		t.AssertEQ(fcfg.GetString("app.version.latest"), "1.0.0")
		t.AssertEQ(fcfg.GetString("database.mysql.mall.database"), "env{mysql_mall_database}")

		config2File := ftest.DataPath("config2.yaml")
		err = fcfg.ReadConfig(config2File, true)
		t.AssertNil(err)
		t.AssertEQ(fcfg.GetString("app.version.latest"), "2.0.0")
		t.AssertEQ(fcfg.GetString("app.version.other"), "other")
	})
}
