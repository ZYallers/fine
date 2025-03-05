package fyaml_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fyaml"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
)

const yamlStr = `
#url属性值
url: https://fine.com
#server地址
server:
    - 121.78.12.34
    - 121.78.12.35
`

func Test_Encode(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		out, err := fyaml.Encode(f.Map{"a": "b"})
		t.AssertNil(err)
		t.Assert(string(out), `a: b`)
	})
}

func Test_EncodeIndent(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		b, err := fyaml.EncodeIndent([]string{"a", "b", "c"}, "####")
		t.AssertNil(err)
		t.Assert(string(b), `####- a
####- b
####- c
`)
	})
}

func Test_Decode(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		result, err := fyaml.Decode([]byte(yamlStr))
		t.AssertNil(err)
		t.Assert(result, map[string]interface{}{
			"url":    "https://fine.com",
			"server": f.Slice{"121.78.12.34", "121.78.12.35"},
		})
	})
}

func Test_DecodeTo(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		result := make(map[string]interface{})
		err := fyaml.DecodeTo([]byte(yamlStr), &result)
		t.AssertNil(err)
		t.Assert(result, map[string]interface{}{
			"url":    "https://fine.com",
			"server": f.Slice{"121.78.12.34", "121.78.12.35"},
		})
	})
}

func Test_ToJson(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		out, err := fyaml.ToJson([]byte(yamlStr))
		t.AssertNil(err)
		t.Assert(string(out), `{"server":["121.78.12.34","121.78.12.35"],"url":"https://fine.com"}`)
	})
}
