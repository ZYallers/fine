package futil_test

import (
	"testing"

	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_GetOrDefaultStr(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.GetOrDefaultStr("a", "b"), "b")
		t.Assert(futil.GetOrDefaultStr("a", "b", "c"), "b")
		t.Assert(futil.GetOrDefaultStr("a"), "a")
	})
}

func Test_GetOrDefaultAny(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.GetOrDefaultAny("a", "b"), "b")
		t.Assert(futil.GetOrDefaultAny("a", "b", "c"), "b")
		t.Assert(futil.GetOrDefaultAny("a"), "a")
		t.Assert(futil.GetOrDefaultAny("a", nil), "a")
		t.Assert(futil.GetOrDefaultAny("a", ""), "")
	})
}
