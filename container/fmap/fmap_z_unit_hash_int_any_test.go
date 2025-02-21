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

func getAny() interface{} {
	return 123
}

func Test_IntAnyMap_Var(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var m fmap.IntAnyMap
		m.Set(1, 1)

		t.Assert(m.Get(1), 1)
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet(2, "2"), "2")
		t.Assert(m.SetIfNotExist(2, "2"), false)

		t.Assert(m.SetIfNotExist(3, 3), true)

		t.Assert(m.Remove(2), "2")
		t.Assert(m.Contains(2), false)

		t.AssertIN(3, m.Keys())
		t.AssertIN(1, m.Keys())
		t.AssertIN(3, m.Values())
		t.AssertIN(1, m.Values())
		m.Flip()
		t.Assert(m.Map(), map[interface{}]int{1: 1, 3: 3})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)
	})
}

func Test_IntAnyMap_Basic(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMap()
		m.Set(1, 1)

		t.Assert(m.Get(1), 1)
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet(2, "2"), "2")
		t.Assert(m.SetIfNotExist(2, "2"), false)

		t.Assert(m.SetIfNotExist(3, 3), true)

		t.Assert(m.Remove(2), "2")
		t.Assert(m.Contains(2), false)

		t.AssertIN(3, m.Keys())
		t.AssertIN(1, m.Keys())
		t.AssertIN(3, m.Values())
		t.AssertIN(1, m.Values())
		m.Flip()
		t.Assert(m.Map(), map[interface{}]int{1: 1, 3: 3})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := fmap.NewIntAnyMapFrom(map[int]interface{}{1: 1, 2: "2"})
		t.Assert(m2.Map(), map[int]interface{}{1: 1, 2: "2"})
	})

	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMap(true)
		m.Set(1, 1)
		t.Assert(m.Map(), map[int]interface{}{1: 1})
	})
}

func Test_IntAnyMap_Set_Fun(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMap()

		m.GetOrSetFunc(1, getAny)
		m.GetOrSetFuncLock(2, getAny)
		t.Assert(m.Get(1), 123)
		t.Assert(m.Get(2), 123)

		t.Assert(m.SetIfNotExistFunc(1, getAny), false)
		t.Assert(m.SetIfNotExistFunc(3, getAny), true)

		t.Assert(m.SetIfNotExistFuncLock(2, getAny), false)
		t.Assert(m.SetIfNotExistFuncLock(4, getAny), true)
	})
}

func Test_IntAnyMap_Batch(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMap()

		m.Sets(map[int]interface{}{1: 1, 2: "2", 3: 3})
		t.Assert(m.Map(), map[int]interface{}{1: 1, 2: "2", 3: 3})
		m.Removes([]int{1, 2})
		t.Assert(m.Map(), map[int]interface{}{3: 3})
	})
}

func Test_IntAnyMap_Iterator_Deadlock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMapFrom(map[int]interface{}{1: 1, 2: 2, 3: "3", 4: 4}, true)
		m.Iterator(func(k int, _ interface{}) bool {
			if k%2 == 0 {
				m.Remove(k)
			}
			return true
		})
		t.Assert(m.Map(), map[int]interface{}{
			1: 1,
			3: "3",
		})
	})
}

func Test_IntAnyMap_Iterator(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[int]interface{}{1: 1, 2: "2"}
		m := fmap.NewIntAnyMapFrom(expect)
		m.Iterator(func(k int, v interface{}) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.Iterator(func(k int, v interface{}) bool {
			i++
			return true
		})
		m.Iterator(func(k int, v interface{}) bool {
			j++
			return false
		})
		t.Assert(i, "2")
		t.Assert(j, 1)
	})

}

func Test_IntAnyMap_Lock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[int]interface{}{1: 1, 2: "2"}
		m := fmap.NewIntAnyMapFrom(expect)
		m.LockFunc(func(m map[int]interface{}) {
			t.Assert(m, expect)
		})
		m.RLockFunc(func(m map[int]interface{}) {
			t.Assert(m, expect)
		})
	})
}

