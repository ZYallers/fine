package fmap_test

import (
	"testing"
	"time"

	"github.com/ZYallers/fine/container/farray"
	"github.com/ZYallers/fine/container/fmap"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_AnyAnyMap_Var(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var m fmap.AnyAnyMap
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

func Test_AnyAnyMap_Basic(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMap()
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

		m2 := fmap.NewAnyAnyMapFrom(map[interface{}]interface{}{1: 1, 2: "2"})
		t.Assert(m2.Map(), map[interface{}]interface{}{1: 1, 2: "2"})
	})
}

func Test_AnyAnyMap_Set_Fun(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMap()

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

func Test_AnyAnyMap_Batch(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMap()

		m.Sets(map[interface{}]interface{}{1: 1, 2: "2", 3: 3})
		t.Assert(m.Map(), map[interface{}]interface{}{1: 1, 2: "2", 3: 3})
		m.Removes([]interface{}{1, 2})
		t.Assert(m.Map(), map[interface{}]interface{}{3: 3})
	})
}

func Test_AnyAnyMap_Iterator_Deadlock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMapFrom(map[interface{}]interface{}{1: 1, 2: "2", "3": "3", "4": 4}, true)
		m.Iterator(func(k interface{}, _ interface{}) bool {
			if fconv.Int(k)%2 == 0 {
				m.Remove(k)
			}
			return true
		})
		t.Assert(m.Map(), map[interface{}]interface{}{
			1:   1,
			"3": "3",
		})
	})
}

func Test_AnyAnyMap_Iterator(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[interface{}]interface{}{1: 1, 2: "2"}
		m := fmap.NewAnyAnyMapFrom(expect)
		m.Iterator(func(k interface{}, v interface{}) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.Iterator(func(k interface{}, v interface{}) bool {
			i++
			return true
		})
		m.Iterator(func(k interface{}, v interface{}) bool {
			j++
			return false
		})
		t.Assert(i, "2")
		t.Assert(j, 1)
	})
}

func Test_AnyAnyMap_Lock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[interface{}]interface{}{1: 1, 2: "2"}
		m := fmap.NewAnyAnyMapFrom(expect)
		m.LockFunc(func(m map[interface{}]interface{}) {
			t.Assert(m, expect)
		})
		m.RLockFunc(func(m map[interface{}]interface{}) {
			t.Assert(m, expect)
		})
	})
}

func Test_AnyAnyMap_Clone(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// clone 方法是深克隆
		m := fmap.NewAnyAnyMapFrom(map[interface{}]interface{}{1: 1, 2: "2"})

		m_clone := m.Clone()
		m.Remove(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.Keys())

		m_clone.Remove(2)
		// 修改clone map,原 map 不影响
		t.AssertIN(2, m.Keys())
	})
}

func Test_AnyAnyMap_Merge(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewAnyAnyMap()
		m2 := fmap.NewAnyAnyMap()
		m1.Set(1, 1)
		m2.Set(2, "2")
		m1.Merge(m2)
		t.Assert(m1.Map(), map[interface{}]interface{}{1: 1, 2: "2"})
		m3 := fmap.NewAnyAnyMapFrom(nil)
		m3.Merge(m2)
		t.Assert(m3.Map(), m2.Map())
	})
}

func Test_AnyAnyMap_Map(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMap()
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

func Test_AnyAnyMap_MapCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMap()
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

func Test_AnyAnyMap_FilterEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMap()
		m.Set(1, 0)
		m.Set(2, 2)
		t.Assert(m.Get(1), 0)
		t.Assert(m.Get(2), 2)
		m.FilterEmpty()
		t.Assert(m.Get(1), nil)
		t.Assert(m.Get(2), 2)
	})
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMap()
		m.Set(1, 0)
		m.Set("time1", time.Time{})
		m.Set("time2", time.Now())
		t.Assert(m.Get(1), 0)
		t.Assert(m.Get("time1"), time.Time{})
		m.FilterEmpty()
		t.Assert(m.Get(1), nil)
		t.Assert(m.Get("time1"), nil)
		t.AssertNE(m.Get("time2"), nil)
	})
}

func Test_AnyAnyMap_Json(t *testing.T) {
	// Marshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		}
		m1 := fmap.NewAnyAnyMapFrom(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(fconv.Map(data))
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(fconv.Map(data))
		t.AssertNil(err)

		m := fmap.New()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
	ftest.C(t, func(t *ftest.T) {
		data := f.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(fconv.Map(data))
		t.AssertNil(err)

		var m fmap.Map
		err = json.UnmarshalUseNumber(b, &m)
		t.AssertNil(err)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
}

func Test_AnyAnyMap_Pop(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMapFrom(f.MapAnyAny{
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
		t.AssertNil(k3)
		t.AssertNil(v3)
	})
}

func Test_AnyAnyMap_Pops(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMapFrom(f.MapAnyAny{
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

func TestAnyAnyMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *fmap.Map
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

func Test_AnyAnyMap_DeepCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewAnyAnyMapFrom(f.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(m.Size(), 2)

		n := m.DeepCopy().(*fmap.AnyAnyMap)
		n.Set("k1", "val1")
		t.AssertNE(m.Get("k1"), n.Get("k1"))
	})
}

func Test_AnyAnyMap_IsSubOf(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewAnyAnyMapFrom(f.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		})
		m2 := fmap.NewAnyAnyMapFrom(f.MapAnyAny{
			"k2": "v2",
		})
		t.Assert(m1.IsSubOf(m2), false)
		t.Assert(m2.IsSubOf(m1), true)
		t.Assert(m2.IsSubOf(m2), true)
	})
}

func Test_AnyAnyMap_Diff(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewAnyAnyMapFrom(f.MapAnyAny{
			"0": "v0",
			"1": "v1",
			2:   "v2",
			3:   3,
		})
		m2 := fmap.NewAnyAnyMapFrom(f.MapAnyAny{
			"0": "v0",
			2:   "v2",
			3:   "v3",
			4:   "v4",
		})
		addedKeys, removedKeys, updatedKeys := m1.Diff(m2)
		t.Assert(addedKeys, []interface{}{4})
		t.Assert(removedKeys, []interface{}{"1"})
		t.Assert(updatedKeys, []interface{}{3})
	})
}
