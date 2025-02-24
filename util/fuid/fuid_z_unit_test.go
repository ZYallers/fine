package fuid_test

import (
	"testing"

	"github.com/ZYallers/fine/container/fset"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fuid"
)

func Test_S(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		set := fset.NewStrSet()
		for i := 0; i < 1000000; i++ {
			s := fuid.S()
			t.Assert(set.AddIfNotExist(s), true)
			t.Assert(len(s), 32)
		}
	})
}

func Test_S_Data(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(len(fuid.S([]byte("123"))), 32)
	})
}
