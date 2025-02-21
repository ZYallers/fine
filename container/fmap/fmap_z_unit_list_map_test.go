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

func Test_ListMap_Var(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var m fmap.ListMap
		m.Set("key1", "val1")
		t.Assert(m.Keys(), []interface{}{"key1"})

		t.Assert(m.Get("key1"), "val1")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("key2", "val2"), "val2")
		t.Assert(m.SetIfNotExist("key2", "val2"), false)

		t.Assert(m.SetIfNotExist("key3", "val3"), true)
		t.Assert(m.Remove("key2"), "val2")
		t.Assert(m.Contains("key2"), false)

		t.AssertIN("key3", m.Keys())
		t.AssertIN("key1", m.Keys())
		t.AssertIN("val3", m.Values())
		t.AssertIN("val1", m.Values())

		m.Flip()

		t.Assert(m.Map(), map[interface{}]interface{}{"val3": "key3", "val1": "key1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)
	})
}

func Test_ListMap_Basic(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		m.Set("key1", "val1")
		t.Assert(m.Keys(), []interface{}{"key1"})

		t.Assert(m.Get("key1"), "val1")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("key2", "val2"), "val2")
		t.Assert(m.SetIfNotExist("key2", "val2"), false)

		t.Assert(m.SetIfNotExist("key3", "val3"), true)
		t.Assert(m.Remove("key2"), "val2")
		t.Assert(m.Contains("key2"), false)

		t.AssertIN("key3", m.Keys())
		t.AssertIN("key1", m.Keys())
		t.AssertIN("val3", m.Values())
		t.AssertIN("val1", m.Values())

		m.Flip()

		t.Assert(m.Map(), map[interface{}]interface{}{"val3": "key3", "val1": "key1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := fmap.NewListMapFrom(map[interface{}]interface{}{1: 1, "key1": "val1"})
		t.Assert(m2.Map(), map[interface{}]interface{}{1: 1, "key1": "val1"})
	})
}

func Test_ListMap_Set_Fun(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		m.GetOrSetFunc("fun", getValue)
		m.GetOrSetFuncLock("funlock", getValue)
		t.Assert(m.Get("funlock"), 3)
		t.Assert(m.Get("fun"), 3)
		m.GetOrSetFunc("fun", getValue)
		t.Assert(m.SetIfNotExistFunc("fun", getValue), false)
		t.Assert(m.SetIfNotExistFuncLock("funlock", getValue), false)
	})
}

func Test_ListMap_Batch(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		m.Sets(map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		t.Assert(m.Map(), map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		m.Removes([]interface{}{"key1", 1})
		t.Assert(m.Map(), map[interface{}]interface{}{"key2": "val2", "key3": "val3"})
	})
}

func Test_ListMap_Iterator(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[interface{}]interface{}{1: 1, "key1": "val1"}

		m := fmap.NewListMapFrom(expect)
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
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_ListMap_Clone(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// clone 方法是深克隆
		m := fmap.NewListMapFrom(map[interface{}]interface{}{1: 1, "key1": "val1"})
		m_clone := m.Clone()
		m.Remove(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.Keys())

		m_clone.Remove("key1")
		// 修改clone map,原 map 不影响
		t.AssertIN("key1", m.Keys())
	})
}

func Test_ListMap_Basic_Merge(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.NewListMap()
		m2 := fmap.NewListMap()
		m1.Set("key1", "val1")
		m2.Set("key2", "val2")
		m1.Merge(m2)
		t.Assert(m1.Map(), map[interface{}]interface{}{"key1": "val1", "key2": "val2"})
		m3 := fmap.NewListMapFrom(nil)
		m3.Merge(m2)
		t.Assert(m3.Map(), m2.Map())
	})
}

func Test_ListMap_Order(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		m.Set("k1", "v1")
		m.Set("k2", "v2")
		m.Set("k3", "v3")
		t.Assert(m.Keys(), f.Slice{"k1", "k2", "k3"})
		t.Assert(m.Values(), f.Slice{"v1", "v2", "v3"})
	})
}

func Test_ListMap_FilterEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		m.Set(1, "")
		m.Set(2, "2")
		t.Assert(m.Size(), 2)
		t.Assert(m.Get(2), "2")
		m.FilterEmpty()
		t.Assert(m.Size(), 1)
		t.Assert(m.Get(2), "2")
	})
}

func Test_ListMap_Json(t *testing.T) {
	// Marshal
	ftest.C(t, func(t *ftest.T) {
		data := f.MapAnyAny{
			"k1": "v1",
		}
		m1 := fmap.NewListMapFrom(data)
		b1, err1 := json.Marshal(m1)
		t.AssertNil(err1)
		b2, err2 := json.Marshal(fconv.Map(data))
		t.AssertNil(err2)
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

		m := fmap.NewListMap()
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

		var m fmap.ListMap
		err = json.UnmarshalUseNumber(b, &m)
		t.AssertNil(err)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
}

func Test_ListMap_Json_Sequence(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		for i := 'z'; i >= 'a'; i-- {
			m.Set(string(i), i)
		}
		b, err := json.Marshal(m)
		t.AssertNil(err)
		t.Assert(b, `{"z":122,"y":121,"x":120,"w":119,"v":118,"u":117,"t":116,"s":115,"r":114,"q":113,"p":112,"o":111,"n":110,"m":109,"l":108,"k":107,"j":106,"i":105,"h":104,"g":103,"f":102,"e":101,"d":100,"c":99,"b":98,"a":97}`)
	})
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		for i := 'a'; i <= 'z'; i++ {
			m.Set(string(i), i)
		}
		b, err := json.Marshal(m)
		t.AssertNil(err)
		t.Assert(b, `{"a":97,"b":98,"c":99,"d":100,"e":101,"f":102,"g":103,"h":104,"i":105,"j":106,"k":107,"l":108,"m":109,"n":110,"o":111,"p":112,"q":113,"r":114,"s":115,"t":116,"u":117,"v":118,"w":119,"x":120,"y":121,"z":122}`)
	})
}

func Test_ListMap_Pop(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMapFrom(f.MapAnyAny{
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

func Test_ListMap_Pops(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMapFrom(f.MapAnyAny{
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

func TestListMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *fmap.ListMap
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
		t.Assert(v.Map.Get("1"), "v1")
		t.Assert(v.Map.Get("2"), "v2")
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
		t.Assert(v.Map.Get("1"), "v1")
		t.Assert(v.Map.Get("2"), "v2")
	})
}

func TestListMap_String(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		m.Set(1, "")
		m.Set(2, "2")
		t.Assert(m.String(), "{\"1\":\"\",\"2\":\"2\"}")

		m1 := fmap.NewListMapFrom(nil)
		t.Assert(m1.String(), "{}")
	})
}

func TestListMap_MarshalJSON(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		m.Set(1, "")
		m.Set(2, "2")
		res, err := m.MarshalJSON()
		t.Assert(res, []byte("{\"1\":\"\",\"2\":\"2\"}"))
		t.AssertNil(err)

		m1 := fmap.NewListMapFrom(nil)
		res, err = m1.MarshalJSON()
		t.Assert(res, []byte("{}"))
		t.AssertNil(err)
	})
}

func TestListMap_DeepCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.NewListMap()
		m.Set(1, "1")
		m.Set(2, "2")
		t.Assert(m.Size(), 2)

		n := m.DeepCopy().(*fmap.ListMap)
		n.Set(1, "val1")
		t.AssertNE(m.Get(1), n.Get(1))
	})
}
