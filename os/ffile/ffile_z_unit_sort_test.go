package ffile_test

import (
	"testing"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/text/fstr"
)

func Test_SortFiles(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		files := []string{
			"/aaa/bbb/ccc.txt",
			"/aaa/bbb/",
			"/aaa/",
			"/aaa",
			"/aaa/ccc/ddd.txt",
			"/bbb",
			"/0123",
			"/ddd",
			"/ccc",
		}
		sortOut := ffile.SortFiles(files)
		res := fstr.Join(sortOut, ",")
		t.AssertEQ(res, fstr.Join([]string{
			"/0123",
			"/aaa",
			"/aaa/",
			"/aaa/bbb/",
			"/aaa/bbb/ccc.txt",
			"/aaa/ccc/ddd.txt",
			"/bbb",
			"/ccc",
			"/ddd",
		}, ","))
	})
}
