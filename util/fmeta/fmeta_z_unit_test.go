package fmeta_test

import (
	"testing"

	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
	"github.com/ZYallers/fine/util/fmeta"
)

func TestMeta_Basic(t *testing.T) {
	type A struct {
		fmeta.Meta `tag:"123" orm:"456"`
		Id         int
		Name       string
	}

	ftest.C(t, func(t *ftest.T) {
		a := &A{
			Id:   100,
			Name: "john",
		}
		t.Assert(len(fmeta.Data(a)), 2)
		t.AssertEQ(fmeta.Get(a, "tag").String(), "123")
		t.AssertEQ(fmeta.Get(a, "orm").String(), "456")
		t.AssertEQ(fmeta.Get(a, "none"), nil)

		b, err := json.Marshal(a)
		t.AssertNil(err)
		t.Assert(b, `{"Id":100,"Name":"john"}`)
	})
}

func TestMeta_Convert_Map(t *testing.T) {
	type A struct {
		fmeta.Meta `tag:"123" orm:"456"`
		Id         int
		Name       string
	}

	ftest.C(t, func(t *ftest.T) {
		a := &A{
			Id:   100,
			Name: "john",
		}
		m := fconv.Map(a)
		t.Assert(len(m), 2)
		t.Assert(m[`Meta`], nil)
	})
}

func TestMeta_Json(t *testing.T) {
	type A struct {
		fmeta.Meta `tag:"123" orm:"456"`
		Id         int
	}

	ftest.C(t, func(t *ftest.T) {
		a := &A{
			Id: 100,
		}
		b, err := json.Marshal(a)
		t.AssertNil(err)
		t.Assert(string(b), `{"Id":100}`)
	})
}
