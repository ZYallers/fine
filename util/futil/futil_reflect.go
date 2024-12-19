package futil

import (
	"github.com/ZYallers/fine/internal/util/utils"
)

// CanCallIsNil Can reflect.Value call reflect.Value.IsNil.
// It can avoid reflect.Value.IsNil panics.
func CanCallIsNil(v interface{}) bool {
	return utils.CanCallIsNil(v)
}
