package ffile_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/os/ftime"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/fconv"
	"github.com/ZYallers/fine/util/fuid"
)

func Test_IsDir(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		paths := "/testfile"
		createDir(paths)
		defer delTestFiles(paths)

		t.Assert(ffile.IsDir(testpath()+paths), true)
		t.Assert(ffile.IsDir("./testfile2"), false)
		t.Assert(ffile.IsDir("./testfile/tt.txt"), false)
		t.Assert(ffile.IsDir(""), false)
	})
}

func Test_IsEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		path := "/testdir_" + fconv.String(ftime.TimestampNano())
		createDir(path)
		defer delTestFiles(path)

		t.Assert(ffile.IsEmpty(testpath()+path), true)
		t.Assert(ffile.IsEmpty(testpath()+path+ffile.Separator+"test.txt"), true)
	})
	ftest.C(t, func(t *ftest.T) {
		path := "/testfile_" + fconv.String(ftime.TimestampNano())
		createTestFile(path, "")
		defer delTestFiles(path)

		t.Assert(ffile.IsEmpty(testpath()+path), true)
	})
	ftest.C(t, func(t *ftest.T) {
		path := "/testfile_" + fconv.String(ftime.TimestampNano())
		createTestFile(path, "1")
		defer delTestFiles(path)

		t.Assert(ffile.IsEmpty(testpath()+path), false)
	})
}

func Test_Create(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			err       error
			filepaths []string
			fileobj   *os.File
		)
		filepaths = append(filepaths, "/testfile_cc1.txt")
		filepaths = append(filepaths, "/testfile_cc2.txt")
		for _, v := range filepaths {
			fileobj, err = ffile.Create(testpath() + v)
			defer delTestFiles(v)
			fileobj.Close()
			t.AssertNil(err)
		}
	})

	ftest.C(t, func(t *ftest.T) {
		tmpPath := ffile.Join(ffile.Temp(), "test/testfile_cc1.txt")
		fileobj, err := ffile.Create(tmpPath)
		defer ffile.Remove(tmpPath)
		t.AssertNE(fileobj, nil)
		t.AssertNil(err)
		fileobj.Close()
	})
}

func Test_Open(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)

		file1 := "/testfile_nc1.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)

		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "./testfile/file1/c1.txt")
		flags = append(flags, false)

		for k, v := range files {
			fileobj, err = ffile.Open(testpath() + v)
			fileobj.Close()
			if flags[k] {
				t.AssertNil(err)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_OpenFile(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)

		files = append(files, "./testfile/file1/nc1.txt")
		flags = append(flags, false)

		f1 := "/testfile_tt.txt"
		createTestFile(f1, "")
		defer delTestFiles(f1)

		files = append(files, f1)
		flags = append(flags, true)

		for k, v := range files {
			fileobj, err = ffile.OpenFile(testpath()+v, os.O_RDWR, 0666)
			fileobj.Close()
			if flags[k] {
				t.AssertNil(err)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_OpenWithFlag(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)

		file1 := "/testfile_t1.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)
		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "/testfiless/dirfiles/t1_no.txt")
		flags = append(flags, false)

		for k, v := range files {
			fileobj, err = ffile.OpenWithFlag(testpath()+v, os.O_RDWR)
			fileobj.Close()
			if flags[k] {
				t.AssertNil(err)
			} else {
				t.AssertNE(err, nil)
			}
		}
	})
}

func Test_OpenWithFlagPerm(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)
		file1 := "/testfile_nc1.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)
		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "/testfileyy/tt.txt")
		flags = append(flags, false)

		for k, v := range files {
			fileobj, err = ffile.OpenWithFlagPerm(testpath()+v, os.O_RDWR, 0666)
			fileobj.Close()
			if flags[k] {
				t.AssertNil(err)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_Exists(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			flag  bool
			files []string
			flags []bool
		)

		file1 := "/testfile_GetContents.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)

		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "./testfile/havefile1/tt_no.txt")
		flags = append(flags, false)

		for k, v := range files {
			flag = ffile.Exists(testpath() + v)
			if flags[k] {
				t.Assert(flag, true)
			} else {
				t.Assert(flag, false)
			}
		}
	})
}

func Test_Pwd(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		paths, err := os.Getwd()
		t.AssertNil(err)
		t.Assert(ffile.Pwd(), paths)

	})
}

func Test_IsFile(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			flag  bool
			files []string
			flags []bool
		)

		file1 := "/testfile_tt.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)
		files = append(files, file1)
		flags = append(flags, true)

		dir1 := "/testfiless"
		createDir(dir1)
		defer delTestFiles(dir1)
		files = append(files, dir1)
		flags = append(flags, false)

		files = append(files, "./testfiledd/tt1.txt")
		flags = append(flags, false)

		for k, v := range files {
			flag = ffile.IsFile(testpath() + v)
			if flags[k] {
				t.Assert(flag, true)
			} else {
				t.Assert(flag, false)
			}

		}
	})
}

