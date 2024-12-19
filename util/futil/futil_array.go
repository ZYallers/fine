package futil

import (
	"github.com/ZYallers/fine/internal/util/utils"
)

// IsArray checks whether given value is array/slice.
// Note that it uses reflect internally implementing this feature.
func IsArray(value interface{}) bool {
	return utils.IsArray(value)
}
