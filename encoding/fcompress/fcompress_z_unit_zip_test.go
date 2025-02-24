package fcompress_test

import (
	"bytes"
	"testing"

	"github.com/ZYallers/fine/encoding/fcompress"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/os/ftime"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_ZipPath(t *testing.T) {
	// file
	ftest.C(t, func(t *ftest.T) {
		srcPath := ftest.DataPath("zip", "path1", "1.txt")
		dstPath := ftest.DataPath("zip", "zip.zip")

		t.Assert(ffile.Exists(dstPath), false)
		t.Assert(fcompress.ZipPath(srcPath, dstPath), nil)
		t.Assert(ffile.Exists(dstPath), true)
		defer ffile.Remove(dstPath)

		// unzip to temporary dir.
		tempDirPath := ffile.Temp(ftime.TimestampNanoStr())
		t.Assert(ffile.Mkdir(tempDirPath), nil)
		t.Assert(fcompress.UnZipFile(dstPath, tempDirPath), nil)
		defer ffile.Remove(tempDirPath)

		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "1.txt")),
			ffile.GetContents(srcPath),
		)
	})
	// multiple files
	ftest.C(t, func(t *ftest.T) {
		var (
			srcPath1 = ftest.DataPath("zip", "path1", "1.txt")
			srcPath2 = ftest.DataPath("zip", "path2", "2.txt")
			dstPath  = ffile.Temp(ftime.TimestampNanoStr(), "zip.zip")
		)
		if p := ffile.Dir(dstPath); !ffile.Exists(p) {
			t.Assert(ffile.Mkdir(p), nil)
		}

		t.Assert(ffile.Exists(dstPath), false)
		err := fcompress.ZipPath(srcPath1+","+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(ffile.Exists(dstPath), true)
		defer ffile.Remove(dstPath)

		// unzip to another temporary dir.
		tempDirPath := ffile.Temp(ftime.TimestampNanoStr())
		t.Assert(ffile.Mkdir(tempDirPath), nil)
		err = fcompress.UnZipFile(dstPath, tempDirPath)
		t.AssertNil(err)
		defer ffile.Remove(tempDirPath)

		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "1.txt")),
			ffile.GetContents(srcPath1),
		)
		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "2.txt")),
			ffile.GetContents(srcPath2),
		)
	})
	// one dir and one file.
	ftest.C(t, func(t *ftest.T) {
		var (
			srcPath1 = ftest.DataPath("zip", "path1")
			srcPath2 = ftest.DataPath("zip", "path2", "2.txt")
			dstPath  = ffile.Temp(ftime.TimestampNanoStr(), "zip.zip")
		)
		if p := ffile.Dir(dstPath); !ffile.Exists(p) {
			t.Assert(ffile.Mkdir(p), nil)
		}

		t.Assert(ffile.Exists(dstPath), false)
		err := fcompress.ZipPath(srcPath1+","+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(ffile.Exists(dstPath), true)
		defer ffile.Remove(dstPath)

		// unzip to another temporary dir.
		tempDirPath := ffile.Temp(ftime.TimestampNanoStr())
		t.Assert(ffile.Mkdir(tempDirPath), nil)
		err = fcompress.UnZipFile(dstPath, tempDirPath)
		t.AssertNil(err)
		defer ffile.Remove(tempDirPath)

		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "path1", "1.txt")),
			ffile.GetContents(ffile.Join(srcPath1, "1.txt")),
		)
		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "2.txt")),
			ffile.GetContents(srcPath2),
		)
	})
	// directory.
	ftest.C(t, func(t *ftest.T) {
		srcPath := ftest.DataPath("zip")
		dstPath := ftest.DataPath("zip", "zip.zip")

		pwd := ffile.Pwd()
		err := ffile.Chdir(srcPath)
		defer ffile.Chdir(pwd)
		t.AssertNil(err)

		t.Assert(ffile.Exists(dstPath), false)
		err = fcompress.ZipPath(srcPath, dstPath)
		t.AssertNil(err)
		t.Assert(ffile.Exists(dstPath), true)
		defer ffile.Remove(dstPath)

		tempDirPath := ffile.Temp(ftime.TimestampNanoStr())
		err = ffile.Mkdir(tempDirPath)
		t.AssertNil(err)

		err = fcompress.UnZipFile(dstPath, tempDirPath)
		t.AssertNil(err)
		defer ffile.Remove(tempDirPath)

		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "zip", "path1", "1.txt")),
			ffile.GetContents(ffile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "zip", "path2", "2.txt")),
			ffile.GetContents(ffile.Join(srcPath, "path2", "2.txt")),
		)
	})
	// multiple directory paths joined using char ','.
	ftest.C(t, func(t *ftest.T) {
		var (
			srcPath  = ftest.DataPath("zip")
			srcPath1 = ftest.DataPath("zip", "path1")
			srcPath2 = ftest.DataPath("zip", "path2")
			dstPath  = ftest.DataPath("zip", "zip.zip")
		)
		pwd := ffile.Pwd()
		err := ffile.Chdir(srcPath)
		defer ffile.Chdir(pwd)
		t.AssertNil(err)

		t.Assert(ffile.Exists(dstPath), false)
		err = fcompress.ZipPath(srcPath1+", "+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(ffile.Exists(dstPath), true)
		defer ffile.Remove(dstPath)

		tempDirPath := ffile.Temp(ftime.TimestampNanoStr())
		err = ffile.Mkdir(tempDirPath)
		t.AssertNil(err)

		zipContent := ffile.GetBytes(dstPath)
		t.AssertGT(len(zipContent), 0)
		err = fcompress.UnZipContent(zipContent, tempDirPath)
		t.AssertNil(err)
		defer ffile.Remove(tempDirPath)

		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "path1", "1.txt")),
			ffile.GetContents(ffile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "path2", "2.txt")),
			ffile.GetContents(ffile.Join(srcPath, "path2", "2.txt")),
		)
	})
}

