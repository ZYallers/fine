package fmap_test

import (
	"strconv"
	"testing"

	"github.com/ZYallers/fine/container/farray"
	"github.com/ZYallers/fine/container/fmap"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_StrAnyMap_Var(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var m fmap.StrAnyMap
		m.Set("a", 1)

		t.Assert(m.Get("a"), 1)
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("b", "2"), "2")
		t.Assert(m.SetIfNotExist("b", "2"), false)

		t.Assert(m.SetIfNotExist("c", 3), true)

		t.Assert(m.Remove("b"), "2")
		t.Assert(m.Contains("b"), false)

		t.AssertIN("c", m.Keys())
		t.AssertIN("a", m.Keys())
		t.AssertIN(3, m.Values())
		t.AssertIN(1, m.Values())

		m.Flip()
		t.Assert(m.Map(), map[string]interface{}{"1": "a", "3": "c"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)
	})
}

func Test_StrAnyMap_Basic(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMap()
		m.Set("a", 1)

		t.Assert(m.Get("a"), 1)
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("b", "2"), "2")
		t.Assert(m.SetIfNotExist("b", "2"), false)

		t.Assert(m.SetIfNotExist("c", 3), true)

		t.Assert(m.Remove("b"), "2")
		t.Assert(m.Contains("b"), false)

		t.AssertIN("c", m.Keys())
		t.AssertIN("a", m.Keys())
		t.AssertIN(3, m.Values())
		t.AssertIN(1, m.Values())

		m.Flip()
		t.Assert(m.Map(), map[string]interface{}{"1": "a", "3": "c"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := fmap.NewStrAnyMapFrom(map[string]interface{}{"a": 1, "b": "2"})
		t.Assert(m2.Map(), map[string]interface{}{"a": 1, "b": "2"})
	})
}

func Test_StrAnyMap_Set_Fun(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMap()

		m.GetOrSetFunc("a", getAny)
		m.GetOrSetFuncLock("b", getAny)
		t.Assert(m.Get("a"), 123)
		t.Assert(m.Get("b"), 123)
		t.Assert(m.SetIfNotExistFunc("a", getAny), false)
		t.Assert(m.SetIfNotExistFunc("c", getAny), true)

		t.Assert(m.SetIfNotExistFuncLock("b", getAny), false)
		t.Assert(m.SetIfNotExistFuncLock("d", getAny), true)
	})
}

func Test_StrAnyMap_Batch(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMap()

		m.Sets(map[string]interface{}{"a": 1, "b": "2", "c": 3})
		t.Assert(m.Map(), map[string]interface{}{"a": 1, "b": "2", "c": 3})
		m.Removes([]string{"a", "b"})
		t.Assert(m.Map(), map[string]interface{}{"c": 3})
	})
}

func Test_StrAnyMap_Iterator_Deadlock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMapFrom(map[string]interface{}{"1": "1", "2": "2", "3": "3", "4": "4"}, true)
		m.Iterator(func(k string, _ interface{}) bool {
			kInt, _ := strconv.Atoi(k)
			if kInt%2 == 0 {
				m.Remove(k)
			}
			return true
		})
		t.Assert(m.Map(), map[string]interface{}{
			"1": "1",
			"3": "3",
		})
	})
}

func Test_StrAnyMap_Iterator(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[string]interface{}{"a": true, "b": false}
		m := fmap.NewStrAnyMapFrom(expect)
		m.Iterator(func(k string, v interface{}) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.Iterator(func(k string, v interface{}) bool {
			i++
			return true
		})
		m.Iterator(func(k string, v interface{}) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_StrAnyMap_Lock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[string]interface{}{"a": true, "b": false}

		m := fmap.NewStrAnyMapFrom(expect)
		m.LockFunc(func(m map[string]interface{}) {
			t.Assert(m, expect)
		})
		m.RLockFunc(func(m map[string]interface{}) {
			t.Assert(m, expect)
		})
	})
}

func Test_StrAnyMap_Clone(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// clone 方法是深克隆
		m := fmap.NewStrAnyMapFrom(map[string]interface{}{"a": 1, "b": "2"})

		cloneMap := m.Clone()
		m.Remove("a")
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN("a", cloneMap.Keys())

		cloneMap.Remove("b")
		// 修改clone map,原 map 不影响
		t.AssertIN("b", m.Keys())
	})
}

func Test_StrAnyMap_Merge(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewStrAnyMap()
		m2 := fmap.NewStrAnyMap()
		m1.Set("a", 1)
		m2.Set("b", "2")
		m1.Merge(m2)
		t.Assert(m1.Map(), map[string]interface{}{"a": 1, "b": "2"})

		m3 := fmap.NewStrAnyMapFrom(nil)
		m3.Merge(m2)
		t.Assert(m3.Map(), m2.Map())
	})
}

func Test_StrAnyMap_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMap()
		m.Set("1", 1)
		m.Set("2", 2)
		t.Assert(m.Get("1"), 1)
		t.Assert(m.Get("2"), 2)
		data := m.Map()
		t.Assert(data["1"], 1)
		t.Assert(data["2"], 2)
		data["3"] = 3
		t.Assert(m.Get("3"), 3)
		m.Set("4", 4)
		t.Assert(data["4"], 4)
	})
}

