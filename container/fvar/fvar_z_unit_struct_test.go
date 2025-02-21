package fvar_test

import (
	"testing"

	"github.com/ZYallers/fine/container/fvar"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func TestVar_Struct(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type StTest struct {
			Test int
		}

		Kv := make(map[string]int, 1)
		Kv["Test"] = 100

		testObj := &StTest{}

		objOne := fvar.New(Kv, true)

		objOne.Struct(testObj)

		t.Assert(testObj.Test, Kv["Test"])
	})
	ftest.C(t, func(t *ftest.T) {
		type StTest struct {
			Test int8
		}
		o := &StTest{}
		v := fvar.New(f.Slice{"Test", "-25"})
		v.Struct(o)
		t.Assert(o.Test, -25)
	})
}

func TestVar_Var_Attribute_Struct(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type User struct {
			Uid  int
			Name string
		}
		user := new(User)
		err := fconv.Struct(
			f.Map{
				"uid":  1,
				"name": "john",
			}, user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, "john")
	})
	ftest.C(t, func(t *ftest.T) {
		type User struct {
			Uid  int
			Name string
		}
		var user *User
		err := fconv.Struct(
			f.Map{
				"uid":  1,
				"name": "john",
			}, &user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, "john")
	})
}

func TestVar_structs(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		paramsArray := []f.Map{}
		params1 := f.Map{
			"uid":  1,
			"name": "golang",
		}
		params2 := f.Map{
			"uid":  2,
			"name": "java",
		}
		paramsArray = append(paramsArray, params1, params2)
		v := fvar.New(paramsArray)
		type target struct {
			Uid  int
			Name string
		}
		var tar []target
		err := v.Structs(&tar)
		t.AssertNil(err)
		t.Assert(len(tar), 2)
		t.Assert(tar[0].Uid, 1)
		t.Assert(tar[0].Name, "golang")
		t.Assert(tar[1].Uid, 2)
		t.Assert(tar[1].Name, "java")
	})

}