func Test_IntAnyMap_Clone(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// clone 方法是深克隆
		m := fmap.NewIntAnyMapFrom(map[int]interface{}{1: 1, 2: "2"})

		m_clone := m.Clone()
		m.Remove(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.Keys())

		m_clone.Remove(2)
		// 修改clone map,原 map 不影响
		t.AssertIN(2, m.Keys())
	})
}

func Test_IntAnyMap_Merge(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewIntAnyMap()
		m2 := fmap.NewIntAnyMap()
		m1.Set(1, 1)
		m2.Set(2, "2")
		m1.Merge(m2)
		t.Assert(m1.Map(), map[int]interface{}{1: 1, 2: "2"})
		m3 := fmap.NewIntAnyMapFrom(nil)
		m3.Merge(m2)
		t.Assert(m3.Map(), m2.Map())
	})
}

func Test_IntAnyMap_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMap()
		m.Set(1, 0)
		m.Set(2, 2)
		t.Assert(m.Get(1), 0)
		t.Assert(m.Get(2), 2)
		data := m.Map()
		t.Assert(data[1], 0)
		t.Assert(data[2], 2)
		data[3] = 3
		t.Assert(m.Get(3), 3)
		m.Set(4, 4)
		t.Assert(data[4], 4)
	})
}

func Test_IntAnyMap_MapCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMap()
		m.Set(1, 0)
		m.Set(2, 2)
		t.Assert(m.Get(1), 0)
		t.Assert(m.Get(2), 2)
		data := m.MapCopy()
		t.Assert(data[1], 0)
		t.Assert(data[2], 2)
		data[3] = 3
		t.Assert(m.Get(3), nil)
		m.Set(4, 4)
		t.Assert(data[4], nil)
	})
}

func Test_IntAnyMap_FilterEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMap()
		m.Set(1, 0)
		m.Set(2, 2)
		t.Assert(m.Size(), 2)
		t.Assert(m.Get(1), 0)
		t.Assert(m.Get(2), 2)
		m.FilterEmpty()
		t.Assert(m.Size(), 1)
		t.Assert(m.Get(2), 2)
	})
}

func Test_IntAnyMap_Json(t *testing.T) {
	// Marshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapIntAny{
			1: "v1",
			2: "v2",
		}
		m1 := fmap.NewIntAnyMapFrom(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapIntAny{
			1: "v1",
			2: "v2",
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		m := fmap.NewIntAnyMap()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.Get(1), data[1])
		t.Assert(m.Get(2), data[2])
	})
}

func Test_IntAnyMap_Pop(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMapFrom(f.MapIntAny{
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
		t.AssertNil(v3)
	})
}

func Test_IntAnyMap_Pops(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMapFrom(f.MapIntAny{
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

func TestIntAnyMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *fmap.IntAnyMap
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

func Test_IntAnyMap_DeepCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewIntAnyMapFrom(f.MapIntAny{
			1: "v1",
			2: "v2",
		})
		t.Assert(m.Size(), 2)

		n := m.DeepCopy().(*fmap.IntAnyMap)
		n.Set(1, "val1")
		t.AssertNE(m.Get(1), n.Get(1))
	})
}

func Test_IntAnyMap_IsSubOf(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewIntAnyMapFrom(f.MapIntAny{
			1: "v1",
			2: "v2",
		})
		m2 := fmap.NewIntAnyMapFrom(f.MapIntAny{
			2: "v2",
		})
		t.Assert(m1.IsSubOf(m2), false)
		t.Assert(m2.IsSubOf(m1), true)
		t.Assert(m2.IsSubOf(m2), true)
	})
}

func Test_IntAnyMap_Diff(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewIntAnyMapFrom(f.MapIntAny{
			0: "v0",
			1: "v1",
			2: "v2",
			3: 3,
		})
		m2 := fmap.NewIntAnyMapFrom(f.MapIntAny{
			0: "v0",
			2: "v2",
			3: "v3",
			4: "v4",
		})
		addedKeys, removedKeys, updatedKeys := m1.Diff(m2)
		t.Assert(addedKeys, []int{4})
		t.Assert(removedKeys, []int{1})
		t.Assert(updatedKeys, []int{3})
	})
}
