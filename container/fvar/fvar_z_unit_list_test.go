package fvar_test

import (
	"testing"

	"github.com/ZYallers/fine/container/fvar"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
)

func TestVar_ListItemValues_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": 99},
			f.Map{"id": 3, "score": 99},
		}
		t.Assert(fvar.New(listMap).ListItemValues("id"), f.Slice{1, 2, 3})
		t.Assert(fvar.New(listMap).ListItemValues("score"), f.Slice{100, 99, 99})
	})
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": nil},
			f.Map{"id": 3, "score": 0},
		}
		t.Assert(fvar.New(listMap).ListItemValues("id"), f.Slice{1, 2, 3})
		t.Assert(fvar.New(listMap).ListItemValues("score"), f.Slice{100, nil, 0})
	})
}

func TestVar_ListItemValues_Struct(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := f.Slice{
			T{1, 100},
			T{2, 99},
			T{3, 0},
		}
		t.Assert(fvar.New(listStruct).ListItemValues("Id"), f.Slice{1, 2, 3})
		t.Assert(fvar.New(listStruct).ListItemValues("Score"), f.Slice{100, 99, 0})
	})
	// Pointer items.
	ftest.C(t, func(t *ftest.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := f.Slice{
			&T{1, 100},
			&T{2, 99},
			&T{3, 0},
		}
		t.Assert(fvar.New(listStruct).ListItemValues("Id"), f.Slice{1, 2, 3})
		t.Assert(fvar.New(listStruct).ListItemValues("Score"), f.Slice{100, 99, 0})
	})
	// Nil element value.
	ftest.C(t, func(t *ftest.T) {
		type T struct {
			Id    int
			Score interface{}
		}
		listStruct := f.Slice{
			T{1, 100},
			T{2, nil},
			T{3, 0},
		}
		t.Assert(fvar.New(listStruct).ListItemValues("Id"), f.Slice{1, 2, 3})
		t.Assert(fvar.New(listStruct).ListItemValues("Score"), f.Slice{100, nil, 0})
	})
}

func TestVar_ListItemValuesUnique(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": 100},
			f.Map{"id": 3, "score": 100},
			f.Map{"id": 4, "score": 100},
			f.Map{"id": 5, "score": 100},
		}
		t.Assert(fvar.New(listMap).ListItemValuesUnique("id"), f.Slice{1, 2, 3, 4, 5})
		t.Assert(fvar.New(listMap).ListItemValuesUnique("score"), f.Slice{100})
	})
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": 100},
			f.Map{"id": 3, "score": 100},
			f.Map{"id": 4, "score": 100},
			f.Map{"id": 5, "score": 99},
		}
		t.Assert(fvar.New(listMap).ListItemValuesUnique("id"), f.Slice{1, 2, 3, 4, 5})
		t.Assert(fvar.New(listMap).ListItemValuesUnique("score"), f.Slice{100, 99})
	})
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": 100},
			f.Map{"id": 3, "score": 0},
			f.Map{"id": 4, "score": 100},
			f.Map{"id": 5, "score": 99},
		}
		t.Assert(fvar.New(listMap).ListItemValuesUnique("id"), f.Slice{1, 2, 3, 4, 5})
		t.Assert(fvar.New(listMap).ListItemValuesUnique("score"), f.Slice{100, 0, 99})
	})
}
