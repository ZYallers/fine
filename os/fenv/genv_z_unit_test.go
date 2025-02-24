package fenv_test

import (
	"os"
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/internal/command"
	"github.com/ZYallers/fine/os/fenv"
	"github.com/ZYallers/fine/os/ftime"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_GEnv_All(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(os.Environ(), fenv.All())
	})
}

func Test_GEnv_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		value := fconv.String(ftime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.Assert(fenv.Map()[key], "TEST")
	})
}

func Test_GEnv_Get(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		value := fconv.String(ftime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(fenv.Get(key).String(), "TEST")
	})
}

func Test_GEnv_GetVar(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		value := fconv.String(ftime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(fenv.Get(key).String(), "TEST")
	})
}

func Test_GEnv_Contains(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		value := fconv.String(ftime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(fenv.Contains(key), true)
		t.AssertEQ(fenv.Contains("none"), false)
	})
}

func Test_GEnv_Set(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		value := fconv.String(ftime.TimestampNano())
		key := "TEST_ENV_" + value
		err := fenv.Set(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(os.Getenv(key), "TEST")
	})
}

func Test_GEnv_SetMap(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		err := fenv.SetMap(f.MapStrStr{
			"K1": "TEST1",
			"K2": "TEST2",
		})
		t.AssertNil(err)
		t.AssertEQ(os.Getenv("K1"), "TEST1")
		t.AssertEQ(os.Getenv("K2"), "TEST2")
	})
}

func Test_GEnv_Build(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := fenv.Build(map[string]string{
			"k1": "v1",
			"k2": "v2",
		})
		t.AssertIN("k1=v1", s)
		t.AssertIN("k2=v2", s)
	})
}

func Test_GEnv_Remove(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		value := fconv.String(ftime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		err = fenv.Remove(key)
		t.AssertNil(err)
		t.AssertEQ(os.Getenv(key), "")
	})
}

func Test_GetWithCmd(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		command.Init("-test", "2")
		t.Assert(fenv.GetWithCmd("TEST"), 2)
	})
	ftest.C(t, func(t *ftest.T) {
		fenv.Set("TEST", "1")
		defer fenv.Remove("TEST")
		command.Init("-test", "2")
		t.Assert(fenv.GetWithCmd("test"), 1)
	})
}

func Test_MapFromEnv(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fenv.MapFromEnv([]string{"a=1", "b=2"})
		t.Assert(m, f.Map{"a": 1, "b": 2})
	})
}

func Test_MapToEnv(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := fenv.MapToEnv(f.MapStrStr{"a": "1"})
		t.Assert(s, []string{"a=1"})
	})
}

func Test_Filter(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := fenv.Filter([]string{"a=1", "a=3"})
		t.Assert(s, []string{"a=3"})
	})
}
