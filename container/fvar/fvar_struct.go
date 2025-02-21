package fvar

import (
	"github.com/ZYallers/fine/util/fconv"
)

// Struct maps value of `v` to `pointer`.
// The parameter `pointer` should be a pointer to a struct instance.
// The parameter `mapping` is used to specify the key-to-attribute mapping rules.
func (v *Var) Struct(pointer interface{}, mapping ...map[string]string) error {
	return fconv.Struct(v.Val(), pointer, mapping...)
}

// Structs converts and returns `v` as given struct slice.
func (v *Var) Structs(pointer interface{}, mapping ...map[string]string) error {
	return fconv.Structs(v.Val(), pointer, mapping...)
}
