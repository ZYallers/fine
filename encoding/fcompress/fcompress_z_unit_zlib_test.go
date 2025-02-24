package fcompress_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fcompress"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_Zlib_UnZlib(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		src := "hello, world\n"
		dst := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
		data, _ := fcompress.Zlib([]byte(src))
		t.Assert(data, dst)

		data, _ = fcompress.UnZlib(dst)
		t.Assert(data, []byte(src))

		data, _ = fcompress.Zlib(nil)
		t.Assert(data, nil)
		data, _ = fcompress.UnZlib(nil)
		t.Assert(data, nil)

		data, _ = fcompress.UnZlib(dst[1:])
		t.Assert(data, nil)
	})
}
