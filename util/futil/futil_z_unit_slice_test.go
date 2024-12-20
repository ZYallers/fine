package futil_test

import (
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/text/fstr"
	"github.com/ZYallers/fine/util/futil"
)

func Test_SliceCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := f.Slice{
			"K1", "v1", "K2", "v2",
		}
		s1 := futil.SliceCopy(s)
		t.Assert(s, s1)
	})
}

func Test_SliceDelete(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := f.Slice{
			"K1", "v1", "K2", "v2",
		}
		t.Assert(futil.SliceDelete(s, 0), f.Slice{
			"v1", "K2", "v2",
		})
		t.Assert(futil.SliceDelete(s, 5), s)
	})
}

func Test_SliceToMap(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := f.Slice{
			"K1", "v1", "K2", "v2",
		}
		m := futil.SliceToMap(s)
		t.Assert(len(m), 2)
		t.Assert(m, f.Map{
			"K1": "v1",
			"K2": "v2",
		})

		m1 := futil.SliceToMap(&s)
		t.Assert(len(m1), 2)
		t.Assert(m1, f.Map{
			"K1": "v1",
			"K2": "v2",
		})
	})
	ftest.C(t, func(t *ftest.T) {
		s := f.Slice{
			"K1", "v1", "K2",
		}
		m := futil.SliceToMap(s)
		t.Assert(len(m), 0)
		t.Assert(m, nil)
	})
	ftest.C(t, func(t *ftest.T) {
		m := futil.SliceToMap(1)
		t.Assert(len(m), 0)
		t.Assert(m, nil)
	})
}

func Test_SliceToMapWithColumnAsKey(t *testing.T) {
	m1 := f.Map{"K1": "v1", "K2": 1}
	m2 := f.Map{"K1": "v2", "K2": 2}
	s := f.Slice{m1, m2}
	ftest.C(t, func(t *ftest.T) {
		m := futil.SliceToMapWithColumnAsKey(s, "K1")
		t.Assert(m, f.MapAnyAny{
			"v1": m1,
			"v2": m2,
		})

		n := futil.SliceToMapWithColumnAsKey(&s, "K1")
		t.Assert(n, f.MapAnyAny{
			"v1": m1,
			"v2": m2,
		})
	})
	ftest.C(t, func(t *ftest.T) {
		m := futil.SliceToMapWithColumnAsKey(s, "K2")
		t.Assert(m, f.MapAnyAny{
			1: m1,
			2: m2,
		})
	})
}

func Test_SliceInsertBefore(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := f.Slice{
			0, 1, 2, 3, 4,
		}
		s2 := futil.SliceInsertBefore(s, 1, 8, 9)
		t.Assert(fstr.JoinAny(s2, " "), "0 8 9 1 2 3 4")
	})
}

func Test_SliceInsertAfter(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := f.Slice{
			0, 1, 2, 3, 4,
		}
		s2 := futil.SliceInsertAfter(s, 1, 8, 9)
		t.Assert(fstr.JoinAny(s2, " "), "0 1 8 9 2 3 4")
	})
}
