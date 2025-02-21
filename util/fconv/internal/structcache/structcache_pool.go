package structcache

import (
	"sync"
)

var (
	poolUsedParamsKeyOrTagNameMap = &sync.Pool{
		New: func() interface{} {
			return make(map[string]struct{})
		},
	}
)

// GetUsedParamsKeyOrTagNameMapFromPool retrieves and returns a map for storing params key or tag name.
func GetUsedParamsKeyOrTagNameMapFromPool() map[string]struct{} {
	return poolUsedParamsKeyOrTagNameMap.Get().(map[string]struct{})
}

// PutUsedParamsKeyOrTagNameMapToPool puts a map for storing params key or tag name into pool for re-usage.
func PutUsedParamsKeyOrTagNameMapToPool(m map[string]struct{}) {
	// need to be cleared before putting back into pool.
	for k := range m {
		delete(m, k)
	}
	poolUsedParamsKeyOrTagNameMap.Put(m)
}
