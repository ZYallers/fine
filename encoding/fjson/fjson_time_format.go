package fjson

import (
	"time"
)

// time format alias
const (
	ANSIC       = "ANSIC"
	UnixDate    = "UnixDate"
	RubyDate    = "RubyDate"
	RFC822      = "RFC822"
	RFC822Z     = "RFC822Z"
	RFC850      = "RFC850"
	RFC1123     = "RFC1123"
	RFC1123Z    = "RFC1123Z"
	RFC3339     = "RFC3339"
	RFC3339Nano = "RFC3339Nano"
	Kitchen     = "Kitchen"
	Stamp       = "Stamp"
	StampMilli  = "StampMilli"
	StampMicro  = "StampMicro"
	StampNano   = "StampNano"
	ShangHai    = "ShangHai"
)

// time zone alias
const (
	Local = "Local"
	UTC   = "UTC"
)

// time tag name
const (
	tagNameTimeFormat   = "time_format"
	tagNameTimeLocation = "time_location"
)

var formatAlias = map[string]string{
	ANSIC:       time.ANSIC,
	UnixDate:    time.UnixDate,
	RubyDate:    time.RubyDate,
	RFC822:      time.RFC822,
	RFC822Z:     time.RFC822Z,
	RFC850:      time.RFC850,
	RFC1123:     time.RFC1123,
	RFC1123Z:    time.RFC1123Z,
	RFC3339:     time.RFC3339,
	RFC3339Nano: time.RFC3339Nano,
	Kitchen:     time.Kitchen,
	Stamp:       time.Stamp,
	StampMilli:  time.StampMilli,
	StampMicro:  time.StampMicro,
	StampNano:   time.StampNano,
	ShangHai:    "2006-01-02 15:04:05",
}

var (
	localeAlias   = map[string]*time.Location{Local: time.Local, UTC: time.UTC}
	defaultFormat = formatAlias[ShangHai]
	defaultLocale = time.Local
)

func AddTimeFormatAlias(alias, format string) {
	formatAlias[alias] = format
}

func AddLocaleAlias(alias string, locale *time.Location) {
	localeAlias[alias] = locale
}

func SetDefaultTimeFormat(timeFormat string, timeLocation *time.Location) {
	defaultFormat = timeFormat
	defaultLocale = timeLocation
}
