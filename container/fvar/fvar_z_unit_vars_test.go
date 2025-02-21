package fvar_test

import (
	"testing"

	"github.com/ZYallers/fine/container/fvar"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
)

func TestVars(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var vs = fvar.Vars{
			fvar.New(1),
			fvar.New(2),
			fvar.New(3),
		}
		t.AssertEQ(vs.Strings(), []string{"1", "2", "3"})
		t.AssertEQ(vs.Interfaces(), []interface{}{1, 2, 3})
		t.AssertEQ(vs.Float32s(), []float32{1, 2, 3})
		t.AssertEQ(vs.Float64s(), []float64{1, 2, 3})
		t.AssertEQ(vs.Ints(), []int{1, 2, 3})
		t.AssertEQ(vs.Int8s(), []int8{1, 2, 3})
		t.AssertEQ(vs.Int16s(), []int16{1, 2, 3})
		t.AssertEQ(vs.Int32s(), []int32{1, 2, 3})
		t.AssertEQ(vs.Int64s(), []int64{1, 2, 3})
		t.AssertEQ(vs.Uints(), []uint{1, 2, 3})
		t.AssertEQ(vs.Uint8s(), []uint8{1, 2, 3})
		t.AssertEQ(vs.Uint16s(), []uint16{1, 2, 3})
		t.AssertEQ(vs.Uint32s(), []uint32{1, 2, 3})
		t.AssertEQ(vs.Uint64s(), []uint64{1, 2, 3})
	})
}

func TestVars_Scan(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type User struct {
			Id   int
			Name string
		}
		var vs = fvar.Vars{
			fvar.New(f.Map{"id": 1, "name": "john"}),
			fvar.New(f.Map{"id": 2, "name": "smith"}),
		}
		var users []User
		err := vs.Scan(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 1)
		t.Assert(users[0].Name, "john")
		t.Assert(users[1].Id, 2)
		t.Assert(users[1].Name, "smith")
	})
}