func Test_StrAnyMap_MapCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMap()
		m.Set("1", 1)
		m.Set("2", 2)
		t.Assert(m.Get("1"), 1)
		t.Assert(m.Get("2"), 2)
		data := m.MapCopy()
		t.Assert(data["1"], 1)
		t.Assert(data["2"], 2)
		data["3"] = 3
		t.Assert(m.Get("3"), nil)
		m.Set("4", 4)
		t.Assert(data["4"], nil)
	})
}

func Test_StrAnyMap_FilterEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMap()
		m.Set("1", 0)
		m.Set("2", 2)
		t.Assert(m.Size(), 2)
		t.Assert(m.Get("1"), 0)
		t.Assert(m.Get("2"), 2)
		m.FilterEmpty()
		t.Assert(m.Size(), 1)
		t.Assert(m.Get("2"), 2)
	})
}

func Test_StrAnyMap_Json(t *testing.T) {
	// Marshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapStrAny{
			"k1": "v1",
			"k2": "v2",
		}
		m1 := fmap.NewStrAnyMapFrom(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapStrAny{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		m := fmap.NewStrAnyMap()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
	ftest.C(t, func(t *ftest.T) {
		data := f.MapStrAny{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		var m fmap.StrAnyMap
		err = json.UnmarshalUseNumber(b, &m)
		t.AssertNil(err)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
}

func Test_StrAnyMap_Pop(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMapFrom(f.MapStrAny{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(m.Size(), 2)

		k1, v1 := m.Pop()
		t.AssertIN(k1, f.Slice{"k1", "k2"})
		t.AssertIN(v1, f.Slice{"v1", "v2"})
		t.Assert(m.Size(), 1)
		k2, v2 := m.Pop()
		t.AssertIN(k2, f.Slice{"k1", "k2"})
		t.AssertIN(v2, f.Slice{"v1", "v2"})
		t.Assert(m.Size(), 0)

		t.AssertNE(k1, k2)
		t.AssertNE(v1, v2)

		k3, v3 := m.Pop()
		t.Assert(k3, "")
		t.Assert(v3, "")
	})
}

func Test_StrAnyMap_Pops(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMapFrom(f.MapStrAny{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
		t.Assert(m.Size(), 3)

		kArray := farray.New()
		vArray := farray.New()
		for k, v := range m.Pops(1) {
			t.AssertIN(k, f.Slice{"k1", "k2", "k3"})
			t.AssertIN(v, f.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.Size(), 2)
		for k, v := range m.Pops(2) {
			t.AssertIN(k, f.Slice{"k1", "k2", "k3"})
			t.AssertIN(v, f.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.Size(), 0)

		t.Assert(kArray.Unique().Len(), 3)
		t.Assert(vArray.Unique().Len(), 3)

		v := m.Pops(1)
		t.AssertNil(v)
		v = m.Pops(-1)
		t.AssertNil(v)
	})
}

func TestStrAnyMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *fmap.StrAnyMap
	}
	// JSON
	ftest.C(t, func(t *ftest.T) {
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"map":  []byte(`{"k1":"v1","k2":"v2"}`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get("k1"), "v1")
		t.Assert(v.Map.Get("k2"), "v2")
	})
	// Map
	ftest.C(t, func(t *ftest.T) {
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"map": f.Map{
				"k1": "v1",
				"k2": "v2",
			},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get("k1"), "v1")
		t.Assert(v.Map.Get("k2"), "v2")
	})
}

func Test_StrAnyMap_DeepCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrAnyMapFrom(f.MapStrAny{
			"key1": "val1",
			"key2": "val2",
		})
		t.Assert(m.Size(), 2)

		n := m.DeepCopy().(*fmap.StrAnyMap)
		n.Set("key1", "v1")
		t.AssertNE(m.Get("key1"), n.Get("key1"))
	})
}

func Test_StrAnyMap_IsSubOf(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewStrAnyMapFrom(f.MapStrAny{
			"k1": "v1",
			"k2": "v2",
		})
		m2 := fmap.NewStrAnyMapFrom(f.MapStrAny{
			"k2": "v2",
		})
		t.Assert(m1.IsSubOf(m2), false)
		t.Assert(m2.IsSubOf(m1), true)
		t.Assert(m2.IsSubOf(m2), true)
	})
}

func Test_StrAnyMap_Diff(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewStrAnyMapFrom(f.MapStrAny{
			"0": "v0",
			"1": "v1",
			"2": "v2",
			"3": 3,
		})
		m2 := fmap.NewStrAnyMapFrom(f.MapStrAny{
			"0": "v0",
			"2": "v2",
			"3": "v3",
			"4": "v4",
		})
		addedKeys, removedKeys, updatedKeys := m1.Diff(m2)
		t.Assert(addedKeys, []string{"4"})
		t.Assert(removedKeys, []string{"1"})
		t.Assert(updatedKeys, []string{"3"})
	})
}
