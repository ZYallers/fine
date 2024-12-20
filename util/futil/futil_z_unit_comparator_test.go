package futil_test

import (
	"testing"

	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_ComparatorString(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorString(1, 1), 0)
		t.Assert(futil.ComparatorString(1, 2), -1)
		t.Assert(futil.ComparatorString(2, 1), 1)
	})
}

func Test_ComparatorInt(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorInt(1, 1), 0)
		t.Assert(futil.ComparatorInt(1, 2), -1)
		t.Assert(futil.ComparatorInt(2, 1), 1)
	})
}

func Test_ComparatorInt8(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorInt8(1, 1), 0)
		t.Assert(futil.ComparatorInt8(1, 2), -1)
		t.Assert(futil.ComparatorInt8(2, 1), 1)
	})
}

func Test_ComparatorInt16(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorInt16(1, 1), 0)
		t.Assert(futil.ComparatorInt16(1, 2), -1)
		t.Assert(futil.ComparatorInt16(2, 1), 1)
	})
}

func Test_ComparatorInt32(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorInt32(1, 1), 0)
		t.Assert(futil.ComparatorInt32(1, 2), -1)
		t.Assert(futil.ComparatorInt32(2, 1), 1)
	})
}

func Test_ComparatorInt64(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorInt64(1, 1), 0)
		t.Assert(futil.ComparatorInt64(1, 2), -1)
		t.Assert(futil.ComparatorInt64(2, 1), 1)
	})
}

func Test_ComparatorUint(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorUint(1, 1), 0)
		t.Assert(futil.ComparatorUint(1, 2), -1)
		t.Assert(futil.ComparatorUint(2, 1), 1)
	})
}

func Test_ComparatorUint8(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorUint8(1, 1), 0)
		t.Assert(futil.ComparatorUint8(2, 6), 252)
		t.Assert(futil.ComparatorUint8(2, 1), 1)
	})
}

func Test_ComparatorUint16(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorUint16(1, 1), 0)
		t.Assert(futil.ComparatorUint16(1, 2), 65535)
		t.Assert(futil.ComparatorUint16(2, 1), 1)
	})
}

func Test_ComparatorUint32(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorUint32(1, 1), 0)
		t.Assert(futil.ComparatorUint32(-1000, 2147483640), 2147482656)
		t.Assert(futil.ComparatorUint32(2, 1), 1)
	})
}

func Test_ComparatorUint64(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorUint64(1, 1), 0)
		t.Assert(futil.ComparatorUint64(1, 2), -1)
		t.Assert(futil.ComparatorUint64(2, 1), 1)
	})
}

func Test_ComparatorFloat32(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorFloat32(1, 1), 0)
		t.Assert(futil.ComparatorFloat32(1, 2), -1)
		t.Assert(futil.ComparatorFloat32(2, 1), 1)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorFloat32(0.1, 0.1), 0)
		t.Assert(futil.ComparatorFloat32(1.1, 2.1), -1)
		t.Assert(futil.ComparatorFloat32(2.1, 1.1), 1)
	})
}

func Test_ComparatorFloat64(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorFloat64(1, 1), 0)
		t.Assert(futil.ComparatorFloat64(1, 2), -1)
		t.Assert(futil.ComparatorFloat64(2, 1), 1)
	})
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorFloat64(0.1, 0.1), 0)
		t.Assert(futil.ComparatorFloat64(1.1, 2.1), -1)
		t.Assert(futil.ComparatorFloat64(2.1, 1.1), 1)
	})
}

func Test_ComparatorByte(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorByte(1, 1), 0)
		t.Assert(futil.ComparatorByte(1, 2), 255)
		t.Assert(futil.ComparatorByte(2, 1), 1)
	})
}

func Test_ComparatorRune(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorRune(1, 1), 0)
		t.Assert(futil.ComparatorRune(1, 2), -1)
		t.Assert(futil.ComparatorRune(2, 1), 1)
	})
}

func Test_ComparatorTime(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		t.Assert(futil.ComparatorTime("2019-06-14", "2019-06-14"), 0)
		t.Assert(futil.ComparatorTime("2019-06-15", "2019-06-14"), 1)
		t.Assert(futil.ComparatorTime("2019-06-13", "2019-06-14"), -1)
	})
}
