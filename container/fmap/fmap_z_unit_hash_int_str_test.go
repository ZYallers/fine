package fmap_test

import (
	"testing"

	"github.com/ZYallers/fine/container/farray"
	"github.com/ZYallers/fine/container/fmap"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func getStr() string {
	return "z"
}

func Test_IntStrMap_Var(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var m fmap.IntStrMap
		m.Set(1, "a")

		t.Assert(m.Get(1), "a")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet(2, "b"), "b")
		t.Assert(m.SetIfNotExist(2, "b"), false)

		t.Assert(m.SetIfNotExist(3, "c"), true)

		t.Assert(m.Remove(2), "b")
		t.Assert(m.Contains(2), false)

		t.AssertIN(3, m.Keys())
		t.AssertIN(1, m.Keys())
		t.AssertIN("a", m.Values())
		t.AssertIN("c", m.Values())

		m_f := fmap.NewIntStrMap()
		m_f.Set(1, "2")
		m_f.Flip()
		t.Assert(m_f.Map(), map[int]string{2: "1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)
	})
}

func Test_IntStrMap_Basic(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMap()
		m.Set(1, "a")

		t.Assert(m.Get(1), "a")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet(1, "a"), "a")
		t.Assert(m.GetOrSet(2, "b"), "b")
		t.Assert(m.SetIfNotExist(2, "b"), false)

		t.Assert(m.SetIfNotExist(3, "c"), true)

		t.Assert(m.Remove(2), "b")
		t.Assert(m.Contains(2), false)

		t.AssertIN(3, m.Keys())
		t.AssertIN(1, m.Keys())
		t.AssertIN("a", m.Values())
		t.AssertIN("c", m.Values())

		// 反转之后不成为以下 map,flip 操作只是翻转原 map
		// t.Assert(m.Map(), map[string]int{"a": 1, "c": 3})
		m_f := fmap.NewIntStrMap()
		m_f.Set(1, "2")
		m_f.Flip()
		t.Assert(m_f.Map(), map[int]string{2: "1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := fmap.NewIntStrMapFrom(map[int]string{1: "a", 2: "b"})
		t.Assert(m2.Map(), map[int]string{1: "a", 2: "b"})
	})

	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMap(true)
		m.Set(1, "val1")
		t.Assert(m.Map(), map[int]string{1: "val1"})
	})
}

func TestIntStrMap_MapStrAny(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMap()
		m.GetOrSetFunc(1, getStr)
		m.GetOrSetFuncLock(2, getStr)
		t.Assert(m.MapStrAny(), f.MapStrAny{"1": "z", "2": "z"})
	})
}

func TestIntStrMap_Sets(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMapFrom(nil)
		m.Sets(f.MapIntStr{1: "z", 2: "z"})
		t.Assert(len(m.Map()), 2)
	})
}

func Test_IntStrMap_Set_Fun(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMap()
		m.GetOrSetFunc(1, getStr)
		m.GetOrSetFuncLock(2, getStr)
		t.Assert(m.GetOrSetFunc(1, getStr), "z")
		t.Assert(m.GetOrSetFuncLock(2, getStr), "z")
		t.Assert(m.Get(1), "z")
		t.Assert(m.Get(2), "z")
		t.Assert(m.SetIfNotExistFunc(1, getStr), false)
		t.Assert(m.SetIfNotExistFunc(3, getStr), true)

		t.Assert(m.SetIfNotExistFuncLock(2, getStr), false)
		t.Assert(m.SetIfNotExistFuncLock(4, getStr), true)
	})

	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMapFrom(nil)
		t.Assert(m.GetOrSetFuncLock(1, getStr), "z")
	})

	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMapFrom(nil)
		t.Assert(m.SetIfNotExistFuncLock(1, getStr), true)
	})
}

