package fcast

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ToTime casts an interface to a time.Time type.
func ToTime(i interface{}) time.Time {
	v, _ := ToTimeE(i)
	return v
}

// ToTimeE casts an interface to a time.Time type.
func ToTimeE(i interface{}) (tim time.Time, err error) {
	return ToTimeInDefaultLocationE(i, time.UTC)
}

func ToTimeInDefaultLocation(i interface{}, location *time.Location) time.Time {
	v, _ := ToTimeInDefaultLocationE(i, location)
	return v
}

// ToTimeInDefaultLocationE casts an empty interface to time.Time,
// interpreting inputs without a timezone to be in the given location,
// or the local timezone if nil.
func ToTimeInDefaultLocationE(i interface{}, location *time.Location) (tim time.Time, err error) {
	i = indirect(i)

	switch v := i.(type) {
	case time.Time:
		return v, nil
	case string:
		return StringToDateInDefaultLocation(v, location)
	case json.Number:
		s, err1 := ToInt64E(v)
		if err1 != nil {
			return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", i, i)
		}
		return time.Unix(s, 0), nil
	case int:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	default:
		return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", i, i)
	}
}

// ToDuration casts an interface to a time.Duration type.
func ToDuration(i interface{}) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// ToDurationE casts an interface to a time.Duration type.
func ToDurationE(i interface{}) (d time.Duration, err error) {
	i = indirect(i)

	switch s := i.(type) {
	case time.Duration:
		return s, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		d = time.Duration(ToInt64(s))
		return
	case float32, float64:
		d = time.Duration(ToFloat64(s))
		return
	case string:
		if strings.ContainsAny(s, "nsuµmh") {
			d, err = time.ParseDuration(s)
		} else {
			d, err = time.ParseDuration(s + "ns")
		}
		return
	case float64EProvider:
		var v float64
		v, err = s.Float64()
		d = time.Duration(v)
		return
	case float64Provider:
		d = time.Duration(s.Float64())
		return
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to Duration", i, i)
		return
	}
}
