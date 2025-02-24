package ffile_test

import (
	"os"
	"testing"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/os/ftime"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fuid"
)

func Test_Copy(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths  = "/testfile_copyfile1.txt"
			topath = "/testfile_copyfile2.txt"
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)

		t.Assert(ffile.Copy(testpath()+paths, testpath()+topath), nil)
		defer delTestFiles(topath)

		t.Assert(ffile.IsFile(testpath()+topath), true)
		t.AssertNE(ffile.Copy(paths, ""), nil)
		t.AssertNE(ffile.Copy("", topath), nil)
	})
}

func Test_Copy_File_To_Dir(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			src = ftest.DataPath("dir1", "file1")
			dst = ffile.Temp(fuid.S(), "dir2")
		)
		err := ffile.Mkdir(dst)
		t.AssertNil(err)
		defer ffile.Remove(dst)

		err = ffile.Copy(src, dst)
		t.AssertNil(err)

		expectPath := ffile.Join(dst, "file1")
		t.Assert(ffile.GetContents(expectPath), ffile.GetContents(src))
	})
}

func Test_Copy_Dir_To_File(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			src = ftest.DataPath("dir1")
			dst = ffile.Temp(fuid.S(), "file2")
		)
		f, err := ffile.Create(dst)
		t.AssertNil(err)
		defer f.Close()
		defer ffile.Remove(dst)

		err = ffile.Copy(src, dst)
		t.AssertNE(err, nil)
	})
}

func Test_CopyFile(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths  = "/testfile_copyfile1.txt"
			topath = "/testfile_copyfile2.txt"
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)

		t.Assert(ffile.CopyFile(testpath()+paths, testpath()+topath), nil)
		defer delTestFiles(topath)

		t.Assert(ffile.IsFile(testpath()+topath), true)
		t.AssertNE(ffile.CopyFile(paths, ""), nil)
		t.AssertNE(ffile.CopyFile("", topath), nil)
	})
	// Content replacement.
	ftest.C(t, func(t *ftest.T) {
		src := ffile.Temp(ftime.TimestampNanoStr())
		dst := ffile.Temp(ftime.TimestampNanoStr())
		srcContent := "1"
		dstContent := "1"
		t.Assert(ffile.PutContents(src, srcContent), nil)
		t.Assert(ffile.PutContents(dst, dstContent), nil)
		t.Assert(ffile.GetContents(src), srcContent)
		t.Assert(ffile.GetContents(dst), dstContent)

		t.Assert(ffile.CopyFile(src, dst), nil)
		t.Assert(ffile.GetContents(src), srcContent)
		t.Assert(ffile.GetContents(dst), srcContent)
	})
	// Set mode
	ftest.C(t, func(t *ftest.T) {
		var (
			src     = "/testfile_copyfile1.txt"
			dst     = "/testfile_copyfile2.txt"
			dstMode = os.FileMode(0600)
		)
		t.AssertNil(createTestFile(src, ""))
		defer delTestFiles(src)

		t.Assert(ffile.CopyFile(testpath()+src, testpath()+dst, ffile.CopyOption{Mode: dstMode}), nil)
		defer delTestFiles(dst)

		dstStat, err := ffile.Stat(testpath() + dst)
		t.AssertNil(err)
		t.Assert(dstStat.Mode().Perm(), dstMode)
	})
	// Preserve src file's mode
	ftest.C(t, func(t *ftest.T) {
		var (
			src = "/testfile_copyfile1.txt"
			dst = "/testfile_copyfile2.txt"
		)
		t.AssertNil(createTestFile(src, ""))
		defer delTestFiles(src)

		t.Assert(ffile.CopyFile(testpath()+src, testpath()+dst, ffile.CopyOption{PreserveMode: true}), nil)
		defer delTestFiles(dst)

		srcStat, err := ffile.Stat(testpath() + src)
		t.AssertNil(err)
		dstStat, err := ffile.Stat(testpath() + dst)
		t.AssertNil(err)
		t.Assert(srcStat.Mode().Perm(), dstStat.Mode().Perm())
	})
}

func Test_CopyDir(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			dirPath1 = "/test-copy-dir1"
			dirPath2 = "/test-copy-dir2"
		)
		haveList := []string{
			"t1.txt",
			"t2.txt",
		}
		createDir(dirPath1)
		for _, v := range haveList {
			t.Assert(createTestFile(dirPath1+"/"+v, ""), nil)
		}
		defer delTestFiles(dirPath1)

		var (
			yfolder  = testpath() + dirPath1
			tofolder = testpath() + dirPath2
		)

		if ffile.IsDir(tofolder) {
			t.Assert(ffile.Remove(tofolder), nil)
			t.Assert(ffile.Remove(""), nil)
		}

		t.Assert(ffile.CopyDir(yfolder, tofolder), nil)
		defer delTestFiles(tofolder)

		t.Assert(ffile.IsDir(yfolder), true)

		for _, v := range haveList {
			t.Assert(ffile.IsFile(yfolder+"/"+v), true)
		}

		t.Assert(ffile.IsDir(tofolder), true)

		for _, v := range haveList {
			t.Assert(ffile.IsFile(tofolder+"/"+v), true)
		}

		t.Assert(ffile.Remove(tofolder), nil)
		t.Assert(ffile.Remove(""), nil)
	})
	// Content replacement.
	ftest.C(t, func(t *ftest.T) {
		src := ffile.Temp(ftime.TimestampNanoStr(), ftime.TimestampNanoStr())
		dst := ffile.Temp(ftime.TimestampNanoStr(), ftime.TimestampNanoStr())
		defer func() {
			ffile.Remove(src)
			ffile.Remove(dst)
		}()
		srcContent := "1"
		dstContent := "1"
		t.Assert(ffile.PutContents(src, srcContent), nil)
		t.Assert(ffile.PutContents(dst, dstContent), nil)
		t.Assert(ffile.GetContents(src), srcContent)
		t.Assert(ffile.GetContents(dst), dstContent)

		err := ffile.CopyDir(ffile.Dir(src), ffile.Dir(dst))
		t.AssertNil(err)
		t.Assert(ffile.GetContents(src), srcContent)
		t.Assert(ffile.GetContents(dst), srcContent)

		t.AssertNE(ffile.CopyDir(ffile.Dir(src), ""), nil)
		t.AssertNE(ffile.CopyDir("", ffile.Dir(dst)), nil)
	})
}
