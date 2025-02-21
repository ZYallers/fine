package ftype_test

import (
	"testing"

	"github.com/ZYallers/fine/container/ftype"
	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_Interface(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t1 := Temp{Name: "gf", Age: 18}
		t2 := Temp{Name: "gf", Age: 19}
		i := ftype.New(t1)
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(t2), t1)
		t.AssertEQ(iClone.Val().(Temp), t2)

		// empty param test
		i1 := ftype.New()
		t.AssertEQ(i1.Val(), nil)

		i2 := ftype.New("gf")
		t.AssertEQ(i2.String(), "gf")
		copyVal := i2.DeepCopy()
		i2.Set("goframe")
		t.AssertNE(copyVal, iClone.Val())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Interface_JSON(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := "i love gf"
		i := ftype.New(s)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := ftype.New()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), s)
	})
}

func Test_Interface_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *ftype.Interface
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
