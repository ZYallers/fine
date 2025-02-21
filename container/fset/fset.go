// Package fset provides kinds of concurrent-safe/unsafe sets.
package fset

import (
	"github.com/ZYallers/fine/internal/rwmutex"
	"github.com/ZYallers/fine/util/fconv"
)

// Set is consisted of interface{} items.
type Set struct {
	mu   rwmutex.RWMutex
	data map[interface{}]struct{}
}

// New create and returns a new set, which contains un-repeated items.
// The parameter `safe` is used to specify whether using set in concurrent-safety,
// which is false in default.
func New(safe ...bool) *Set {
	return NewSet(safe...)
}

// NewSet create and returns a new set, which contains un-repeated items.
// Also see New.
func NewSet(safe ...bool) *Set {
	return &Set{
		data: make(map[interface{}]struct{}),
		mu:   rwmutex.Create(safe...),
	}
}

// NewFrom returns a new set from `items`.
// Parameter `items` can be either a variable of any type, or a slice.
func NewFrom(items interface{}, safe ...bool) *Set {
	m := make(map[interface{}]struct{})
	for _, v := range fconv.Interfaces(items) {
		m[v] = struct{}{}
	}
	return &Set{
		data: m,
		mu:   rwmutex.Create(safe...),
	}
}
