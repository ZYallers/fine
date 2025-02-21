package ftype

// Any is a struct for concurrent-safe operation for type any.
type Any = Interface

// NewAny creates and returns a concurrent-safe object for any type,
// with given initial value `value`.
func NewAny(value ...interface{}) *Any {
	t := &Any{}
	if len(value) > 0 && value[0] != nil {
		t.value.Store(value[0])
	}
	return t
}
