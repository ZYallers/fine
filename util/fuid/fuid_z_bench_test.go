package fuid_test

import (
	"testing"

	"github.com/ZYallers/fine/util/fuid"
)

func Benchmark_S(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fuid.S()
		}
	})
}

func Benchmark_S_Data_1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fuid.S([]byte("123"))
		}
	})
}

func Benchmark_S_Data_2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fuid.S([]byte("123"), []byte("456"))
		}
	})
}
