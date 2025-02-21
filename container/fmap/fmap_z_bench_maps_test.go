package fmap_test

import (
	"testing"

	"github.com/ZYallers/fine/container/fmap"
)

var hashMap = fmap.New(true)

var listMap = fmap.NewListMap(true)

func Benchmark_HashMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			hashMap.Set(i, i)
			i++
		}
	})
}

func Benchmark_ListMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			listMap.Set(i, i)
			i++
		}
	})
}

func Benchmark_HashMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			hashMap.Get(i)
			i++
		}
	})
}

func Benchmark_ListMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			listMap.Get(i)
			i++
		}
	})
}
