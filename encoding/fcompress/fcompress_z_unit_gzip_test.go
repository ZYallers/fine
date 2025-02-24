package fcompress_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fcompress"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/os/ftime"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_Gzip_UnGzip(t *testing.T) {
	var (
		src  = "Hello World!!"
		gzip = []byte{
			0x1f, 0x8b, 0x08, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0xff,
			0xf2, 0x48, 0xcd, 0xc9, 0xc9,
			0x57, 0x08, 0xcf, 0x2f, 0xca,
			0x49, 0x51, 0x54, 0x04, 0x04,
			0x00, 0x00, 0xff, 0xff, 0x9d,
			0x24, 0xa8, 0xd1, 0x0d, 0x00,
			0x00, 0x00,
		}
	)

	ftest.C(t, func(t *ftest.T) {
		arr := []byte(src)
		data, _ := fcompress.Gzip(arr)
		t.Assert(data, gzip)

		data, _ = fcompress.UnGzip(gzip)
		t.Assert(data, arr)

		data, _ = fcompress.UnGzip(gzip[1:])
		t.Assert(data, nil)
	})
}

func Test_Gzip_UnGzip_File(t *testing.T) {
	var (
		srcPath  = ftest.DataPath("gzip", "file.txt")
		dstPath1 = ffile.Temp(ftime.TimestampNanoStr(), "gzip.zip")
		dstPath2 = ffile.Temp(ftime.TimestampNanoStr(), "file.txt")
	)

	// Compress.
	ftest.C(t, func(t *ftest.T) {
		err := fcompress.GzipFile(srcPath, dstPath1, 9)
		t.AssertNil(err)
		defer ffile.Remove(dstPath1)
		t.Assert(ffile.Exists(dstPath1), true)

		// Decompress.
		err = fcompress.UnGzipFile(dstPath1, dstPath2)
		t.AssertNil(err)
		defer ffile.Remove(dstPath2)
		t.Assert(ffile.Exists(dstPath2), true)

		t.Assert(ffile.GetContents(srcPath), ffile.GetContents(dstPath2))
	})
}
