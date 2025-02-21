package fvar

import (
	"github.com/ZYallers/fine/util/fconv"
)

// Scan automatically checks the type of `pointer` and converts `params` to `pointer`. It supports `pointer`
// with type of `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` for converting.
//
// See fconv.Scan.
func (v *Var) Scan(pointer interface{}, mapping ...map[string]string) error {
	return fconv.Scan(v.Val(), pointer, mapping...)
}
