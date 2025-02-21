package ferror_test

import (
	"errors"
	"testing"

	"github.com/ZYallers/fine/errors/ferror"
	"github.com/ZYallers/fine/test/ftest"
)

func TestError_New(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		err := ferror.New(200, "success")
		t.Assert(err.Code(), 200)
		t.Assert(err.Msg(), "success")
	})
	ftest.C(t, func(t *ftest.T) {
		err := ferror.New(400, errors.New("not found"))
		t.Assert(err.Code(), 400)
		t.Assert(err.Msg(), "not found")
	})
}

func TestError_Code(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		err := ferror.New(501, "bad request")
		t.Assert(err.Code(), 501)
	})
}

func TestError_Msg(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		err := ferror.New(404, "not found")
		t.Assert(err.Msg(), "not found")
	})
}

func TestError_Error(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		err := ferror.New(404, "not found")
		t.Assert(err.Error(), "404:not found")
	})
	ftest.C(t, func(t *ftest.T) {
		var err *ferror.Error
		t.Assert(err.Error(), "")
	})
}
