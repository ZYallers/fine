package flist_test

import (
	"testing"

	"github.com/ZYallers/fine/container/flist"
)

var (
	l = flist.New(true)
)

func Benchmark_PushBack(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			l.PushBack(i)
			i++
		}
	})
}

func Benchmark_PushFront(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			l.PushFront(i)
			i++
		}
	})
}

func Benchmark_Len(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Len()
		}
	})
}

func Benchmark_PopFront(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.PopFront()
		}
	})
}

func Benchmark_PopBack(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.PopBack()
		}
	})
}
