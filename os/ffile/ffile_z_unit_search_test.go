package ffile_test

import (
	"path/filepath"
	"testing"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/test/ftest"
)

func Test_Search(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1  string = "/testfiless"
			paths2  string = "./testfile/dirfiles_no"
			tpath   string
			tpath2  string
			tempstr string
			ypaths1 string
			err     error
		)

		createDir(paths1)
		defer delTestFiles(paths1)
		ypaths1 = paths1

		tpath, err = ffile.Search(testpath() + paths1)
		t.AssertNil(err)

		tpath = filepath.ToSlash(tpath)

		// 自定义优先路径
		tpath2, err = ffile.Search(testpath() + paths1)
		t.AssertNil(err)
		tpath2 = filepath.ToSlash(tpath2)

		tempstr = testpath()
		paths1 = tempstr + paths1
		paths1 = filepath.ToSlash(paths1)

		t.Assert(tpath, paths1)

		t.Assert(tpath2, tpath)

		// 测试给定目录
		tpath2, err = ffile.Search(paths1, "testfiless")
		tpath2 = filepath.ToSlash(tpath2)
		tempss := filepath.ToSlash(paths1)
		t.Assert(tpath2, tempss)

		// 测试当前目录
		tempstr, _ = filepath.Abs("./")
		tempstr = testpath()
		paths1 = tempstr + ypaths1
		paths1 = filepath.ToSlash(paths1)

		t.Assert(tpath2, paths1)

		// 测试目录不存在时
		_, err = ffile.Search(paths2)
		t.AssertNE(err, nil)

	})
}
