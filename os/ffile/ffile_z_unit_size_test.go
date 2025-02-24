package ffile_test

import (
	"testing"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
)

func Test_Size(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1 string = "/testfile_t1.txt"
			sizes  int64
		)

		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)

		sizes = ffile.Size(testpath() + paths1)
		t.Assert(sizes, 14)

		sizes = ffile.Size("")
		t.Assert(sizes, 0)

	})
}

func Test_SizeFormat(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1 = "/testfile_t1.txt"
			sizes  string
		)

		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)

		sizes = ffile.SizeFormat(testpath() + paths1)
		t.Assert(sizes, "14.00B")

		sizes = ffile.SizeFormat("")
		t.Assert(sizes, "0.00B")

	})
}

func Test_StrToSize(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(ffile.StrToSize("0.00B"), 0)
		t.Assert(ffile.StrToSize("16.00B"), 16)
		t.Assert(ffile.StrToSize("1.00K"), 1024)
		t.Assert(ffile.StrToSize("1.00KB"), 1024)
		t.Assert(ffile.StrToSize("1.00KiloByte"), 1024)
		t.Assert(ffile.StrToSize("15.26M"), fconv.Int64(15.26*1024*1024))
		t.Assert(ffile.StrToSize("15.26MB"), fconv.Int64(15.26*1024*1024))
		t.Assert(ffile.StrToSize("1.49G"), fconv.Int64(1.49*1024*1024*1024))
		t.Assert(ffile.StrToSize("1.49GB"), fconv.Int64(1.49*1024*1024*1024))
		t.Assert(ffile.StrToSize("8.73T"), fconv.Int64(8.73*1024*1024*1024*1024))
		t.Assert(ffile.StrToSize("8.73TB"), fconv.Int64(8.73*1024*1024*1024*1024))
		t.Assert(ffile.StrToSize("8.53P"), fconv.Int64(8.53*1024*1024*1024*1024*1024))
		t.Assert(ffile.StrToSize("8.53PB"), fconv.Int64(8.53*1024*1024*1024*1024*1024))
		t.Assert(ffile.StrToSize("8.01EB"), fconv.Int64(8.01*1024*1024*1024*1024*1024*1024))
		t.Assert(ffile.StrToSize("0.01ZB"), fconv.Int64(0.01*1024*1024*1024*1024*1024*1024*1024))
		t.Assert(ffile.StrToSize("0.01YB"), fconv.Int64(0.01*1024*1024*1024*1024*1024*1024*1024*1024))
		t.Assert(ffile.StrToSize("0.01BB"), fconv.Int64(0.01*1024*1024*1024*1024*1024*1024*1024*1024*1024))
		t.Assert(ffile.StrToSize("0.01AB"), fconv.Int64(-1))
		t.Assert(ffile.StrToSize("123456789"), 123456789)
	})
}

func Test_FormatSize(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(ffile.FormatSize(0), "0.00B")
		t.Assert(ffile.FormatSize(16), "16.00B")

		t.Assert(ffile.FormatSize(1024), "1.00K")

		t.Assert(ffile.FormatSize(16000000), "15.26M")

		t.Assert(ffile.FormatSize(1600000000), "1.49G")

		t.Assert(ffile.FormatSize(9600000000000), "8.73T")
		t.Assert(ffile.FormatSize(9600000000000000), "8.53P")
	})
}

func Test_ReadableSize(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {

		var (
			paths1 string = "/testfile_t1.txt"
		)
		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)
		t.Assert(ffile.ReadableSize(testpath()+paths1), "14.00B")
		t.Assert(ffile.ReadableSize(""), "0.00B")

	})
}
