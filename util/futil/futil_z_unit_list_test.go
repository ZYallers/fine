package futil_test

import (
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_ListItemValues_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": 99},
			f.Map{"id": 3, "score": 98},
		}
		t.Assert(futil.ListItemValues(listMap, "id"), f.Slice{1, 2, 3})
		t.Assert(futil.ListItemValues(&listMap, "id"), f.Slice{1, 2, 3})
		t.Assert(futil.ListItemValues(listMap, "score"), f.Slice{100, 99, 98})
	})
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": nil},
			f.Map{"id": 3, "score": 0},
		}
		t.Assert(futil.ListItemValues(listMap, "id"), f.Slice{1, 2, 3})
		t.Assert(futil.ListItemValues(listMap, "score"), f.Slice{100, nil, 0})
	})
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{}
		t.Assert(len(futil.ListItemValues(listMap, "id")), 0)
	})
}

func Test_ListItemValues_Map_SubKey(t *testing.T) {
	type Scores struct {
		Math    int
		English int
	}
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "scores": Scores{100, 60}},
			f.Map{"id": 2, "scores": Scores{0, 100}},
			f.Map{"id": 3, "scores": Scores{59, 99}},
		}
		t.Assert(futil.ListItemValues(listMap, "scores", "Math"), f.Slice{100, 0, 59})
		t.Assert(futil.ListItemValues(listMap, "scores", "English"), f.Slice{60, 100, 99})
		t.Assert(futil.ListItemValues(listMap, "scores", "PE"), f.Slice{})
	})
}

func Test_ListItemValues_Map_Array_SubKey(t *testing.T) {
	type Scores struct {
		Math    int
		English int
	}
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "scores": []Scores{{1, 2}, {3, 4}}},
			f.Map{"id": 2, "scores": []Scores{{5, 6}, {7, 8}}},
			f.Map{"id": 3, "scores": []Scores{{9, 10}, {11, 12}}},
		}
		t.Assert(futil.ListItemValues(listMap, "scores", "Math"), f.Slice{1, 3, 5, 7, 9, 11})
		t.Assert(futil.ListItemValues(listMap, "scores", "English"), f.Slice{2, 4, 6, 8, 10, 12})
		t.Assert(futil.ListItemValues(listMap, "scores", "PE"), f.Slice{})
	})
}

func Test_ListItemValues_Struct(t *testing.T) {
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
		t.Assert(futil.ListItemValues(listStruct, "Id"), f.Slice{1, 2, 3})
		t.Assert(futil.ListItemValues(listStruct, "Score"), f.Slice{100, 99, 0})
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
		t.Assert(futil.ListItemValues(listStruct, "Id"), f.Slice{1, 2, 3})
		t.Assert(futil.ListItemValues(listStruct, "Score"), f.Slice{100, 99, 0})
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
		t.Assert(futil.ListItemValues(listStruct, "Id"), f.Slice{1, 2, 3})
		t.Assert(futil.ListItemValues(listStruct, "Score"), f.Slice{100, nil, 0})
	})
}

func Test_ListItemValues_Struct_SubKey(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type Student struct {
			Id    int
			Score float64
		}
		type Class struct {
			Total    int
			Students []Student
		}
		listStruct := f.Slice{
			Class{2, []Student{{1, 1}, {2, 2}}},
			Class{3, []Student{{3, 3}, {4, 4}, {5, 5}}},
			Class{1, []Student{{6, 6}}},
		}
		t.Assert(futil.ListItemValues(listStruct, "Total"), f.Slice{2, 3, 1})
		t.Assert(futil.ListItemValues(listStruct, "Students"), `[[{"Id":1,"Score":1},{"Id":2,"Score":2}],[{"Id":3,"Score":3},{"Id":4,"Score":4},{"Id":5,"Score":5}],[{"Id":6,"Score":6}]]`)
		t.Assert(futil.ListItemValues(listStruct, "Students", "Id"), f.Slice{1, 2, 3, 4, 5, 6})
	})
	ftest.C(t, func(t *ftest.T) {
		type Student struct {
			Id    int
			Score float64
		}
		type Class struct {
			Total    int
			Students []*Student
		}
		listStruct := f.Slice{
			&Class{2, []*Student{{1, 1}, {2, 2}}},
			&Class{3, []*Student{{3, 3}, {4, 4}, {5, 5}}},
			&Class{1, []*Student{{6, 6}}},
		}
		t.Assert(futil.ListItemValues(listStruct, "Total"), f.Slice{2, 3, 1})
		t.Assert(futil.ListItemValues(listStruct, "Students"), `[[{"Id":1,"Score":1},{"Id":2,"Score":2}],[{"Id":3,"Score":3},{"Id":4,"Score":4},{"Id":5,"Score":5}],[{"Id":6,"Score":6}]]`)
		t.Assert(futil.ListItemValues(listStruct, "Students", "Id"), f.Slice{1, 2, 3, 4, 5, 6})
	})
}

