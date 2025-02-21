package fvar

import (
	"github.com/ZYallers/fine/util/futil"
)

// ListItemValues retrieves and returns the elements of all item struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
func (v *Var) ListItemValues(key interface{}) (values []interface{}) {
	return futil.ListItemValues(v.Val(), key)
}

// ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
func (v *Var) ListItemValuesUnique(key string) []interface{} {
	return futil.ListItemValuesUnique(v.Val(), key)
}