func Test_Info(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			err    error
			paths  string = "/testfile_t1.txt"
			files  os.FileInfo
			files2 os.FileInfo
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)
		files, err = ffile.Stat(testpath() + paths)
		t.AssertNil(err)

		files2, err = os.Stat(testpath() + paths)
		t.AssertNil(err)
		t.Assert(files, files2)
	})
}

func Test_Move(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths     string = "/ovetest"
			filepaths string = "/testfile_ttn1.txt"
			topath    string = "/testfile_ttn2.txt"
		)
		createDir("/ovetest")
		createTestFile(paths+filepaths, "a")

		defer delTestFiles(paths)

		yfile := testpath() + paths + filepaths
		tofile := testpath() + paths + topath

		t.Assert(ffile.Move(yfile, tofile), nil)

		// 检查移动后的文件是否真实存在
		_, err := os.Stat(tofile)
		t.Assert(os.IsNotExist(err), false)
	})
}

func Test_Rename(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths  string = "/testfiles"
			ypath  string = "/testfilettm1.txt"
			topath string = "/testfilettm2.txt"
		)
		createDir(paths)
		createTestFile(paths+ypath, "a")
		defer delTestFiles(paths)

		ypath = testpath() + paths + ypath
		topath = testpath() + paths + topath

		t.Assert(ffile.Rename(ypath, topath), nil)
		t.Assert(ffile.IsFile(topath), true)

		t.AssertNE(ffile.Rename("", ""), nil)

	})

}

func Test_DirNames(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths    string = "/testdirs"
			err      error
			readlist []string
		)
		havelist := []string{
			"t1.txt",
			"t2.txt",
		}

		// 创建测试文件
		createDir(paths)
		for _, v := range havelist {
			createTestFile(paths+"/"+v, "")
		}
		defer delTestFiles(paths)

		readlist, err = ffile.DirNames(testpath() + paths)

		t.AssertNil(err)
		t.AssertIN(readlist, havelist)

		_, err = ffile.DirNames("")
		t.AssertNE(err, nil)

	})
}

func Test_Glob(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths      string = "/testfiles/*.txt"
			dirpath    string = "/testfiles"
			err        error
			resultlist []string
		)

		havelist1 := []string{
			"t1.txt",
			"t2.txt",
		}

		havelist2 := []string{
			testpath() + "/testfiles/t1.txt",
			testpath() + "/testfiles/t2.txt",
		}

		// ===============================构建测试文件
		createDir(dirpath)
		for _, v := range havelist1 {
			createTestFile(dirpath+"/"+v, "")
		}
		defer delTestFiles(dirpath)

		resultlist, err = ffile.Glob(testpath()+paths, true)
		t.AssertNil(err)
		t.Assert(resultlist, havelist1)

		resultlist, err = ffile.Glob(testpath()+paths, false)

		t.AssertNil(err)
		t.Assert(formatpaths(resultlist), formatpaths(havelist2))

		_, err = ffile.Glob("", true)
		t.AssertNil(err)

	})
}

func Test_Remove(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths string = "/testfile_t1.txt"
		)
		createTestFile(paths, "")
		t.Assert(ffile.Remove(testpath()+paths), nil)

		t.Assert(ffile.Remove(""), nil)

		defer delTestFiles(paths)

	})
}

func Test_IsReadable(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1 string = "/testfile_GetContents.txt"
			paths2 string = "./testfile_GetContents_no.txt"
		)

		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		t.Assert(ffile.IsReadable(testpath()+paths1), true)
		t.Assert(ffile.IsReadable(paths2), false)

	})
}

func Test_IsWritable(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1 string = "/testfile_GetContents.txt"
			paths2 string = "./testfile_GetContents_no.txt"
		)

		createTestFile(paths1, "")
		defer delTestFiles(paths1)
		t.Assert(ffile.IsWritable(testpath()+paths1), true)
		t.Assert(ffile.IsWritable(paths2), false)

	})
}

func Test_Chmod(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1 string = "/testfile_GetContents.txt"
			paths2 string = "./testfile_GetContents_no.txt"
		)
		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		t.Assert(ffile.Chmod(testpath()+paths1, 0777), nil)
		t.AssertNE(ffile.Chmod(paths2, 0777), nil)

	})
}

// 获取绝对目录地址
func Test_RealPath(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1    string = "/testfile_files"
			readlPath string

			tempstr string
		)

		createDir(paths1)
		defer delTestFiles(paths1)

		readlPath = ffile.RealPath("./")

		tempstr, _ = filepath.Abs("./")

		t.Assert(readlPath, tempstr)

		t.Assert(ffile.RealPath("./nodirs"), "")

	})
}

// 获取当前执行文件的目录
func Test_SelfPath(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1    string
			readlPath string
			tempstr   string
		)
		readlPath = ffile.SelfPath()
		readlPath = filepath.ToSlash(readlPath)

		tempstr, _ = filepath.Abs(os.Args[0])
		paths1 = filepath.ToSlash(tempstr)
		paths1 = strings.Replace(paths1, "./", "/", 1)

		t.Assert(readlPath, paths1)

	})
}

