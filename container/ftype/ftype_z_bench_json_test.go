package ftype_test

import (
	"testing"

	"github.com/ZYallers/fine/container/ftype"
	"github.com/ZYallers/fine/internal/json"
)

var (
	vBool      = ftype.NewBool()
	vByte      = ftype.NewByte()
	vBytes     = ftype.NewBytes()
	vFloat32   = ftype.NewFloat32()
	vFloat64   = ftype.NewFloat64()
	vInt       = ftype.NewInt()
	vInt32     = ftype.NewInt32()
	vInt64     = ftype.NewInt64()
	vInterface = ftype.NewInterface()
	vString    = ftype.NewString()
	vUint      = ftype.NewUint()
	vUint32    = ftype.NewUint32()
	vUint64    = ftype.NewUint64()
)

func Benchmark_Bool_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vBool)
	}
}

func Benchmark_Byte_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vByte)
	}
}

func Benchmark_Bytes_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vBytes)
	}
}

func Benchmark_Float32_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vFloat32)
	}
}

func Benchmark_Float64_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vFloat64)
	}
}

func Benchmark_Int_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInt)
	}
}

func Benchmark_Int32_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInt32)
	}
}

func Benchmark_Int64_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInt64)
	}
}

func Benchmark_Interface_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInterface)
	}
}

func Benchmark_String_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vString)
	}
}

func Benchmark_Uint_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vUint)
	}
}

func Benchmark_Uint32_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vUint64)
	}
}
