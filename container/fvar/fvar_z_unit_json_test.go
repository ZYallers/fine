package fvar_test

import (
	"math"
	"testing"

	"github.com/ZYallers/fine/container/fvar"
	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
)

func TestVar_Json(t *testing.T) {
	// Marshal
	ftest.C(t, func(t *ftest.T) {
		s := "i love gf"
		v := fvar.New(s)
		b1, err1 := json.Marshal(v)
		b2, err2 := json.Marshal(s)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})

	ftest.C(t, func(t *ftest.T) {
		s := int64(math.MaxInt64)
		v := fvar.New(s)
		b1, err1 := json.Marshal(v)
		b2, err2 := json.Marshal(s)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})

	// Unmarshal
	ftest.C(t, func(t *ftest.T) {
		s := "i love gf"
		v := fvar.New(nil)
		b, err := json.Marshal(s)
		t.AssertNil(err)

		err = json.UnmarshalUseNumber(b, v)
		t.AssertNil(err)
		t.Assert(v.String(), s)
	})

	ftest.C(t, func(t *ftest.T) {
		var v fvar.Var
		s := "i love gf"
		b, err := json.Marshal(s)
		t.AssertNil(err)

		err = json.UnmarshalUseNumber(b, &v)
		t.AssertNil(err)
		t.Assert(v.String(), s)
	})
}
