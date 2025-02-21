package fvar_test

import (
	"testing"

	"github.com/ZYallers/fine/container/fvar"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
)

func TestVar_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := f.Map{
			"k1": "v1",
			"k2": "v2",
		}
		objOne := fvar.New(m, true)
		t.Assert(objOne.Map()["k1"], m["k1"])
		t.Assert(objOne.Map()["k2"], m["k2"])
	})
}

func TestVar_MapToMap(t *testing.T) {
	// map[int]int -> map[string]string
	// empty original map.
	ftest.C(t, func(t *ftest.T) {
		m1 := f.MapIntInt{}
		m2 := f.MapStrStr{}
		t.Assert(fvar.New(m1).MapToMap(&m2), nil)
		t.Assert(len(m1), len(m2))
	})
	// map[int]int -> map[string]string
	ftest.C(t, func(t *ftest.T) {
		m1 := f.MapIntInt{
			1: 100,
			2: 200,
		}
		m2 := f.MapStrStr{}
		t.Assert(fvar.New(m1).MapToMap(&m2), nil)
		t.Assert(m2["1"], m1[1])
		t.Assert(m2["2"], m1[2])
	})
	// map[string]interface{} -> map[string]string
	ftest.C(t, func(t *ftest.T) {
		m1 := f.Map{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := f.MapStrStr{}
		t.Assert(fvar.New(m1).MapToMap(&m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// map[string]string -> map[string]interface{}
	ftest.C(t, func(t *ftest.T) {
		m1 := f.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := f.Map{}
		t.Assert(fvar.New(m1).MapToMap(&m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// map[string]interface{} -> map[interface{}]interface{}
	ftest.C(t, func(t *ftest.T) {
		m1 := f.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := f.MapAnyAny{}
		t.Assert(fvar.New(m1).MapToMap(&m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
}

func TestVar_MapStrVar(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := f.Map{
			"k1": "v1",
			"k2": "v2",
		}
		objOne := fvar.New(m, true)
		t.Assert(objOne.MapStrVar(), "{\"k1\":\"v1\",\"k2\":\"v2\"}")

		objEmpty := fvar.New(f.Map{})
		t.Assert(objEmpty.MapStrVar(), "")
	})
}
