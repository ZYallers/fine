package fdebug_test

import (
	"testing"

	"github.com/ZYallers/fine/debug/fdebug"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/text/fstr"
)

func Test_CallerPackage(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(fdebug.CallerPackage(), "github.com/ZYallers/fine/test/ftest")
	})
}

func Test_CallerFunction(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(fdebug.CallerFunction(), "C")
	})
}

func Test_CallerFilePath(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(fstr.Contains(fdebug.CallerFilePath(), "ftest_util.go"), true)
	})
}

func Test_CallerDirectory(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(fstr.Contains(fdebug.CallerDirectory(), "ftest"), true)
	})
}

func Test_CallerFileLine(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Log(fdebug.CallerFileLine())
		t.Assert(fstr.Contains(fdebug.CallerFileLine(), "ftest_util.go:30"), true)
	})
}

func Test_CallerFileLineShort(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log(fdebug.CallerFileLineShort())
		t.Assert(fstr.Contains(fdebug.CallerFileLineShort(), "ftest_util.go:30"), true)
	})
}

func Test_FuncPath(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(fdebug.FuncPath(Test_FuncPath), "github.com/ZYallers/fine/debug/fdebug_test.Test_FuncPath")
	})
}

func Test_FuncName(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(fdebug.FuncName(Test_FuncName), "fdebug_test.Test_FuncName")
	})
}

func Test_PrintStack(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		fdebug.PrintStack()
	})
}

func Test_GoroutineId(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.AssertGT(fdebug.GoroutineId(), 0)
	})
}

func Test_Stack(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log(fdebug.Stack())
		t.Assert(fstr.Contains(fdebug.Stack(), "ftest_util.go:30"), true)
	})
}

func Test_StackWithFilter(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log(fdebug.StackWithFilter([]string{"github.com"}))
		t.Assert(fstr.Contains(fdebug.StackWithFilter([]string{"github.com"}), "ftest_util.go:30"), true)
	})
}