func Test_ZipPathWriter(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			srcPath  = ftest.DataPath("zip")
			srcPath1 = ftest.DataPath("zip", "path1")
			srcPath2 = ftest.DataPath("zip", "path2")
		)
		pwd := ffile.Pwd()
		err := ffile.Chdir(srcPath)
		defer ffile.Chdir(pwd)
		t.AssertNil(err)

		writer := bytes.NewBuffer(nil)
		t.Assert(writer.Len(), 0)
		err = fcompress.ZipPathWriter(srcPath1+", "+srcPath2, writer)
		t.AssertNil(err)
		t.AssertGT(writer.Len(), 0)

		tempDirPath := ffile.Temp(ftime.TimestampNanoStr())
		err = ffile.Mkdir(tempDirPath)
		t.AssertNil(err)

		zipContent := writer.Bytes()
		t.AssertGT(len(zipContent), 0)
		err = fcompress.UnZipContent(zipContent, tempDirPath)
		t.AssertNil(err)
		defer ffile.Remove(tempDirPath)

		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "path1", "1.txt")),
			ffile.GetContents(ffile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "path2", "2.txt")),
			ffile.GetContents(ffile.Join(srcPath, "path2", "2.txt")),
		)
	})
}

func Test_ZipPathContent(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			srcPath  = ftest.DataPath("zip")
			srcPath1 = ftest.DataPath("zip", "path1")
			srcPath2 = ftest.DataPath("zip", "path2")
		)
		pwd := ffile.Pwd()
		err := ffile.Chdir(srcPath)
		defer ffile.Chdir(pwd)
		t.AssertNil(err)

		tempDirPath := ffile.Temp(ftime.TimestampNanoStr())
		err = ffile.Mkdir(tempDirPath)
		t.AssertNil(err)

		zipContent, err := fcompress.ZipPathContent(srcPath1 + ", " + srcPath2)
		t.AssertGT(len(zipContent), 0)
		err = fcompress.UnZipContent(zipContent, tempDirPath)
		t.AssertNil(err)
		defer ffile.Remove(tempDirPath)

		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "path1", "1.txt")),
			ffile.GetContents(ffile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			ffile.GetContents(ffile.Join(tempDirPath, "path2", "2.txt")),
			ffile.GetContents(ffile.Join(srcPath, "path2", "2.txt")),
		)
	})
}
