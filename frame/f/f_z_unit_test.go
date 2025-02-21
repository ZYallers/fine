package f_test

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/ZYallers/fine/container/farray"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

var (
	ctx = context.TODO()
)

func Test_NewVar(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.NewVar(1).Int(), 1)
		t.Assert(f.NewVar(1, true).Int(), 1)
	})
}

func Test_Dump(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		f.Dump("Fine")
	})
}

func Test_DumpTo(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		f.DumpTo(os.Stdout, "Fine", futil.DumpOption{})
	})
}

func Test_DumpWithType(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		f.DumpWithType("Fine", 123)
	})
}

func Test_DumpWithOption(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		f.DumpWithOption("Fine", futil.DumpOption{})
	})
}

func Test_Try(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		f.Try(ctx, func(ctx context.Context) {
			f.Dump("Fine")
		})
	})
}

func Test_TryCatch(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		f.TryCatch(ctx, func(ctx context.Context) {
			f.Dump("Fine")
		}, func(ctx context.Context, exception error) {
			f.Dump(exception)
		})
	})
	ftest.C(t, func(t *ftest.T) {
		f.TryCatch(ctx, func(ctx context.Context) {
			f.Throw("Fine")
		}, func(ctx context.Context, exception error) {
			t.Assert(exception.Error(), "Fine")
		})
	})
}

func Test_IsNil(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.IsNil(nil), true)
		t.Assert(f.IsNil(0), false)
		t.Assert(f.IsNil("Fine"), false)
	})
}

func Test_IsEmpty(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(f.IsEmpty(nil), true)
		t.Assert(f.IsEmpty(0), true)
		t.Assert(f.IsEmpty("Fine"), false)
	})
}

func Test_Go(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = farray.NewArray(true)
		)
		wg.Add(1)
		f.Go(context.Background(), func(ctx context.Context) {
			defer wg.Done()
			array.Append(1)
		}, nil)
		wg.Wait()
		t.Assert(array.Len(), 1)
	})
}
