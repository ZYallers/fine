package futil_test

import (
	"bytes"
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/internal/util/utils"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_Copy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.Copy(0), 0)
		t.Assert(futil.Copy(1), 1)
		t.Assert(futil.Copy("a"), "a")
		t.Assert(futil.Copy(nil), nil)
	})
	ftest.C(t, func(t *ftest.T) {
		src := f.Map{
			"k1": "v1",
			"k2": "v2",
		}
		dst := futil.Copy(src)
		t.Assert(dst, src)

		dst.(f.Map)["k3"] = "v3"
		t.Assert(src, f.Map{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(dst, f.Map{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
	})
}

func Test_CopyBytes(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var b bytes.Buffer
		_, _ = b.WriteString("a")
		res, err := futil.CopyBytes(&b)
		t.Assert(err, nil)
		t.Assert(res, []byte("a"))
	})
}

func Test_DrainBody(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		b := utils.NewReadCloser([]byte("a"), true)
		origin, value, err := futil.DrainBody(b)
		r1, _ := futil.CopyBytes(origin)
		r2, _ := futil.CopyBytes(value)
		t.Assert(string(r1), "a")
		t.Assert(string(r2), "a")
		t.Assert(err, nil)
	})
}
