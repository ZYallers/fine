package ftype_test

import (
	"math"
	"testing"

	"github.com/ZYallers/fine/container/ftype"
	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_Float32(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		i := ftype.NewFloat32(0)
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(0.1), float32(0))
		t.AssertEQ(iClone.Val(), float32(0.1))

		// empty param test
		i1 := ftype.NewFloat32()
		t.AssertEQ(i1.Val(), float32(0))

		i2 := ftype.NewFloat32(1.23)
		t.AssertEQ(i2.Add(3.21), float32(4.44))
		t.AssertEQ(i2.Cas(4.45, 5.55), false)
		t.AssertEQ(i2.Cas(4.44, 5.55), true)
		t.AssertEQ(i2.String(), "5.55")

		copyVal := i2.DeepCopy()
		i2.Set(float32(6.66))
		t.AssertNE(copyVal, iClone.Val())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Float32_JSON(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		v := float32(math.MaxFloat32)
		i := ftype.NewFloat32(v)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())

		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := ftype.NewFloat32()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), v)
	})
}

func Test_Float32_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *ftype.Float32
	}
	ftest.C(t, func(t *ftest.T) {
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123.456",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "123.456")
	})
}
