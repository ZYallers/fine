package fmap_test

import (
	"testing"

	"github.com/ZYallers/fine/container/fmap"
	"github.com/ZYallers/fine/test/ftest"
)

func getValue() interface{} {
	return 3
}

func Test_Map_Var(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var m fmap.Map
		m.Set(1, 11)
		t.Assert(m.Get(1), 11)
	})
	ftest.C(t, func(t *ftest.T) {
		var m fmap.IntAnyMap
		m.Set(1, 11)
		t.Assert(m.Get(1), 11)
	})
	ftest.C(t, func(t *ftest.T) {
		var m fmap.IntIntMap
		m.Set(1, 11)
		t.Assert(m.Get(1), 11)
	})
	ftest.C(t, func(t *ftest.T) {
		var m fmap.IntStrMap
		m.Set(1, "11")
		t.Assert(m.Get(1), "11")
	})
	ftest.C(t, func(t *ftest.T) {
		var m fmap.StrAnyMap
		m.Set("1", "11")
		t.Assert(m.Get("1"), "11")
	})
	ftest.C(t, func(t *ftest.T) {
		var m fmap.StrStrMap
		m.Set("1", "11")
		t.Assert(m.Get("1"), "11")
	})
	ftest.C(t, func(t *ftest.T) {
		var m fmap.StrIntMap
		m.Set("1", 11)
		t.Assert(m.Get("1"), 11)
	})
	ftest.C(t, func(t *ftest.T) {
		var m fmap.ListMap
		m.Set("1", 11)
		t.Assert(m.Get("1"), 11)
	})
}

func Test_Map_Basic(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.New()
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

		m2 := fmap.NewFrom(map[interface{}]interface{}{1: 1, "key1": "val1"})
		t.Assert(m2.Map(), map[interface{}]interface{}{1: 1, "key1": "val1"})
	})
}

func Test_Map_Set_Fun(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.New()
		m.GetOrSetFunc("fun", getValue)
		m.GetOrSetFuncLock("funlock", getValue)
		t.Assert(m.Get("funlock"), 3)
		t.Assert(m.Get("fun"), 3)
		m.GetOrSetFunc("fun", getValue)
		t.Assert(m.SetIfNotExistFunc("fun", getValue), false)
		t.Assert(m.SetIfNotExistFuncLock("funlock", getValue), false)
	})
}

func Test_Map_Batch(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m := fmap.New()
		m.Sets(map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		t.Assert(m.Map(), map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		m.Removes([]interface{}{"key1", 1})
		t.Assert(m.Map(), map[interface{}]interface{}{"key2": "val2", "key3": "val3"})
	})
}

func Test_Map_Iterator(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[interface{}]interface{}{1: 1, "key1": "val1"}

		m := fmap.NewFrom(expect)
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

func Test_Map_Lock(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		expect := map[interface{}]interface{}{1: 1, "key1": "val1"}
		m := fmap.NewFrom(expect)
		m.LockFunc(func(m map[interface{}]interface{}) {
			t.Assert(m, expect)
		})
		m.RLockFunc(func(m map[interface{}]interface{}) {
			t.Assert(m, expect)
		})
	})
}

func Test_Map_Clone(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// clone 方法是深克隆
		m := fmap.NewFrom(map[interface{}]interface{}{1: 1, "key1": "val1"})
		m_clone := m.Clone()
		m.Remove(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.Keys())

		m_clone.Remove("key1")
		// 修改clone map,原 map 不影响
		t.AssertIN("key1", m.Keys())
	})
}

func Test_Map_Basic_Merge(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		m1 := fmap.New()
		m2 := fmap.New()
		m1.Set("key1", "val1")
		m2.Set("key2", "val2")
		m1.Merge(m2)
		t.Assert(m1.Map(), map[interface{}]interface{}{"key1": "val1", "key2": "val2"})
	})
}