func Test_IntStrMap_Batch(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMap()
		m.Sets(map[int]string{1: "a", 2: "b", 3: "c"})
		t.Assert(m.Map(), map[int]string{1: "a", 2: "b", 3: "c"})
		m.Removes([]int{1, 2})
		t.Assert(m.Map(), map[int]interface{}{3: "c"})
	})
}

func Test_IntStrMap_Iterator_Deadlock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMapFrom(map[int]string{1: "1", 2: "2", 3: "3", 4: "4"}, true)
		m.Iterator(func(k int, _ string) bool {
			if k%2 == 0 {
				m.Remove(k)
			}
			return true
		})
		t.Assert(m.Map(), map[int]string{
			1: "1",
			3: "3",
		})
	})
}

func Test_IntStrMap_Iterator(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[int]string{1: "a", 2: "b"}
		m := fmap.NewIntStrMapFrom(expect)
		m.Iterator(func(k int, v string) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.Iterator(func(k int, v string) bool {
			i++
			return true
		})
		m.Iterator(func(k int, v string) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_IntStrMap_Lock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[int]string{1: "a", 2: "b", 3: "c"}
		m := fmap.NewIntStrMapFrom(expect)
		m.LockFunc(func(m map[int]string) {
			t.Assert(m, expect)
		})
		m.RLockFunc(func(m map[int]string) {
			t.Assert(m, expect)
		})
	})
}

func Test_IntStrMap_Clone(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// clone 方法是深克隆
		m := fmap.NewIntStrMapFrom(map[int]string{1: "a", 2: "b", 3: "c"})

		m_clone := m.Clone()
		m.Remove(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.Keys())

		m_clone.Remove(2)
		// 修改clone map,原 map 不影响
		t.AssertIN(2, m.Keys())
	})
}

func Test_IntStrMap_Merge(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewIntStrMap()
		m2 := fmap.NewIntStrMap()
		m1.Set(1, "a")
		m2.Set(2, "b")
		m1.Merge(m2)
		t.Assert(m1.Map(), map[int]string{1: "a", 2: "b"})

		m3 := fmap.NewIntStrMapFrom(nil)
		m3.Merge(m2)
		t.Assert(m3.Map(), m2.Map())
	})
}

func Test_IntStrMap_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMap()
		m.Set(1, "0")
		m.Set(2, "2")
		t.Assert(m.Get(1), "0")
		t.Assert(m.Get(2), "2")
		data := m.Map()
		t.Assert(data[1], "0")
		t.Assert(data[2], "2")
		data[3] = "3"
		t.Assert(m.Get(3), "3")
		m.Set(4, "4")
		t.Assert(data[4], "4")
	})
}

func Test_IntStrMap_MapCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMap()
		m.Set(1, "0")
		m.Set(2, "2")
		t.Assert(m.Get(1), "0")
		t.Assert(m.Get(2), "2")
		data := m.MapCopy()
		t.Assert(data[1], "0")
		t.Assert(data[2], "2")
		data[3] = "3"
		t.Assert(m.Get(3), "")
		m.Set(4, "4")
		t.Assert(data[4], "")
	})
}

func Test_IntStrMap_FilterEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMap()
		m.Set(1, "")
		m.Set(2, "2")
		t.Assert(m.Size(), 2)
		t.Assert(m.Get(2), "2")
		m.FilterEmpty()
		t.Assert(m.Size(), 1)
		t.Assert(m.Get(2), "2")
	})
}

func Test_IntStrMap_Json(t *testing.T) {
	// Marshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapIntStr{
			1: "v1",
			2: "v2",
		}
		m1 := fmap.NewIntStrMapFrom(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapIntStr{
			1: "v1",
			2: "v2",
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		m := fmap.NewIntStrMap()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.Get(1), data[1])
		t.Assert(m.Get(2), data[2])
	})
}

