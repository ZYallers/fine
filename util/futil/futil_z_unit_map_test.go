package futil_test

import (
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_MapCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := f.Map{
			"k1": "v1",
		}
		m2 := futil.MapCopy(m1)
		m2["k2"] = "v2"
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], "v1")
		t.Assert(m2["k2"], "v2")
	})
}

func Test_MapContains(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := f.Map{
			"k1": "v1",
		}
		t.Assert(futil.MapContains(m1, "k1"), true)
		t.Assert(futil.MapContains(m1, "K1"), false)
		t.Assert(futil.MapContains(m1, "k2"), false)
		m2 := f.Map{}
		t.Assert(futil.MapContains(m2, "k1"), false)
	})
}

func Test_MapDelete(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := f.Map{
			"k1": "v1",
		}
		futil.MapDelete(m1, "k1")
		futil.MapDelete(m1, "K1")
		m2 := f.Map{}
		futil.MapDelete(m2, "k1")
	})
}

func Test_MapMerge(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := f.Map{
			"k1": "v1",
		}
		m2 := f.Map{
			"k2": "v2",
		}
		m3 := f.Map{
			"k3": "v3",
		}
		futil.MapMerge(m1, m2, m3, nil)
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], "v2")
		t.Assert(m1["k3"], "v3")
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
	})
}

func Test_MapMergeCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := f.Map{
			"k1": "v1",
		}
		m2 := f.Map{
			"k2": "v2",
		}
		m3 := f.Map{
			"k3": "v3",
		}
		m := futil.MapMergeCopy(m1, m2, m3, nil)
		t.Assert(m["k1"], "v1")
		t.Assert(m["k2"], "v2")
		t.Assert(m["k3"], "v3")
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
	})
}

func Test_MapPossibleItemByKey(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := f.Map{
			"name":     "guo",
			"NickName": "john",
		}
		k, v := futil.MapPossibleItemByKey(m, "NAME")
		t.Assert(k, "name")
		t.Assert(v, "guo")

		k, v = futil.MapPossibleItemByKey(m, "nick name")
		t.Assert(k, "NickName")
		t.Assert(v, "john")

		k, v = futil.MapPossibleItemByKey(m, "none")
		t.Assert(k, "")
		t.Assert(v, nil)
	})
}

func Test_MapContainsPossibleKey(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := f.Map{
			"name":     "guo",
			"NickName": "john",
		}
		t.Assert(futil.MapContainsPossibleKey(m, "name"), true)
		t.Assert(futil.MapContainsPossibleKey(m, "NAME"), true)
		t.Assert(futil.MapContainsPossibleKey(m, "nickname"), true)
		t.Assert(futil.MapContainsPossibleKey(m, "nick name"), true)
		t.Assert(futil.MapContainsPossibleKey(m, "nick_name"), true)
		t.Assert(futil.MapContainsPossibleKey(m, "nick-name"), true)
		t.Assert(futil.MapContainsPossibleKey(m, "nick.name"), true)
		t.Assert(futil.MapContainsPossibleKey(m, "none"), false)
	})
}

func Test_MapOmitEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := f.Map{
			"k1": "john",
			"e1": "",
			"e2": 0,
			"e3": nil,
			"k2": "smith",
		}
		futil.MapOmitEmpty(m)
		t.Assert(len(m), 2)
		t.AssertNE(m["k1"], nil)
		t.AssertNE(m["k2"], nil)
		m1 := f.Map{}
		futil.MapOmitEmpty(m1)
		t.Assert(len(m1), 0)
	})
}

func Test_MapToSlice(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := f.Map{
			"k1": "v1",
			"k2": "v2",
		}
		s := futil.MapToSlice(m)
		t.Assert(len(s), 4)
		t.AssertIN(s[0], f.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[1], f.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[2], f.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[3], f.Slice{"k1", "k2", "v1", "v2"})
		s1 := futil.MapToSlice(&m)
		t.Assert(len(s1), 4)
		t.AssertIN(s1[0], f.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s1[1], f.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s1[2], f.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s1[3], f.Slice{"k1", "k2", "v1", "v2"})
	})
	ftest.C(t, func(t *ftest.T) {
		m := f.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		s := futil.MapToSlice(m)
		t.Assert(len(s), 4)
		t.AssertIN(s[0], f.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[1], f.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[2], f.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[3], f.Slice{"k1", "k2", "v1", "v2"})
	})
	ftest.C(t, func(t *ftest.T) {
		m := f.MapStrStr{}
		s := futil.MapToSlice(m)
		t.Assert(len(s), 0)
	})
	ftest.C(t, func(t *ftest.T) {
		s := futil.MapToSlice(1)
		t.Assert(s, nil)
	})
}
