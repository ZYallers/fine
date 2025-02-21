package farray_test

import (
	"strings"
	"testing"

	"github.com/ZYallers/fine/container/farray"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
	"github.com/ZYallers/fine/util/futil"
)

func Test_Array_Var(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var array farray.Array
		expect := []int{2, 3, 1}
		array.Append(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	ftest.C(t, func(t *ftest.T) {
		var array farray.IntArray
		expect := []int{2, 3, 1}
		array.Append(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	ftest.C(t, func(t *ftest.T) {
		var array farray.StrArray
		expect := []string{"b", "a"}
		array.Append("b", "a")
		t.Assert(array.Slice(), expect)
	})
	ftest.C(t, func(t *ftest.T) {
		var array farray.SortedArray
		array.SetComparator(futil.ComparatorInt)
		expect := []int{1, 2, 3}
		array.Add(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	ftest.C(t, func(t *ftest.T) {
		var array farray.SortedIntArray
		expect := []int{1, 2, 3}
		array.Add(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	ftest.C(t, func(t *ftest.T) {
		var array farray.SortedStrArray
		expect := []string{"a", "b", "c"}
		array.Add("c", "a", "b")
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedIntArray_Var(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var array farray.SortedIntArray
		expect := []int{1, 2, 3}
		array.Add(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
}

func Test_IntArray_Unique(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := []int{1, 2, 3, 4, 5, 6}
		array := farray.NewIntArray()
		array.Append(1, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6)
		array.Unique()
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedIntArray1(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		array := farray.NewSortedIntArray()
		for i := 10; i > -1; i-- {
			array.Add(i)
		}
		t.Assert(array.Slice(), expect)
		t.Assert(array.Add().Slice(), expect)
	})
}

func Test_SortedIntArray2(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		array := farray.NewSortedIntArray()
		for i := 0; i <= 10; i++ {
			array.Add(i)
		}
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedStrArray1(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array1 := farray.NewSortedStrArray()
		array2 := farray.NewSortedStrArray(true)
		for i := 10; i > -1; i-- {
			array1.Add(fconv.String(i))
			array2.Add(fconv.String(i))
		}
		t.Assert(array1.Slice(), expect)
		t.Assert(array2.Slice(), expect)
	})

}

func Test_SortedStrArray2(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array := farray.NewSortedStrArray()
		for i := 0; i <= 10; i++ {
			array.Add(fconv.String(i))
		}
		t.Assert(array.Slice(), expect)
		array.Add()
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedArray1(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array := farray.NewSortedArray(func(v1, v2 interface{}) int {
			return strings.Compare(fconv.String(v1), fconv.String(v2))
		})
		for i := 10; i > -1; i-- {
			array.Add(fconv.String(i))
		}
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedArray2(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(fconv.String(v1), fconv.String(v2))
		}
		array := farray.NewSortedArray(func1)
		array2 := farray.NewSortedArray(func1, true)
		for i := 0; i <= 10; i++ {
			array.Add(fconv.String(i))
			array2.Add(fconv.String(i))
		}
		t.Assert(array.Slice(), expect)
		t.Assert(array.Add().Slice(), expect)
		t.Assert(array2.Slice(), expect)
	})
}

func TestNewFromCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		a1 := []interface{}{"100", "200", "300", "400", "500", "600"}
		array1 := farray.NewFromCopy(a1)
		t.AssertIN(array1.PopRands(2), a1)
		t.Assert(len(array1.PopRands(1)), 1)
		t.Assert(len(array1.PopRands(9)), 3)
	})
}
