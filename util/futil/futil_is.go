package futil

import (
	"github.com/ZYallers/fine/internal/util/utils"
)

// IsNil checks whether `value` is nil, especially for interface{} type value.
func IsNil(value interface{}) bool {
	return utils.IsNil(value)
}

// IsEmpty checks whether `value` is empty.
func IsEmpty(value interface{}) bool {
	return utils.IsEmpty(value)
}

// IsInt checks whether `value` is type of int.
func IsInt(value interface{}) bool {
	return utils.IsInt(value)
}

// IsUint checks whether `value` is type of uint.
func IsUint(value interface{}) bool {
	return utils.IsUint(value)
}

// IsFloat checks whether `value` is type of float.
func IsFloat(value interface{}) bool {
	return utils.IsFloat(value)
}

// IsSlice checks whether `value` is type of slice.
func IsSlice(value interface{}) bool {
	return utils.IsSlice(value)
}

// IsMap checks whether `value` is type of map.
func IsMap(value interface{}) bool {
	return utils.IsMap(value)
}

// IsStruct checks whether `value` is type of struct.
func IsStruct(value interface{}) bool {
	return utils.IsStruct(value)
}