func Test_IntStrMap_Pop(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMapFrom(f.MapIntStr{
			1: "v1",
			2: "v2",
		})
		t.Assert(m.Size(), 2)

		k1, v1 := m.Pop()
		t.AssertIN(k1, f.Slice{1, 2})
		t.AssertIN(v1, f.Slice{"v1", "v2"})
		t.Assert(m.Size(), 1)
		k2, v2 := m.Pop()
		t.AssertIN(k2, f.Slice{1, 2})
		t.AssertIN(v2, f.Slice{"v1", "v2"})
		t.Assert(m.Size(), 0)

		t.AssertNE(k1, k2)
		t.AssertNE(v1, v2)

		k3, v3 := m.Pop()
		t.Assert(k3, 0)
		t.Assert(v3, "")
	})
}

func Test_IntStrMap_Pops(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMapFrom(f.MapIntStr{
			1: "v1",
			2: "v2",
			3: "v3",
		})
		t.Assert(m.Size(), 3)

		kArray := farray.New()
		vArray := farray.New()
		for k, v := range m.Pops(1) {
			t.AssertIN(k, f.Slice{1, 2, 3})
			t.AssertIN(v, f.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.Size(), 2)
		for k, v := range m.Pops(2) {
			t.AssertIN(k, f.Slice{1, 2, 3})
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

func TestIntStrMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *fmap.IntStrMap
	}
	// JSON
	ftest.C(t, func(t *ftest.T) {
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"map":  []byte(`{"1":"v1","2":"v2"}`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get(1), "v1")
		t.Assert(v.Map.Get(2), "v2")
	})
	// Map
	ftest.C(t, func(t *ftest.T) {
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"map": f.MapIntAny{
				1: "v1",
				2: "v2",
			},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get(1), "v1")
		t.Assert(v.Map.Get(2), "v2")
	})
}

func TestIntStrMap_Replace(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMapFrom(f.MapIntStr{
			1: "v1",
			2: "v2",
			3: "v3",
		})

		t.Assert(m.Get(1), "v1")
		t.Assert(m.Get(2), "v2")
		t.Assert(m.Get(3), "v3")

		m.Replace(f.MapIntStr{
			1: "v2",
			2: "v3",
			3: "v1",
		})

		t.Assert(m.Get(1), "v2")
		t.Assert(m.Get(2), "v3")
		t.Assert(m.Get(3), "v1")
	})
}

func TestIntStrMap_String(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMapFrom(f.MapIntStr{
			1: "v1",
			2: "v2",
			3: "v3",
		})
		t.Assert(m.String(), "{\"1\":\"v1\",\"2\":\"v2\",\"3\":\"v3\"}")

		m = nil
		t.Assert(len(m.String()), 0)
	})
}

func Test_IntStrMap_DeepCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntStrMapFrom(f.MapIntStr{
			1: "val1",
			2: "val2",
		})
		t.Assert(m.Size(), 2)

		n := m.DeepCopy().(*fmap.IntStrMap)
		n.Set(1, "v1")
		t.AssertNE(m.Get(1), n.Get(1))
	})
}

func Test_IntStrMap_IsSubOf(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewIntStrMapFrom(f.MapIntStr{
			1: "v1",
			2: "v2",
		})
		m2 := fmap.NewIntStrMapFrom(f.MapIntStr{
			2: "v2",
		})
		t.Assert(m1.IsSubOf(m2), false)
		t.Assert(m2.IsSubOf(m1), true)
		t.Assert(m2.IsSubOf(m2), true)
	})
}

func Test_IntStrMap_Diff(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewIntStrMapFrom(f.MapIntStr{
			0: "0",
			1: "1",
			2: "2",
			3: "3",
		})
		m2 := fmap.NewIntStrMapFrom(f.MapIntStr{
			0: "0",
			2: "2",
			3: "31",
			4: "4",
		})
		addedKeys, removedKeys, updatedKeys := m1.Diff(m2)
		t.Assert(addedKeys, []int{4})
		t.Assert(removedKeys, []int{1})
		t.Assert(updatedKeys, []int{3})
	})
}
