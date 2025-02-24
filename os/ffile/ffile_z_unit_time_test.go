package ffile_test

import (
	"os"
	"testing"
	"time"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_MTime(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {

		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.AssertNil(err)

		t.Assert(ffile.MTime(testpath()+file1), fileobj.ModTime())
		t.Assert(ffile.MTime(""), "")
	})
}

func Test_MTimeMillisecond(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.AssertNil(err)

		time.Sleep(time.Millisecond * 100)
		t.AssertGE(
			ffile.MTimestampMilli(testpath()+file1),
			fileobj.ModTime().UnixNano()/1000000,
		)
		t.Assert(ffile.MTimestampMilli(""), -1)
	})
}
