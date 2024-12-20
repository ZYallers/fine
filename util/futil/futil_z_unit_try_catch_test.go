package futil_test

import (
	"context"
	"testing"

	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
	"github.com/pkg/errors"
)

func Test_Try(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		s := `Try test`
		err := futil.Try(context.TODO(), func(ctx context.Context) {
			panic(s)
		})
		t.Assert(err, s)
	})
	ftest.C(t, func(t *ftest.T) {
		s := `Try test`
		err := futil.Try(context.TODO(), func(ctx context.Context) {
			panic(errors.New(s))
		})
		t.Assert(err, s)
	})
}

func Test_TryCatch(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		futil.TryCatch(context.TODO(), func(ctx context.Context) {
			panic("TryCatch test")
		}, nil)
	})

	ftest.C(t, func(t *ftest.T) {
		futil.TryCatch(context.TODO(), func(ctx context.Context) {
			panic("TryCatch test")
		}, func(ctx context.Context, err error) {
			t.Assert(err, "TryCatch test")
		})
	})

	ftest.C(t, func(t *ftest.T) {
		futil.TryCatch(context.TODO(), func(ctx context.Context) {
			panic(errors.New("TryCatch test"))
		}, func(ctx context.Context, err error) {
			t.Assert(err, "TryCatch test")
		})
	})
}

func Test_Throw(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		defer func() {
			t.Assert(recover(), "Throw test")
		}()
		futil.Throw("Throw test")
	})
}
