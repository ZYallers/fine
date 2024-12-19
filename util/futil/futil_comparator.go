package futil

import (
	"strings"

	"github.com/ZYallers/fine/util/fconv"
)

// Comparator is a function that compare a and b, and returns the result as int.
//
// Should return a number:
//
//	negative , if a < b
//	zero     , if a == b
//	positive , if a > b
type Comparator func(a, b interface{}) int

// ComparatorString provides a fast comparison on strings.
func ComparatorString(a, b interface{}) int {
	return strings.Compare(fconv.String(a), fconv.String(b))
}

// ComparatorInt provides a basic comparison on int.
func ComparatorInt(a, b interface{}) int {
	return fconv.Int(a) - fconv.Int(b)
}

// ComparatorInt8 provides a basic comparison on int8.
func ComparatorInt8(a, b interface{}) int {
	return int(fconv.Int8(a) - fconv.Int8(b))
}

// ComparatorInt16 provides a basic comparison on int16.
func ComparatorInt16(a, b interface{}) int {
	return int(fconv.Int16(a) - fconv.Int16(b))
}

// ComparatorInt32 provides a basic comparison on int32.
func ComparatorInt32(a, b interface{}) int {
	return int(fconv.Int32(a) - fconv.Int32(b))
}

// ComparatorInt64 provides a basic comparison on int64.
func ComparatorInt64(a, b interface{}) int {
	return int(fconv.Int64(a) - fconv.Int64(b))
}

// ComparatorUint provides a basic comparison on uint.
func ComparatorUint(a, b interface{}) int {
	return int(fconv.Uint(a) - fconv.Uint(b))
}

// ComparatorUint8 provides a basic comparison on uint8.
func ComparatorUint8(a, b interface{}) int {
	return int(fconv.Uint8(a) - fconv.Uint8(b))
}

// ComparatorUint16 provides a basic comparison on uint16.
func ComparatorUint16(a, b interface{}) int {
	return int(fconv.Uint16(a) - fconv.Uint16(b))
}

// ComparatorUint32 provides a basic comparison on uint32.
func ComparatorUint32(a, b interface{}) int {
	return int(fconv.Uint32(a) - fconv.Uint32(b))
}

// ComparatorUint64 provides a basic comparison on uint64.
func ComparatorUint64(a, b interface{}) int {
	return int(fconv.Uint64(a) - fconv.Uint64(b))
}

// ComparatorFloat32 provides a basic comparison on float32.
func ComparatorFloat32(a, b interface{}) int {
	aFloat := fconv.Float32(a)
	bFloat := fconv.Float32(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// ComparatorFloat64 provides a basic comparison on float64.
func ComparatorFloat64(a, b interface{}) int {
	aFloat := fconv.Float64(a)
	bFloat := fconv.Float64(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// ComparatorByte provides a basic comparison on byte.
func ComparatorByte(a, b interface{}) int {
	return int(fconv.Byte(a) - fconv.Byte(b))
}

// ComparatorRune provides a basic comparison on rune.
func ComparatorRune(a, b interface{}) int {
	return int(fconv.Rune(a) - fconv.Rune(b))
}

// ComparatorTime provides a basic comparison on time.Time.
func ComparatorTime(a, b interface{}) int {
	aTime := fconv.Time(a)
	bTime := fconv.Time(b)
	switch {
	case aTime.After(bTime):
		return 1
	case aTime.Before(bTime):
		return -1
	default:
		return 0
	}
}
