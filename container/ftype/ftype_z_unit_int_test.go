package ftype_test

import (
	"sync"
	"testing"

	"github.com/ZYallers/fine/container/ftype"
	"github.com/ZYallers/fine/internal/json"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_Int(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var wg sync.WaitGroup
		addTimes := 1000
		i := ftype.NewInt(0)
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(1), 0)
		t.AssertEQ(iClone.Val(), 1)
		for index := 0; index < addTimes; index++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				i.Add(1)
			}()
		}
		wg.Wait()
		t.AssertEQ(addTimes, i.Val())

		// empty param test
		i1 := ftype.NewInt()
		t.AssertEQ(i1.Val(), 0)

		i2 := ftype.NewInt(11)
		t.AssertEQ(i2.Add(1), 12)
		t.AssertEQ(i2.Cas(11, 13), false)
		t.AssertEQ(i2.Cas(12, 13), true)
		t.AssertEQ(i2.String(), "13")

		copyVal := i2.DeepCopy()
		i2.Set(14)
		t.AssertNE(copyVal, iClone.Val())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Int_JSON(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		v := 666
		i := ftype.NewInt(v)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := ftype.NewInt()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), v)
	})
}

func Test_Int_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *ftype.Int
	}
	ftest.C(t, func(t *ftest.T) {
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "123")
	})
}
