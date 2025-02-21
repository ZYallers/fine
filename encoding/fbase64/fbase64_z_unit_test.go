package fbase64_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fbase64"
	"github.com/ZYallers/fine/test/ftest"
)

type testPair struct {
	decoded, encoded string
}

var pairs = []testPair{
	// RFC 3548 examples
	{"\x14\xfb\x9c\x03\xd9\x7e", "FPucA9l+"},
	{"\x14\xfb\x9c\x03\xd9", "FPucA9k="},
	{"\x14\xfb\x9c\x03", "FPucAw=="},

	// RFC 4648 examples
	{"", ""},
	{"f", "Zg=="},
	{"fo", "Zm8="},
	{"foo", "Zm9v"},
	{"foob", "Zm9vYg=="},
	{"fooba", "Zm9vYmE="},
	{"foobar", "Zm9vYmFy"},

	// Wikipedia examples
	{"sure.", "c3VyZS4="},
	{"sure", "c3VyZQ=="},
	{"sur", "c3Vy"},
	{"su", "c3U="},
	{"leasure.", "bGVhc3VyZS4="},
	{"easure.", "ZWFzdXJlLg=="},
	{"asure.", "YXN1cmUu"},
	{"sure.", "c3VyZS4="},
}

func Test_Basic(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		for k := range pairs {
			// Encode
			t.Assert(fbase64.Encode([]byte(pairs[k].decoded)), []byte(pairs[k].encoded))
			t.Assert(fbase64.EncodeToString([]byte(pairs[k].decoded)), pairs[k].encoded)
			t.Assert(fbase64.EncodeString(pairs[k].decoded), pairs[k].encoded)

			// Decode
			r1, _ := fbase64.Decode([]byte(pairs[k].encoded))
			t.Assert(r1, []byte(pairs[k].decoded))

			r2, _ := fbase64.DecodeString(pairs[k].encoded)
			t.Assert(r2, []byte(pairs[k].decoded))

			r3, _ := fbase64.DecodeToString(pairs[k].encoded)
			t.Assert(r3, pairs[k].decoded)
		}
	})
}

func Test_File(t *testing.T) {
	path := ftest.DataPath("test")
	expect := "dGVzdA=="
	ftest.C(t, func(t *ftest.T) {
		b, err := fbase64.EncodeFile(path)
		t.AssertNil(err)
		t.Assert(string(b), expect)
	})
	ftest.C(t, func(t *ftest.T) {
		s, err := fbase64.EncodeFileToString(path)
		t.AssertNil(err)
		t.Assert(s, expect)
	})
}

func Test_File_Error(t *testing.T) {
	path := "none-exist-file"
	expect := ""
	ftest.C(t, func(t *ftest.T) {
		b, err := fbase64.EncodeFile(path)
		t.AssertNE(err, nil)
		t.Assert(string(b), expect)
	})
	ftest.C(t, func(t *ftest.T) {
		s, err := fbase64.EncodeFileToString(path)
		t.AssertNE(err, nil)
		t.Assert(s, expect)
	})
}
