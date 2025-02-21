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

func Test_StrStrMap_Var(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var m fmap.StrStrMap
		m.Set("a", "a")

		t.Assert(m.Get("a"), "a")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("b", "b"), "b")
		t.Assert(m.SetIfNotExist("b", "b"), false)

		t.Assert(m.SetIfNotExist("c", "c"), true)

		t.Assert(m.Remove("b"), "b")
		t.Assert(m.Contains("b"), false)

		t.AssertIN("c", m.Keys())
		t.AssertIN("a", m.Keys())
		t.AssertIN("a", m.Values())
		t.AssertIN("c", m.Values())

		m.Flip()

		t.Assert(m.Map(), map[string]string{"a": "a", "c": "c"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)
	})
}

func Test_StrStrMap_Basic(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMap()
		m.Set("a", "a")

		t.Assert(m.Get("a"), "a")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("b", "b"), "b")
		t.Assert(m.SetIfNotExist("b", "b"), false)

		t.Assert(m.SetIfNotExist("c", "c"), true)

		t.Assert(m.Remove("b"), "b")
		t.Assert(m.Contains("b"), false)

		t.AssertIN("c", m.Keys())
		t.AssertIN("a", m.Keys())
		t.AssertIN("a", m.Values())
		t.AssertIN("c", m.Values())

		m.Flip()

		t.Assert(m.Map(), map[string]string{"a": "a", "c": "c"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := fmap.NewStrStrMapFrom(map[string]string{"a": "a", "b": "b"})
		t.Assert(m2.Map(), map[string]string{"a": "a", "b": "b"})
	})
}

func Test_StrStrMap_Set_Fun(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMap()

		m.GetOrSetFunc("a", getStr)
		m.GetOrSetFuncLock("b", getStr)
		t.Assert(m.Get("a"), "z")
		t.Assert(m.Get("b"), "z")
		t.Assert(m.SetIfNotExistFunc("a", getStr), false)
		t.Assert(m.SetIfNotExistFunc("c", getStr), true)

		t.Assert(m.SetIfNotExistFuncLock("b", getStr), false)
		t.Assert(m.SetIfNotExistFuncLock("d", getStr), true)
	})

	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMapFrom(nil)

		t.Assert(m.GetOrSetFuncLock("b", getStr), "z")
	})
}

func Test_StrStrMap_Batch(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMap()

		m.Sets(map[string]string{"a": "a", "b": "b", "c": "c"})
		t.Assert(m.Map(), map[string]string{"a": "a", "b": "b", "c": "c"})
		m.Removes([]string{"a", "b"})
		t.Assert(m.Map(), map[string]string{"c": "c"})
	})
}

func Test_StrStrMap_Iterator_Deadlock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMapFrom(map[string]string{"1": "1", "2": "2", "3": "3", "4": "4"}, true)
		m.Iterator(func(k string, _ string) bool {
			kInt, _ := strconv.Atoi(k)
			if kInt%2 == 0 {
				m.Remove(k)
			}
			return true
		})
		t.Assert(m.Size(), 2)
	})
}

func Test_StrStrMap_Iterator(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[string]string{"a": "a", "b": "b"}
		m := fmap.NewStrStrMapFrom(expect)
		m.Iterator(func(k string, v string) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.Iterator(func(k string, v string) bool {
			i++
			return true
		})
		m.Iterator(func(k string, v string) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_StrStrMap_Lock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[string]string{"a": "a", "b": "b"}

		m := fmap.NewStrStrMapFrom(expect)
		m.LockFunc(func(m map[string]string) {
			t.Assert(m, expect)
		})
		m.RLockFunc(func(m map[string]string) {
			t.Assert(m, expect)
		})
	})
}

func Test_StrStrMap_Clone(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// clone 方法是深克隆
		m := fmap.NewStrStrMapFrom(map[string]string{"a": "a", "b": "b", "c": "c"})

		m_clone := m.Clone()
		m.Remove("a")
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN("a", m_clone.Keys())

		m_clone.Remove("b")
		// 修改clone map,原 map 不影响
		t.AssertIN("b", m.Keys())
	})
}

