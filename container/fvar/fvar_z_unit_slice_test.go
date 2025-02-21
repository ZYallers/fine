package fvar_test

import (
	"testing"

	"github.com/ZYallers/fine/container/fvar"
	"github.com/ZYallers/fine/test/ftest"
)

func TestVar_Ints(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, true)
		t.Assert(objOne.Ints()[0], arr[0])
	})
}

func TestVar_Uints(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, true)
		t.Assert(objOne.Uints()[0], arr[0])
	})
}

func TestVar_Int64s(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, true)
		t.Assert(objOne.Int64s()[0], arr[0])
	})
}

func TestVar_Uint64s(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, true)
		t.Assert(objOne.Uint64s()[0], arr[0])
	})
}

func TestVar_Floats(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []float64{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, true)
		t.Assert(objOne.Floats()[0], arr[0])
	})
}

func TestVar_Float32s(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []float32{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, true)
		t.AssertEQ(objOne.Float32s(), arr)
	})
}

func TestVar_Float64s(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []float64{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, true)
		t.AssertEQ(objOne.Float64s(), arr)
	})
}

func TestVar_Strings(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []string{"hello", "world"}
		objOne := fvar.New(arr, true)
		t.Assert(objOne.Strings()[0], arr[0])
	})
}

func TestVar_Interfaces(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, true)
		t.Assert(objOne.Interfaces(), arr)
	})
}

func TestVar_Slice(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, true)
		t.Assert(objOne.Slice(), arr)
	})
}

func TestVar_Array(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, false)
		t.Assert(objOne.Array(), arr)
	})
}

func TestVar_Vars(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var arr = []int{1, 2, 3, 4, 5}
		objOne := fvar.New(arr, false)
		t.Assert(len(objOne.Vars()), 5)
		t.Assert(objOne.Vars()[0].Int(), 1)
		t.Assert(objOne.Vars()[4].Int(), 5)

		objEmpty := fvar.New([]int{})
		t.Assert(objEmpty.Vars(), nil)
	})
}
