package fbinary_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fbinary"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_BeEncodeAndBeDecode(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		for k, v := range testData {
			ve := fbinary.BeEncode(v)
			ve1 := fbinary.BeEncodeByLength(len(ve), v)

			// t.Logf("%s:%v, encoded:%v\n", k, v, ve)
			switch v.(type) {
			case int:
				t.Assert(fbinary.BeDecodeToInt(ve), v)
				t.Assert(fbinary.BeDecodeToInt(ve1), v)
			case int8:
				t.Assert(fbinary.BeDecodeToInt8(ve), v)
				t.Assert(fbinary.BeDecodeToInt8(ve1), v)
			case int16:
				t.Assert(fbinary.BeDecodeToInt16(ve), v)
				t.Assert(fbinary.BeDecodeToInt16(ve1), v)
			case int32:
				t.Assert(fbinary.BeDecodeToInt32(ve), v)
				t.Assert(fbinary.BeDecodeToInt32(ve1), v)
			case int64:
				t.Assert(fbinary.BeDecodeToInt64(ve), v)
				t.Assert(fbinary.BeDecodeToInt64(ve1), v)
			case uint:
				t.Assert(fbinary.BeDecodeToUint(ve), v)
				t.Assert(fbinary.BeDecodeToUint(ve1), v)
			case uint8:
				t.Assert(fbinary.BeDecodeToUint8(ve), v)
				t.Assert(fbinary.BeDecodeToUint8(ve1), v)
			case uint16:
				t.Assert(fbinary.BeDecodeToUint16(ve1), v)
				t.Assert(fbinary.BeDecodeToUint16(ve), v)
			case uint32:
				t.Assert(fbinary.BeDecodeToUint32(ve1), v)
				t.Assert(fbinary.BeDecodeToUint32(ve), v)
			case uint64:
				t.Assert(fbinary.BeDecodeToUint64(ve), v)
				t.Assert(fbinary.BeDecodeToUint64(ve1), v)
			case bool:
				t.Assert(fbinary.BeDecodeToBool(ve), v)
				t.Assert(fbinary.BeDecodeToBool(ve1), v)
			case string:
				t.Assert(fbinary.BeDecodeToString(ve), v)
				t.Assert(fbinary.BeDecodeToString(ve1), v)
			case float32:
				t.Assert(fbinary.BeDecodeToFloat32(ve), v)
				t.Assert(fbinary.BeDecodeToFloat32(ve1), v)
			case float64:
				t.Assert(fbinary.BeDecodeToFloat64(ve), v)
				t.Assert(fbinary.BeDecodeToFloat64(ve1), v)
			default:
				if v == nil {
					continue
				}
				res := make([]byte, len(ve))
				err := fbinary.BeDecode(ve, res)
				if err != nil {
					t.Errorf("test data: %s, %v, error:%v", k, v, err)
				}
				t.Assert(res, v)
			}
		}
	})
}

func Test_BeEncodeStruct(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		user := User{"wenzi1", 999, "www.baidu.com"}
		ve := fbinary.BeEncode(user)
		s := fbinary.BeDecodeToString(ve)
		t.Assert(string(s), s)
	})
}
