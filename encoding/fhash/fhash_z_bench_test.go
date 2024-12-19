package fhash_test

import (
	"testing"

	"github.com/ZYallers/fine/encoding/fhash"
)

var (
	str = []byte("This is the test string for hash.")
)

func Benchmark_AP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.AP(str)
	}
}
func Benchmark_AP64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.AP64(str)
	}
}

func Benchmark_BKDR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.BKDR(str)
	}
}
func Benchmark_BKDR64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.BKDR64(str)
	}
}

func Benchmark_DJB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.DJB(str)
	}
}
func Benchmark_DJB64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.DJB64(str)
	}
}

func Benchmark_ELF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.ELF(str)
	}
}
func Benchmark_ELF64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.ELF64(str)
	}
}

func Benchmark_JS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.JS(str)
	}
}
func Benchmark_JS64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.JS64(str)
	}
}

func Benchmark_PJW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.PJW(str)
	}
}
func Benchmark_PJW64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.PJW64(str)
	}
}

func Benchmark_RS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.RS(str)
	}
}
func Benchmark_RS64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.RS64(str)
	}
}

func Benchmark_SDBM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.SDBM(str)
	}
}
func Benchmark_SDBM64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fhash.SDBM64(str)
	}
}
