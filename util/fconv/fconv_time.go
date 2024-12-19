package fconv

import (
	"time"

	"github.com/ZYallers/fine/os/ftime"
	"github.com/ZYallers/fine/util/fconv/internal/localinterface"
	"github.com/ZYallers/fine/util/futil"
)

// Time converts `any` to time.Time.
func Time(any interface{}, format ...string) time.Time {
	// It's already this type.
	if len(format) == 0 {
		if v, ok := any.(time.Time); ok {
			return v
		}
	}
	if t := FTime(any, format...); t != nil {
		return t.Time
	}
	return time.Time{}
}

// Duration converts `any` to time.Duration.
// If `any` is string, then it uses time.ParseDuration to convert it.
// If `any` is numeric, then it converts `any` as nanoseconds.
func Duration(any interface{}) time.Duration {
	// It's already this type.
	if v, ok := any.(time.Duration); ok {
		return v
	}
	s := String(any)
	if !futil.IsNumeric(s) {
		d, _ := ftime.ParseDuration(s)
		return d
	}
	return time.Duration(Int64(any))
}

// FTime converts `any` to *ftime.Time.
// The parameter `format` can be used to specify the format of `any`.
// It returns the converted value that matched the first format of the formats slice.
// If no `format` given, it converts `any` using ftime.NewFromTimeStamp if `any` is numeric,
// or using ftime.StrToTime if `any` is string.
func FTime(any interface{}, format ...string) *ftime.Time {
	if any == nil {
		return nil
	}
	if v, ok := any.(localinterface.IFTime); ok {
		return v.FTime(format...)
	}
	// It's already this type.
	if len(format) == 0 {
		if v, ok := any.(*ftime.Time); ok {
			return v
		}
		if t, ok := any.(time.Time); ok {
			return ftime.New(t)
		}
		if t, ok := any.(*time.Time); ok {
			return ftime.New(t)
		}
	}
	s := String(any)
	if len(s) == 0 {
		return ftime.New()
	}
	// Priority conversion using given format.
	if len(format) > 0 {
		for _, item := range format {
			t, err := ftime.StrToTimeFormat(s, item)
			if t != nil && err == nil {
				return t
			}
		}
		return nil
	}
	if futil.IsNumeric(s) {
		return ftime.NewFromTimeStamp(Int64(s))
	} else {
		t, _ := ftime.StrToTime(s)
		return t
	}
}
