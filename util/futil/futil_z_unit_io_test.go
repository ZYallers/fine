package futil_test

import (
	"io/ioutil"
	"testing"

	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_NewReadCloser(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		r := futil.NewReadCloser([]byte("test"), true)
		buf := make([]byte, 4)
		_, _ = r.Read(buf)
		t.Assert(string(buf), "test")
		b, _ := ioutil.ReadAll(r)
		t.Assert(string(b), "test")
	})
}
