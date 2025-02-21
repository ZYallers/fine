package fmap_test

import (
	"strconv"
	"testing"

	"github.com/ZYallers/fine/container/fmap"
)

var anyAnyMapUnsafe = fmap.New()

var intIntMapUnsafe = fmap.NewIntIntMap()

var intAnyMapUnsafe = fmap.NewIntAnyMap()

var intStrMapUnsafe = fmap.NewIntStrMap()

var strIntMapUnsafe = fmap.NewStrIntMap()

var strAnyMapUnsafe = fmap.NewStrAnyMap()

var strStrMapUnsafe = fmap.NewStrStrMap()

// Writing benchmarks.

func Benchmark_Unsafe_IntIntMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intIntMapUnsafe.Set(i, i)
	}
}

func Benchmark_Unsafe_IntAnyMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intAnyMapUnsafe.Set(i, i)
	}
}

func Benchmark_Unsafe_IntStrMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intStrMapUnsafe.Set(i, strconv.Itoa(i))
	}
}

func Benchmark_Unsafe_AnyAnyMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anyAnyMapUnsafe.Set(i, i)
	}
}

func Benchmark_Unsafe_StrIntMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strIntMapUnsafe.Set(strconv.Itoa(i), i)
	}
}

func Benchmark_Unsafe_StrAnyMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strAnyMapUnsafe.Set(strconv.Itoa(i), i)
	}
}

func Benchmark_Unsafe_StrStrMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStrMapUnsafe.Set(strconv.Itoa(i), strconv.Itoa(i))
	}
}

// Reading benchmarks.

func Benchmark_Unsafe_IntIntMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intIntMapUnsafe.Get(i)
	}
}

func Benchmark_Unsafe_IntAnyMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intAnyMapUnsafe.Get(i)
	}
}

func Benchmark_Unsafe_IntStrMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intStrMapUnsafe.Get(i)
	}
}

func Benchmark_Unsafe_AnyAnyMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anyAnyMapUnsafe.Get(i)
	}
}

func Benchmark_Unsafe_StrIntMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strIntMapUnsafe.Get(strconv.Itoa(i))
	}
}

func Benchmark_Unsafe_StrAnyMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strAnyMapUnsafe.Get(strconv.Itoa(i))
	}
}

func Benchmark_Unsafe_StrStrMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStrMapUnsafe.Get(strconv.Itoa(i))
	}
}
