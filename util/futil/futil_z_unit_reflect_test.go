package futil_test

import (
	"reflect"
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_OriginValueAndKind(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var s = "s"
		out := futil.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	ftest.C(t, func(t *ftest.T) {
		var s = "s"
		out := futil.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	ftest.C(t, func(t *ftest.T) {
		var s []int
		out := futil.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	ftest.C(t, func(t *ftest.T) {
		var s []int
		out := futil.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}

func Test_OriginTypeAndKind(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var s = "s"
		out := futil.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	ftest.C(t, func(t *ftest.T) {
		var s = "s"
		out := futil.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	ftest.C(t, func(t *ftest.T) {
		var s []int
		out := futil.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	ftest.C(t, func(t *ftest.T) {
		var s []int
		out := futil.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}

func Test_CanCallIsNil(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.CanCallIsNil(""), false)
		t.Assert(futil.CanCallIsNil(0), false)
		t.Assert(futil.CanCallIsNil(nil), false)
		t.Assert(futil.CanCallIsNil(f.Slice{}), false)
		t.Assert(futil.CanCallIsNil(f.Map{}), false)
		type T struct {
			A int
			B string
		}
		t.Assert(futil.CanCallIsNil(T{}), false)
		t.Assert(futil.CanCallIsNil(&T{}), false)
		var t2 *T
		t.Assert(futil.CanCallIsNil(t2), false)
		var ch chan int
		t.Assert(futil.CanCallIsNil(ch), false)
	})
}
