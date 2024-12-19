package futil

import (
	"reflect"

	"github.com/ZYallers/fine/internal/reflection"
)

type (
	OriginValueAndKindOutput = reflection.OriginValueAndKindOutput
	OriginTypeAndKindOutput  = reflection.OriginTypeAndKindOutput
)

// OriginValueAndKind retrieves and returns the original reflect value and kind.
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput) {
	return reflection.OriginValueAndKind(value)
}

// OriginTypeAndKind retrieves and returns the original reflect type and kind.
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput) {
	return reflection.OriginTypeAndKind(value)
}

// CanCallIsNil Can reflect.Value call reflect.Value.IsNil.
// It can avoid reflect.Value.IsNil panics.
func CanCallIsNil(v interface{}) bool {
	rv, ok := v.(reflect.Value)
	if !ok {
		return false
	}
	switch rv.Kind() {
	case reflect.Interface, reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return true
	default:
		return false
	}
}
