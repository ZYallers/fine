package ftype_test

import (
	"testing"

	"github.com/ZYallers/fine/container/ftype"
	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_String(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		i := ftype.NewString("abc")
		iClone := i.Clone()
		t.AssertEQ(iClone.Set("123"), "abc")
		t.AssertEQ(iClone.Val(), "123")
		t.AssertEQ(iClone.String(), "123")
		//
		copyVal := iClone.DeepCopy()
		iClone.Set("124")
		t.AssertNE(copyVal, iClone.Val())
		iClone = nil
		copyVal = iClone.DeepCopy()
		t.AssertNil(copyVal)
		// empty param test
		i1 := ftype.NewString()
		t.AssertEQ(i1.Val(), "")
	})
}

func Test_String_JSON(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := "i love fine"
		i1 := ftype.NewString(s)
		b1, err1 := json.Marshal(i1)
		b2, err2 := json.Marshal(i1.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := ftype.NewString()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), s)
	})
}

func Test_String_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *ftype.String
	}
	ftest.C(t, func(t *ftest.T) {
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "123")
	})
}
