package fxml_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fxml"
	"github.com/ZYallers/fine/test/ftest"
)

const xmlData = `<?xml version="1.0" encoding="UTF-8"?><users><age>12</age><name>李四</name></users>`

func Test_XmlToData(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		srcJson, err := fxml.ToJson([]byte(xmlData))
		t.AssertNil(err)
		t.Assert(string(srcJson), `{"users":{"age":"12","name":"李四"}}`)
	})

}

func Test_Decode(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		srcMap, err := fxml.Decode([]byte(xmlData))
		t.AssertNil(err)
		t.Assert(srcMap, map[string]interface{}{"users": map[string]interface{}{"age": "12", "name": "李四"}})
	})
}

func Test_DecodeWithoutRoot(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		srcMap, err := fxml.DecodeWithoutRoot([]byte(xmlData))
		t.AssertNil(err)
		t.Assert(srcMap, map[string]interface{}{"age": "12", "name": "李四"})
	})
}

func Test_Encode(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		srcMap := map[string]interface{}{"users": map[string]interface{}{"age": "12", "name": "李四"}}
		srcXml, err := fxml.Encode(srcMap)
		t.AssertNil(err)
		xmlHead := `<?xml version="1.0" encoding="UTF-8"?>`
		t.Assert(xmlHead+string(srcXml), xmlData)
	})
}

func Test_EncodeWithIndent(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		srcMap := map[string]interface{}{"age": "12", "name": "李四"}
		srcXml, err := fxml.EncodeWithIndent(srcMap, "users")
		t.AssertNil(err)
		t.Assert(string(srcXml), `<users>
	<age>12</age>
	<name>李四</name>
</users>`)
	})
}