func Test_StrStrMap_Merge(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewStrStrMap()
		m2 := fmap.NewStrStrMap()
		m1.Set("a", "a")
		m2.Set("b", "b")
		m1.Merge(m2)
		t.Assert(m1.Map(), map[string]string{"a": "a", "b": "b"})
		m3 := fmap.NewStrStrMapFrom(nil)
		m3.Merge(m2)
		t.Assert(m3.Map(), m2.Map())
	})
}

func Test_StrStrMap_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMap()
		m.Set("1", "1")
		m.Set("2", "2")
		t.Assert(m.Get("1"), "1")
		t.Assert(m.Get("2"), "2")
		data := m.Map()
		t.Assert(data["1"], "1")
		t.Assert(data["2"], "2")
		data["3"] = "3"
		t.Assert(m.Get("3"), "3")
		m.Set("4", "4")
		t.Assert(data["4"], "4")
	})
}

func Test_StrStrMap_MapCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMap()
		m.Set("1", "1")
		m.Set("2", "2")
		t.Assert(m.Get("1"), "1")
		t.Assert(m.Get("2"), "2")
		data := m.MapCopy()
		t.Assert(data["1"], "1")
		t.Assert(data["2"], "2")
		data["3"] = "3"
		t.Assert(m.Get("3"), "")
		m.Set("4", "4")
		t.Assert(data["4"], "")
	})
}

func Test_StrStrMap_FilterEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMap()
		m.Set("1", "")
		m.Set("2", "2")
		t.Assert(m.Size(), 2)
		t.Assert(m.Get("1"), "")
		t.Assert(m.Get("2"), "2")
		m.FilterEmpty()
		t.Assert(m.Size(), 1)
		t.Assert(m.Get("2"), "2")
	})
}

func Test_StrStrMap_Json(t *testing.T) {
	// Marshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m1 := fmap.NewStrStrMapFrom(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		m := fmap.NewStrStrMap()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
	ftest.C(t, func(t *ftest.T) {
		data := f.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		var m fmap.StrStrMap
		err = json.UnmarshalUseNumber(b, &m)
		t.AssertNil(err)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
}

func Test_StrStrMap_Pop(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMapFrom(f.MapStrStr{
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

func Test_StrStrMap_Pops(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMapFrom(f.MapStrStr{
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

func TestStrStrMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *fmap.StrStrMap
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

func Test_StrStrMap_DeepCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewStrStrMapFrom(f.MapStrStr{
			"key1": "val1",
			"key2": "val2",
		})
		t.Assert(m.Size(), 2)

		n := m.DeepCopy().(*fmap.StrStrMap)
		n.Set("key1", "v1")
		t.AssertNE(m.Get("key1"), n.Get("key1"))
	})
}

func Test_StrStrMap_IsSubOf(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewStrStrMapFrom(f.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		})
		m2 := fmap.NewStrStrMapFrom(f.MapStrStr{
			"k2": "v2",
		})
		t.Assert(m1.IsSubOf(m2), false)
		t.Assert(m2.IsSubOf(m1), true)
		t.Assert(m2.IsSubOf(m2), true)
	})
}

func Test_StrStrMap_Diff(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewStrStrMapFrom(f.MapStrStr{
			"0": "0",
			"1": "1",
			"2": "2",
			"3": "3",
		})
		m2 := fmap.NewStrStrMapFrom(f.MapStrStr{
			"0": "0",
			"2": "2",
			"3": "31",
			"4": "4",
		})
		addedKeys, removedKeys, updatedKeys := m1.Diff(m2)
		t.Assert(addedKeys, []string{"4"})
		t.Assert(removedKeys, []string{"1"})
		t.Assert(updatedKeys, []string{"3"})
	})
}
