package futil_test

import (
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_IsNil(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsNil(0), false)
		t.Assert(futil.IsNil(nil), true)
		t.Assert(futil.IsNil(f.Map{}), false)
		t.Assert(futil.IsNil(f.Slice{}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsNil(1), false)
		t.Assert(futil.IsNil(0.1), false)
		t.Assert(futil.IsNil(f.Map{"k": "v"}), false)
		t.Assert(futil.IsNil(f.Slice{0}), false)
	})
}

func Test_IsEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsEmpty(0), true)
		t.Assert(futil.IsEmpty(nil), true)
		t.Assert(futil.IsEmpty(f.Map{}), true)
		t.Assert(futil.IsEmpty(f.Slice{}), true)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsEmpty(1), false)
		t.Assert(futil.IsEmpty(0.1), false)
		t.Assert(futil.IsEmpty(f.Map{"k": "v"}), false)
		t.Assert(futil.IsEmpty(f.Slice{0}), false)
	})
}

func Test_IsInt(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsInt(0), true)
		t.Assert(futil.IsInt(nil), false)
		t.Assert(futil.IsInt(f.Map{}), false)
		t.Assert(futil.IsInt(f.Slice{}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsInt(1), true)
		t.Assert(futil.IsInt(-1), true)
		t.Assert(futil.IsInt(0.1), false)
		t.Assert(futil.IsInt(f.Map{"k": "v"}), false)
		t.Assert(futil.IsInt(f.Slice{0}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsInt(int8(1)), true)
		t.Assert(futil.IsInt(uint8(1)), false)
	})
}

func Test_IsUint(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsUint(0), false)
		t.Assert(futil.IsUint(nil), false)
		t.Assert(futil.IsUint(f.Map{}), false)
		t.Assert(futil.IsUint(f.Slice{}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsUint(1), false)
		t.Assert(futil.IsUint(-1), false)
		t.Assert(futil.IsUint(0.1), false)
		t.Assert(futil.IsUint(f.Map{"k": "v"}), false)
		t.Assert(futil.IsUint(f.Slice{0}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsUint(int8(1)), false)
		t.Assert(futil.IsUint(uint8(1)), true)
	})
}

func Test_IsFloat(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsFloat(0), false)
		t.Assert(futil.IsFloat(nil), false)
		t.Assert(futil.IsFloat(f.Map{}), false)
		t.Assert(futil.IsFloat(f.Slice{}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsFloat(1), false)
		t.Assert(futil.IsFloat(-1), false)
		t.Assert(futil.IsFloat(0.1), true)
		t.Assert(futil.IsFloat(float64(1)), true)
		t.Assert(futil.IsFloat(f.Map{"k": "v"}), false)
		t.Assert(futil.IsFloat(f.Slice{0}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsFloat(int8(1)), false)
		t.Assert(futil.IsFloat(uint8(1)), false)
	})
}

func Test_IsArray(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsArray([]interface{}{1, "2"}), true)
		t.Assert(futil.IsArray([]int{1, 2}), true)
		var val1, val2 int = 1, 2
		t.Assert(futil.IsArray([]*int{&val1, &val2}), true)
		type T struct {
			A string
			B int
		}
		t.Assert(futil.IsArray([]T{{A: "1", B: 2}}), true)
		t.Assert(futil.IsArray([]*T{&T{A: "1", B: 2}}), true)
	})
}

func Test_IsSlice(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsSlice(0), false)
		t.Assert(futil.IsSlice(nil), false)
		t.Assert(futil.IsSlice(f.Map{}), false)
		t.Assert(futil.IsSlice(f.Slice{}), true)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsSlice(1), false)
		t.Assert(futil.IsSlice(-1), false)
		t.Assert(futil.IsSlice(0.1), false)
		t.Assert(futil.IsSlice(float64(1)), false)
		t.Assert(futil.IsSlice(f.Map{"k": "v"}), false)
		t.Assert(futil.IsSlice(f.Slice{0}), true)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsSlice(int8(1)), false)
		t.Assert(futil.IsSlice(uint8(1)), false)
	})
}

func Test_IsMap(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsMap(0), false)
		t.Assert(futil.IsMap(nil), false)
		t.Assert(futil.IsMap(f.Map{}), true)
		t.Assert(futil.IsMap(f.Slice{}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsMap(1), false)
		t.Assert(futil.IsMap(-1), false)
		t.Assert(futil.IsMap(0.1), false)
		t.Assert(futil.IsMap(float64(1)), false)
		t.Assert(futil.IsMap(f.Map{"k": "v"}), true)
		t.Assert(futil.IsMap(f.Slice{0}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsMap(int8(1)), false)
		t.Assert(futil.IsMap(uint8(1)), false)
	})
}

func Test_IsStruct(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsStruct(0), false)
		t.Assert(futil.IsStruct(nil), false)
		t.Assert(futil.IsStruct(f.Map{}), false)
		t.Assert(futil.IsStruct(f.Slice{}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.IsStruct(1), false)
		t.Assert(futil.IsStruct(-1), false)
		t.Assert(futil.IsStruct(0.1), false)
		t.Assert(futil.IsStruct(float64(1)), false)
		t.Assert(futil.IsStruct(f.Map{"k": "v"}), false)
		t.Assert(futil.IsStruct(f.Slice{0}), false)
	})
	ftest.C(t, func(t *ftest.T) {
		a := &struct{}{}
		t.Assert(futil.IsStruct(a), true)
		t.Assert(futil.IsStruct(*a), true)
		t.Assert(futil.IsStruct(&a), true)
	})
}