func Test_ListItemValuesUnique(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": 100},
			f.Map{"id": 3, "score": 100},
			f.Map{"id": 4, "score": 100},
			f.Map{"id": 5, "score": 100},
		}
		t.Assert(futil.ListItemValuesUnique(listMap, "id"), f.Slice{1, 2, 3, 4, 5})
		t.Assert(futil.ListItemValuesUnique(listMap, "score"), f.Slice{100})
	})
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": 100},
			f.Map{"id": 3, "score": 100},
			f.Map{"id": 4, "score": 100},
			f.Map{"id": 5, "score": 99},
		}
		t.Assert(futil.ListItemValuesUnique(listMap, "id"), f.Slice{1, 2, 3, 4, 5})
		t.Assert(futil.ListItemValuesUnique(listMap, "score"), f.Slice{100, 99})
	})
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": 100},
			f.Map{"id": 3, "score": 0},
			f.Map{"id": 4, "score": 100},
			f.Map{"id": 5, "score": 99},
		}
		t.Assert(futil.ListItemValuesUnique(listMap, "id"), f.Slice{1, 2, 3, 4, 5})
		t.Assert(futil.ListItemValuesUnique(listMap, "score"), f.Slice{100, 0, 99})
	})
}

func Test_ListItemValuesUnique_Struct_SubKey(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type Student struct {
			Id    int
			Score float64
		}
		type Class struct {
			Total    int
			Students []Student
		}
		listStruct := f.Slice{
			Class{2, []Student{{1, 1}, {1, 2}}},
			Class{3, []Student{{2, 3}, {2, 4}, {5, 5}}},
			Class{1, []Student{{6, 6}}},
		}
		t.Assert(futil.ListItemValuesUnique(listStruct, "Total"), f.Slice{2, 3, 1})
		t.Assert(futil.ListItemValuesUnique(listStruct, "Students", "Id"), f.Slice{1, 2, 5, 6})
	})
	ftest.C(t, func(t *ftest.T) {
		type Student struct {
			Id    int
			Score float64
		}
		type Class struct {
			Total    int
			Students []*Student
		}
		listStruct := f.Slice{
			&Class{2, []*Student{{1, 1}, {1, 2}}},
			&Class{3, []*Student{{2, 3}, {2, 4}, {5, 5}}},
			&Class{1, []*Student{{6, 6}}},
		}
		t.Assert(futil.ListItemValuesUnique(listStruct, "Total"), f.Slice{2, 3, 1})
		t.Assert(futil.ListItemValuesUnique(listStruct, "Students", "Id"), f.Slice{1, 2, 5, 6})
	})
}

func Test_ListItemValuesUnique_Map_Array_SubKey(t *testing.T) {
	type Scores struct {
		Math    int
		English int
	}
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": 1, "scores": []Scores{{1, 2}, {1, 2}}},
			f.Map{"id": 2, "scores": []Scores{{5, 8}, {5, 8}}},
			f.Map{"id": 3, "scores": []Scores{{9, 10}, {11, 12}}},
		}
		t.Assert(futil.ListItemValuesUnique(listMap, "scores", "Math"), f.Slice{1, 5, 9, 11})
		t.Assert(futil.ListItemValuesUnique(listMap, "scores", "English"), f.Slice{2, 8, 10, 12})
		t.Assert(futil.ListItemValuesUnique(listMap, "scores", "PE"), f.Slice{})
	})
}

func Test_ListItemValuesUnique_Binary_ID(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		listMap := f.List{
			f.Map{"id": []byte{1}, "score": 100},
			f.Map{"id": []byte{2}, "score": 100},
			f.Map{"id": []byte{3}, "score": 100},
			f.Map{"id": []byte{4}, "score": 100},
			f.Map{"id": []byte{4}, "score": 100},
		}
		t.Assert(futil.ListItemValuesUnique(listMap, "id"), f.Slice{[]byte{1}, []byte{2}, []byte{3}, []byte{4}})
		t.Assert(futil.ListItemValuesUnique(listMap, "score"), f.Slice{100})
	})
}

func Test_ListToMapByKey(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := f.List{
			f.Map{"id": 1, "score": 100},
			f.Map{"id": 2, "score": 99},
			f.Map{"id": 3, "score": 98},
			f.Map{"id": 4, "score": 97},
			f.Map{"id": 5, "score": 96},
		}
		idMap := futil.ListToMapByKey(m, "id")
		keys := futil.Keys(idMap)
		t.Assert(len(keys), 5)
		t.AssertIN(1, keys)
		t.AssertIN(2, keys)
		t.AssertIN(3, keys)
		t.AssertIN(4, keys)
		t.AssertIN(5, keys)
	})
}
