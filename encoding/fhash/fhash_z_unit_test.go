package fhash_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fhash"
	"github.com/ZYallers/fine/test/ftest"
)

var (
	strBasic = []byte("This is the test string for hash.")
)

func Test_AP(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log("AP", fhash.AP(strBasic))
		t.Assert(fhash.AP(strBasic), uint32(3998202516))
	})
	ftest.C(t, func(t *ftest.T) {
		// t.Log("AP64", fhash.AP64(strBasic))
		t.Assert(fhash.AP64(strBasic), uint64(2531023058543352243))
	})
}

func Test_BKDR(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log("BKDR", fhash.BKDR(strBasic))
		t.Assert(fhash.BKDR(strBasic), uint32(200645773))
	})
	ftest.C(t, func(t *ftest.T) {
		// t.Log("BKDR64", fhash.BKDR64(strBasic))
		t.Assert(fhash.BKDR64(strBasic), uint64(4214762819217104013))
	})
}

func Test_DJB(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log("DJB", fhash.DJB(strBasic))
		t.Assert(fhash.DJB(strBasic), uint32(959862602))
	})
	ftest.C(t, func(t *ftest.T) {
		// t.Log("DJB64", fhash.DJB64(strBasic))
		t.Assert(fhash.DJB64(strBasic), uint64(2519720351310960458))
	})
}

func Test_ELF(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log("ELF", fhash.ELF(strBasic))
		t.Assert(fhash.ELF(strBasic), uint32(7244206))
	})
	ftest.C(t, func(t *ftest.T) {
		// t.Log("ELF64", fhash.ELF64(strBasic))
		t.Assert(fhash.ELF64(strBasic), uint64(31150))
	})
}

func Test_JS(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log("JS", fhash.JS(strBasic))
		t.Assert(fhash.JS(strBasic), uint32(498688898))
	})
	ftest.C(t, func(t *ftest.T) {
		// t.Log("JS64", fhash.JS64(strBasic))
		t.Assert(fhash.JS64(strBasic), uint64(13410163655098759877))
	})
}

func Test_PJW(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log("PJW", fhash.PJW(strBasic))
		t.Assert(fhash.PJW(strBasic), uint32(7244206))
	})
	ftest.C(t, func(t *ftest.T) {
		// t.Log("PJW64", fhash.PJW64(strBasic))
		t.Assert(fhash.PJW64(strBasic), uint64(31150))
	})
}

func Test_RS(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log("RS", fhash.RS(strBasic))
		t.Assert(fhash.RS(strBasic), uint32(1944033799))
	})
	ftest.C(t, func(t *ftest.T) {
		// t.Log("RS64", fhash.RS64(strBasic))
		t.Assert(fhash.RS64(strBasic), uint64(13439708950444349959))
	})
}

func Test_SDBM(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		// t.Log("SDBM", fhash.SDBM(strBasic))
		t.Assert(fhash.SDBM(strBasic), uint32(1069170245))
	})
	ftest.C(t, func(t *ftest.T) {
		// t.Log("SDBM64", fhash.SDBM64(strBasic))
		t.Assert(fhash.SDBM64(strBasic), uint64(9881052176572890693))
	})
}
