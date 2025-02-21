package fbinary_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fbinary"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_LeEncodeAndLeDecode(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		for k, v := range testData {
			ve := fbinary.LeEncode(v)
			ve1 := fbinary.LeEncodeByLength(len(ve), v)

			// t.Logf("%s:%v, encoded:%v\n", k, v, ve)
			switch v.(type) {
			case int:
				t.Assert(fbinary.LeDecodeToInt(ve), v)
				t.Assert(fbinary.LeDecodeToInt(ve1), v)
			case int8:
				t.Assert(fbinary.LeDecodeToInt8(ve), v)
				t.Assert(fbinary.LeDecodeToInt8(ve1), v)
			case int16:
				t.Assert(fbinary.LeDecodeToInt16(ve), v)
				t.Assert(fbinary.LeDecodeToInt16(ve1), v)
			case int32:
				t.Assert(fbinary.LeDecodeToInt32(ve), v)
				t.Assert(fbinary.LeDecodeToInt32(ve1), v)
			case int64:
				t.Assert(fbinary.LeDecodeToInt64(ve), v)
				t.Assert(fbinary.LeDecodeToInt64(ve1), v)
			case uint:
				t.Assert(fbinary.LeDecodeToUint(ve), v)
				t.Assert(fbinary.LeDecodeToUint(ve1), v)
			case uint8:
				t.Assert(fbinary.LeDecodeToUint8(ve), v)
				t.Assert(fbinary.LeDecodeToUint8(ve1), v)
			case uint16:
				t.Assert(fbinary.LeDecodeToUint16(ve1), v)
				t.Assert(fbinary.LeDecodeToUint16(ve), v)
			case uint32:
				t.Assert(fbinary.LeDecodeToUint32(ve1), v)
				t.Assert(fbinary.LeDecodeToUint32(ve), v)
			case uint64:
				t.Assert(fbinary.LeDecodeToUint64(ve), v)
				t.Assert(fbinary.LeDecodeToUint64(ve1), v)
			case bool:
				t.Assert(fbinary.LeDecodeToBool(ve), v)
				t.Assert(fbinary.LeDecodeToBool(ve1), v)
			case string:
				t.Assert(fbinary.LeDecodeToString(ve), v)
				t.Assert(fbinary.LeDecodeToString(ve1), v)
			case float32:
				t.Assert(fbinary.LeDecodeToFloat32(ve), v)
				t.Assert(fbinary.LeDecodeToFloat32(ve1), v)
			case float64:
				t.Assert(fbinary.LeDecodeToFloat64(ve), v)
				t.Assert(fbinary.LeDecodeToFloat64(ve1), v)
			default:
				if v == nil {
					continue
				}
				res := make([]byte, len(ve))
				err := fbinary.LeDecode(ve, res)
				if err != nil {
					t.Errorf("test data: %s, %v, error:%v", k, v, err)
				}
				t.Assert(res, v)
			}
		}
	})
}

func Test_LeEncodeStruct(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		user := User{"wenzi1", 999, "www.baidu.com"}
		ve := fbinary.LeEncode(user)
		s := fbinary.LeDecodeToString(ve)
		t.Assert(s, s)
	})
}
