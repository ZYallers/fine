package futil_test

import (
	"reflect"
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_Keys(t *testing.T) {
	// not support int
	ftest.C(t, func(t *ftest.T) {
		var val int = 1
		keys := futil.Keys(reflect.ValueOf(val))
		t.AssertEQ(len(keys), 0)
	})
	// map
	ftest.C(t, func(t *ftest.T) {
		keys := futil.Keys(map[int]int{1: 10, 2: 20})
		t.AssertIN("1", keys)
		t.AssertIN("2", keys)

		strKeys := futil.Keys(map[string]interface{}{"key1": 1, "key2": 2})
		t.AssertIN("key1", strKeys)
		t.AssertIN("key2", strKeys)
	})
	// *map
	ftest.C(t, func(t *ftest.T) {
		keys := futil.Keys(&map[int]int{1: 10, 2: 20})
		t.AssertIN("1", keys)
		t.AssertIN("2", keys)
	})
	// *struct
	ftest.C(t, func(t *ftest.T) {
		type T struct {
			A string
			B int
		}
		keys := futil.Keys(new(T))
		t.Assert(keys, f.SliceStr{"A", "B"})
	})
	// *struct nil
	ftest.C(t, func(t *ftest.T) {
		type T struct {
			A string
			B int
		}
		var pointer *T
		keys := futil.Keys(pointer)
		t.Assert(keys, f.SliceStr{"A", "B"})
	})
	// **struct nil
	ftest.C(t, func(t *ftest.T) {
		type T struct {
			A string
			B int
		}
		var pointer *T
		keys := futil.Keys(&pointer)
		t.Assert(keys, f.SliceStr{"A", "B"})
	})
}

func Test_Values(t *testing.T) {
	// not support int
	ftest.C(t, func(t *ftest.T) {
		var val int = 1
		keys := futil.Values(reflect.ValueOf(val))
		futil.Dump(keys)
		t.AssertEQ(len(keys), 0)
	})
	// map
	ftest.C(t, func(t *ftest.T) {
		values := futil.Values(map[int]int{1: 10, 2: 20})
		futil.Dump(values)
		t.AssertIN(10, values)
		t.AssertIN(20, values)
	})
	// *map
	ftest.C(t, func(t *ftest.T) {
		keys := futil.Values(&map[int]int{1: 10, 2: 20})
		futil.Dump(keys)
		t.AssertIN(10, keys)
		t.AssertIN(20, keys)
	})
	// struct
	ftest.C(t, func(t *ftest.T) {
		type T struct {
			A string
			B int
		}
		keys := futil.Values(T{A: "1", B: 2})
		futil.Dump(keys)
		t.Assert(keys, f.Slice{"1", 2})
	})
}
