package futil_test

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_Go(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var (
			wg      = sync.WaitGroup{}
			counter uint32
		)
		wg.Add(1)
		futil.Go(context.TODO(), func(ctx context.Context) {
			defer wg.Done()
			atomic.AddUint32(&counter, 1)
		}, nil)
		wg.Wait()
		t.Assert(atomic.LoadUint32(&counter), 1)
	})

	// recover
	ftest.C(t, func(t *ftest.T) {
		var (
			wg      = sync.WaitGroup{}
			counter uint32
		)
		wg.Add(1)
		futil.Go(context.TODO(), func(ctx context.Context) {
			defer wg.Done()
			panic("error")
			atomic.AddUint32(&counter, 1)
		}, nil)
		wg.Wait()
		t.Assert(atomic.LoadUint32(&counter), 0)
	})

	// catch error
	ftest.C(t, func(t *ftest.T) {
		var (
			wg  = sync.WaitGroup{}
			err error
		)
		wg.Add(1)
		futil.Go(context.TODO(), func(ctx context.Context) {
			panic("error")
		}, func(ctx context.Context, exception error) {
			defer wg.Done()
			err = exception
		})
		wg.Wait()
		t.Assert(err.Error(), "error")
	})
}
