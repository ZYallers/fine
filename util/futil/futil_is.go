package futil

import (
	"reflect"

	"github.com/ZYallers/fine/internal/empty"
)

// IsEmpty checks given `value` empty or not.
// It returns false if `value` is: integer(0), bool(false), slice/map(len=0), nil;
// or else returns true.
func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}

// IsTypeOf checks and returns whether the type of `value` and `valueInExpectType` equal.
func IsTypeOf(value, valueInExpectType interface{}) bool {
	return reflect.TypeOf(value) == reflect.TypeOf(valueInExpectType)
}
