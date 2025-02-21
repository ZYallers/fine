package fvar_test

import (
	"testing"

	"github.com/ZYallers/fine/container/fvar"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
)

func TestVar_IsNil(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(0).IsNil(), false)
		t.Assert(f.NewVar(nil).IsNil(), true)
		t.Assert(f.NewVar(f.Map{}).IsNil(), false)
		t.Assert(f.NewVar(f.Slice{}).IsNil(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(1).IsNil(), false)
		t.Assert(f.NewVar(0.1).IsNil(), false)
		t.Assert(f.NewVar(f.Map{"k": "v"}).IsNil(), false)
		t.Assert(f.NewVar(f.Slice{0}).IsNil(), false)
	})
}

func TestVar_IsEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(0).IsEmpty(), true)
		t.Assert(f.NewVar(nil).IsEmpty(), true)
		t.Assert(f.NewVar(f.Map{}).IsEmpty(), true)
		t.Assert(f.NewVar(f.Slice{}).IsEmpty(), true)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(1).IsEmpty(), false)
		t.Assert(f.NewVar(0.1).IsEmpty(), false)
		t.Assert(f.NewVar(f.Map{"k": "v"}).IsEmpty(), false)
		t.Assert(f.NewVar(f.Slice{0}).IsEmpty(), false)
	})
}

func TestVar_IsInt(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(0).IsInt(), true)
		t.Assert(f.NewVar(nil).IsInt(), false)
		t.Assert(f.NewVar(f.Map{}).IsInt(), false)
		t.Assert(f.NewVar(f.Slice{}).IsInt(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(1).IsInt(), true)
		t.Assert(f.NewVar(-1).IsInt(), true)
		t.Assert(f.NewVar(0.1).IsInt(), false)
		t.Assert(f.NewVar(f.Map{"k": "v"}).IsInt(), false)
		t.Assert(f.NewVar(f.Slice{0}).IsInt(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(int8(1)).IsInt(), true)
		t.Assert(f.NewVar(uint8(1)).IsInt(), false)
	})
}

func TestVar_IsUint(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(0).IsUint(), false)
		t.Assert(f.NewVar(nil).IsUint(), false)
		t.Assert(f.NewVar(f.Map{}).IsUint(), false)
		t.Assert(f.NewVar(f.Slice{}).IsUint(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(1).IsUint(), false)
		t.Assert(f.NewVar(-1).IsUint(), false)
		t.Assert(f.NewVar(0.1).IsUint(), false)
		t.Assert(f.NewVar(f.Map{"k": "v"}).IsUint(), false)
		t.Assert(f.NewVar(f.Slice{0}).IsUint(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(int8(1)).IsUint(), false)
		t.Assert(f.NewVar(uint8(1)).IsUint(), true)
	})
}

func TestVar_IsFloat(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(0).IsFloat(), false)
		t.Assert(f.NewVar(nil).IsFloat(), false)
		t.Assert(f.NewVar(f.Map{}).IsFloat(), false)
		t.Assert(f.NewVar(f.Slice{}).IsFloat(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(1).IsFloat(), false)
		t.Assert(f.NewVar(-1).IsFloat(), false)
		t.Assert(f.NewVar(0.1).IsFloat(), true)
		t.Assert(f.NewVar(float64(1)).IsFloat(), true)
		t.Assert(f.NewVar(f.Map{"k": "v"}).IsFloat(), false)
		t.Assert(f.NewVar(f.Slice{0}).IsFloat(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(int8(1)).IsFloat(), false)
		t.Assert(f.NewVar(uint8(1)).IsFloat(), false)
	})
}

func TestVar_IsSlice(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(0).IsSlice(), false)
		t.Assert(f.NewVar(nil).IsSlice(), false)
		t.Assert(f.NewVar(f.Map{}).IsSlice(), false)
		t.Assert(f.NewVar(f.Slice{}).IsSlice(), true)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(1).IsSlice(), false)
		t.Assert(f.NewVar(-1).IsSlice(), false)
		t.Assert(f.NewVar(0.1).IsSlice(), false)
		t.Assert(f.NewVar(float64(1)).IsSlice(), false)
		t.Assert(f.NewVar(f.Map{"k": "v"}).IsSlice(), false)
		t.Assert(f.NewVar(f.Slice{0}).IsSlice(), true)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(int8(1)).IsSlice(), false)
		t.Assert(f.NewVar(uint8(1)).IsSlice(), false)
	})
}

func TestVar_IsMap(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(0).IsMap(), false)
		t.Assert(f.NewVar(nil).IsMap(), false)
		t.Assert(f.NewVar(f.Map{}).IsMap(), true)
		t.Assert(f.NewVar(f.Slice{}).IsMap(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(1).IsMap(), false)
		t.Assert(f.NewVar(-1).IsMap(), false)
		t.Assert(f.NewVar(0.1).IsMap(), false)
		t.Assert(f.NewVar(float64(1)).IsMap(), false)
		t.Assert(f.NewVar(f.Map{"k": "v"}).IsMap(), true)
		t.Assert(f.NewVar(f.Slice{0}).IsMap(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(int8(1)).IsMap(), false)
		t.Assert(f.NewVar(uint8(1)).IsMap(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(fvar.New(fvar.New("asd")).IsMap(), false)
		t.Assert(fvar.New(&f.Map{"k": "v"}).IsMap(), true)
	})
}

func TestVar_IsStruct(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(0).IsStruct(), false)
		t.Assert(f.NewVar(nil).IsStruct(), false)
		t.Assert(f.NewVar(f.Map{}).IsStruct(), false)
		t.Assert(f.NewVar(f.Slice{}).IsStruct(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(1).IsStruct(), false)
		t.Assert(f.NewVar(-1).IsStruct(), false)
		t.Assert(f.NewVar(0.1).IsStruct(), false)
		t.Assert(f.NewVar(float64(1)).IsStruct(), false)
		t.Assert(f.NewVar(f.Map{"k": "v"}).IsStruct(), false)
		t.Assert(f.NewVar(f.Slice{0}).IsStruct(), false)
	})
	ftest.C(t, func(t *ftest.T) {
		a := &struct {
		}{}
		t.Assert(f.NewVar(a).IsStruct(), true)
		t.Assert(f.NewVar(*a).IsStruct(), true)
		t.Assert(f.NewVar(&a).IsStruct(), true)
	})
}
