package fvar_test

import (
	"bytes"
	"encoding/binary"
	"testing"
	"time"

	"github.com/ZYallers/fine/container/fvar"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_Set(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var v fvar.Var
		v.Set(123.456)
		t.Assert(v.Val(), 123.456)
	})
	ftest.C(t, func(t *ftest.T) {
		var v fvar.Var
		v.Set(123.456)
		t.Assert(v.Val(), 123.456)
	})

	ftest.C(t, func(t *ftest.T) {
		objOne := fvar.New("old", true)
		objOneOld, _ := objOne.Set("new").(string)
		t.Assert(objOneOld, "old")

		objTwo := fvar.New("old", false)
		objTwoOld, _ := objTwo.Set("new").(string)
		t.Assert(objTwoOld, "old")
	})
}

func Test_Val(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		objOne := fvar.New(1, true)
		objOneOld, _ := objOne.Val().(int)
		t.Assert(objOneOld, 1)

		objTwo := fvar.New(1, false)
		objTwoOld, _ := objTwo.Val().(int)
		t.Assert(objTwoOld, 1)

		objOne = nil
		t.Assert(objOne.Val(), nil)
	})
}

func Test_Interface(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		objOne := fvar.New(1, true)
		objOneOld, _ := objOne.Interface().(int)
		t.Assert(objOneOld, 1)

		objTwo := fvar.New(1, false)
		objTwoOld, _ := objTwo.Interface().(int)
		t.Assert(objTwoOld, 1)
	})
}

func Test_IsNil(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		objOne := fvar.New(nil, true)
		t.Assert(objOne.IsNil(), true)

		objTwo := fvar.New("noNil", false)
		t.Assert(objTwo.IsNil(), false)

	})
}

func Test_Bytes(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		x := int32(1)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, x)

		objOne := fvar.New(bytesBuffer.Bytes(), true)

		bBuf := bytes.NewBuffer(objOne.Bytes())
		var y int32
		binary.Read(bBuf, binary.BigEndian, &y)

		t.Assert(x, y)

	})
}

func Test_String(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var str string = "hello"
		objOne := fvar.New(str, true)
		t.Assert(objOne.String(), str)

	})
}

func Test_Bool(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var ok bool = true
		objOne := fvar.New(ok, true)
		t.Assert(objOne.Bool(), ok)

		ok = false
		objTwo := fvar.New(ok, true)
		t.Assert(objTwo.Bool(), ok)

	})
}

func Test_Int(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num int = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Int(), num)

	})
}

func Test_Int8(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num int8 = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Int8(), num)

	})
}

func Test_Int16(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num int16 = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Int16(), num)

	})
}

func Test_Int32(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num int32 = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Int32(), num)

	})
}

func Test_Int64(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num int64 = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Int64(), num)

	})
}

func Test_Uint(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num uint = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Uint(), num)

	})
}

func Test_Uint8(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num uint8 = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Uint8(), num)

	})
}

func Test_Uint16(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num uint16 = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Uint16(), num)

	})
}

func Test_Uint32(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num uint32 = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Uint32(), num)

	})
}

func Test_Uint64(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num uint64 = 1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Uint64(), num)

	})
}

func Test_Float32(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num float32 = 1.1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Float32(), num)

	})
}

func Test_Float64(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var num float64 = 1.1
		objOne := fvar.New(num, true)
		t.Assert(objOne.Float64(), num)

	})
}

func Test_Time(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var timeUnix int64 = 1556242660
		objOne := fvar.New(timeUnix, true)
		t.Assert(objOne.Time().Unix(), timeUnix)
	})
}

func Test_GTime(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var timeUnix int64 = 1556242660
		objOne := fvar.New(timeUnix, true)
		t.Assert(objOne.FTime().Unix(), timeUnix)
	})
}

func Test_Duration(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var timeUnix int64 = 1556242660
		objOne := fvar.New(timeUnix, true)
		t.Assert(objOne.Duration(), time.Duration(timeUnix))
	})
}

func Test_UnmarshalJson(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type V struct {
			Name string
			Var  *fvar.Var
		}
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "v",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.String(), "v")
	})
	ftest.C(t, func(t *ftest.T) {
		type V struct {
			Name string
			Var  fvar.Var
		}
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "v",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.String(), "v")
	})
}

func Test_UnmarshalValue(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type V struct {
			Name string
			Var  *fvar.Var
		}
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "v",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.String(), "v")
	})
	ftest.C(t, func(t *ftest.T) {
		type V struct {
			Name string
			Var  fvar.Var
		}
		var v *V
		err := fconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "v",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.String(), "v")
	})
}

func Test_Copy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		src := f.Map{
			"k1": "v1",
			"k2": "v2",
		}
		srcVar := fvar.New(src)
		dstVar := srcVar.Copy()
		t.Assert(srcVar.Map(), src)
		t.Assert(dstVar.Map(), src)

		dstVar.Map()["k3"] = "v3"
		t.Assert(srcVar.Map(), f.Map{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(dstVar.Map(), f.Map{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
	})
}

func Test_DeepCopy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		src := f.Map{
			"k1": "v1",
			"k2": "v2",
		}
		srcVar := fvar.New(src)
		copyVar := srcVar.DeepCopy().(*fvar.Var)
		copyVar.Set(f.Map{
			"k3": "v3",
			"k4": "v4",
		})
		t.AssertNE(srcVar, copyVar)

		srcVar = nil
		t.AssertNil(srcVar.DeepCopy())
	})
}
