package ffile_test

import (
	"testing"

	"github.com/ZYallers/fine/container/farray"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_ScanDir(t *testing.T) {
	teatPath := ftest.DataPath()
	ftest.C(t, func(t *ftest.T) {
		files, err := ffile.ScanDir(teatPath, "*", false)
		t.AssertNil(err)
		t.AssertIN(teatPath+ffile.Separator+"dir1", files)
		t.AssertIN(teatPath+ffile.Separator+"dir2", files)
		t.AssertNE(teatPath+ffile.Separator+"dir1"+ffile.Separator+"file1", files)
	})
	ftest.C(t, func(t *ftest.T) {
		files, err := ffile.ScanDir(teatPath, "*", true)
		t.AssertNil(err)
		t.AssertIN(teatPath+ffile.Separator+"dir1", files)
		t.AssertIN(teatPath+ffile.Separator+"dir2", files)
		t.AssertIN(teatPath+ffile.Separator+"dir1"+ffile.Separator+"file1", files)
		t.AssertIN(teatPath+ffile.Separator+"dir2"+ffile.Separator+"file2", files)
	})
}

func Test_ScanDirFunc(t *testing.T) {
	teatPath := ftest.DataPath()
	ftest.C(t, func(t *ftest.T) {
		files, err := ffile.ScanDirFunc(teatPath, "*", true, func(path string) string {
			if ffile.Name(path) != "file1" {
				return ""
			}
			return path
		})
		t.AssertNil(err)
		t.Assert(len(files), 1)
		t.Assert(ffile.Name(files[0]), "file1")
	})
}

func Test_ScanDirFile(t *testing.T) {
	teatPath := ftest.DataPath()
	ftest.C(t, func(t *ftest.T) {
		files, err := ffile.ScanDirFile(teatPath, "*", false)
		t.AssertNil(err)
		t.Assert(len(files), 0)
	})
	ftest.C(t, func(t *ftest.T) {
		files, err := ffile.ScanDirFile(teatPath, "*", true)
		t.AssertNil(err)
		t.AssertNI(teatPath+ffile.Separator+"dir1", files)
		t.AssertNI(teatPath+ffile.Separator+"dir2", files)
		t.AssertIN(teatPath+ffile.Separator+"dir1"+ffile.Separator+"file1", files)
		t.AssertIN(teatPath+ffile.Separator+"dir2"+ffile.Separator+"file2", files)
	})
}

func Test_ScanDirFileFunc(t *testing.T) {
	teatPath := ftest.DataPath()
	ftest.C(t, func(t *ftest.T) {
		array := farray.New()
		files, err := ffile.ScanDirFileFunc(teatPath, "*", false, func(path string) string {
			array.Append(1)
			return path
		})
		t.AssertNil(err)
		t.Assert(len(files), 0)
		t.Assert(array.Len(), 0)
	})
	ftest.C(t, func(t *ftest.T) {
		array := farray.New()
		files, err := ffile.ScanDirFileFunc(teatPath, "*", true, func(path string) string {
			array.Append(1)
			if ffile.Basename(path) == "file1" {
				return path
			}
			return ""
		})
		t.AssertNil(err)
		t.Assert(len(files), 1)
		t.Assert(array.Len(), 3)
	})
}