func Test_SelfDir(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1    string
			readlPath string
			tempstr   string
		)
		readlPath = ffile.SelfDir()

		tempstr, _ = filepath.Abs(os.Args[0])
		paths1 = filepath.Dir(tempstr)

		t.Assert(readlPath, paths1)

	})
}

func Test_Basename(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1    string = "/testfilerr_GetContents.txt"
			readlPath string
		)

		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		readlPath = ffile.Basename(testpath() + paths1)
		t.Assert(readlPath, "testfilerr_GetContents.txt")

	})
}

func Test_Dir(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1    string = "/testfiless"
			readlPath string
		)
		createDir(paths1)
		defer delTestFiles(paths1)

		readlPath = ffile.Dir(testpath() + paths1)

		t.Assert(readlPath, testpath())

		t.Assert(len(ffile.Dir(".")) > 0, true)
	})
}

func Test_Ext(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			paths1   string = "/testfile_GetContents.txt"
			dirpath1        = "/testdirs"
		)
		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		createDir(dirpath1)
		defer delTestFiles(dirpath1)

		t.Assert(ffile.Ext(testpath()+paths1), ".txt")
		t.Assert(ffile.Ext(testpath()+dirpath1), "")
	})

	ftest.C(t, func(t *ftest.T) {
		t.Assert(ffile.Ext("/var/www/test.js"), ".js")
		t.Assert(ffile.Ext("/var/www/test.min.js"), ".js")
		t.Assert(ffile.Ext("/var/www/test.js?1"), ".js")
		t.Assert(ffile.Ext("/var/www/test.min.js?v1"), ".js")
	})
}

func Test_ExtName(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(ffile.ExtName("/var/www/test.js"), "js")
		t.Assert(ffile.ExtName("/var/www/test.min.js"), "js")
		t.Assert(ffile.ExtName("/var/www/test.js?v=1"), "js")
		t.Assert(ffile.ExtName("/var/www/test.min.js?v=1"), "js")
	})
}

func Test_TempDir(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(ffile.Temp(), os.TempDir())
	})
}

func Test_Mkdir(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			tpath string = "/testfile/createdir"
			err   error
		)

		defer delTestFiles("/testfile")

		err = ffile.Mkdir(testpath() + tpath)
		t.AssertNil(err)

		err = ffile.Mkdir("")
		t.AssertNE(err, nil)

		err = ffile.Mkdir(testpath() + tpath + "2/t1")
		t.AssertNil(err)

	})
}

func Test_Stat(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			tpath1   = "/testfile_t1.txt"
			tpath2   = "./testfile_t1_no.txt"
			err      error
			fileiofo os.FileInfo
		)

		createTestFile(tpath1, "a")
		defer delTestFiles(tpath1)

		fileiofo, err = ffile.Stat(testpath() + tpath1)
		t.AssertNil(err)

		t.Assert(fileiofo.Size(), 1)

		_, err = ffile.Stat(tpath2)
		t.AssertNE(err, nil)

	})
}

func Test_MainPkgPath(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		reads := ffile.MainPkgPath()
		t.Assert(reads, "")
	})
}

func Test_SelfName(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(len(ffile.SelfName()) > 0, true)
	})
}

func Test_MTimestamp(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(ffile.MTimestamp(ffile.Temp()) > 0, true)
	})
}

func Test_RemoveFile_RemoveAll(t *testing.T) {
	// safe deleting single file.
	ftest.C(t, func(t *ftest.T) {
		path := ffile.Temp(fuid.S())
		err := ffile.PutContents(path, "1")
		t.AssertNil(err)
		t.Assert(ffile.Exists(path), true)

		err = ffile.RemoveFile(path)
		t.AssertNil(err)
		t.Assert(ffile.Exists(path), false)
	})
	// error deleting dir which is not empty.
	ftest.C(t, func(t *ftest.T) {
		var (
			err       error
			dirPath   = ffile.Temp(fuid.S())
			filePath1 = ffile.Join(dirPath, fuid.S())
			filePath2 = ffile.Join(dirPath, fuid.S())
		)
		err = ffile.PutContents(filePath1, "1")
		t.AssertNil(err)
		t.Assert(ffile.Exists(filePath1), true)

		err = ffile.PutContents(filePath2, "2")
		t.AssertNil(err)
		t.Assert(ffile.Exists(filePath2), true)

		err = ffile.RemoveFile(dirPath)
		t.AssertNE(err, nil)
		t.Assert(ffile.Exists(filePath1), true)
		t.Assert(ffile.Exists(filePath2), true)

		err = ffile.RemoveAll(dirPath)
		t.AssertNil(err)
		t.Assert(ffile.Exists(filePath1), false)
		t.Assert(ffile.Exists(filePath2), false)
	})
}
