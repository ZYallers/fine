package ftype_test

import (
	"testing"

	"github.com/ZYallers/fine/container/ftype"
	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_Bool(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		i := ftype.NewBool(true)
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(false), true)
		t.AssertEQ(iClone.Val(), false)

		i1 := ftype.NewBool(false)
		iClone1 := i1.Clone()
		t.AssertEQ(iClone1.Set(true), false)
		t.AssertEQ(iClone1.Val(), true)

		t.AssertEQ(iClone1.Cas(false, true), false)
		t.AssertEQ(iClone1.String(), "true")
		t.AssertEQ(iClone1.Cas(true, false), true)
		t.AssertEQ(iClone1.String(), "false")

		copyVal := i1.DeepCopy()
		iClone.Set(true)
		t.AssertNE(copyVal, iClone.Val())
		iClone = nil
		copyVal = iClone.DeepCopy()
		t.AssertNil(copyVal)

		// empty param test
		i2 := ftype.NewBool()
		t.AssertEQ(i2.Val(), false)
	})
}

func Test_Bool_JSON(t *testing.T) {
	// Marshal
	ftest.C(t, func(t *ftest.T) {
		i := ftype.NewBool(true)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)
	})
	ftest.C(t, func(t *ftest.T) {
		i := ftype.NewBool(false)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)
	})
	// Unmarshal
	ftest.C(t, func(t *ftest.T) {
		var err error
		i := ftype.NewBool()
		err = json.UnmarshalUseNumber([]byte("true"), &i)
		t.AssertNil(err)
		t.Assert(i.Val(), true)
		err = json.UnmarshalUseNumber([]byte("false"), &i)
		t.AssertNil(err)
		t.Assert(i.Val(), false)
		err = json.UnmarshalUseNumber([]byte("1"), &i)
		t.AssertNil(err)
		t.Assert(i.Val(), true)
		err = json.UnmarshalUseNumber([]byte("0"), &i)
		t.AssertNil(err)
		t.Assert(i.Val(), false)
	})

	ftest.C(t, func(t *ftest.T) {
		i := ftype.NewBool(true)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := ftype.NewBool()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), i.Val())
	})
	ftest.C(t, func(t *ftest.T) {
		i := ftype.NewBool(false)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := ftype.NewBool()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), i.Val())
	})
}

func Test_Bool_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *ftype.Bool
	}
	ftest.C(t, func(t *ftest.T) {
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "true",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), true)
	})
	ftest.C(t, func(t *ftest.T) {
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "false",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), false)
	})
}
