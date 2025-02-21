package fjson

import (
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

type CustomTimeExtension struct {
	jsoniter.DummyExtension
}

func (extension *CustomTimeExtension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
	for _, binding := range structDescriptor.Fields {
		var typeErr error
		var isPtr bool
		typeName := binding.Field.Type().String()
		if typeName == "time.Time" {
			isPtr = false
		} else if typeName == "*time.Time" {
			isPtr = true
		} else {
			continue
		}

		timeFormat := defaultFormat
		formatTag := binding.Field.Tag().Get(tagNameTimeFormat)
		if format, ok := formatAlias[formatTag]; ok {
			timeFormat = format
		} else if formatTag != "" {
			timeFormat = formatTag
		}
		locale := defaultLocale
		if localeTag := binding.Field.Tag().Get(tagNameTimeLocation); localeTag != "" {
			if loc, ok := localeAlias[localeTag]; ok {
				locale = loc
			} else {
				loc, err := time.LoadLocation(localeTag)
				if err != nil {
					typeErr = err
				} else {
					locale = loc
				}
			}
		}

		binding.Encoder = &funcEncoder{
			fun: func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
				if typeErr != nil {
					stream.Error = typeErr
					return
				}

				var tp *time.Time
				if isPtr {
					tpp := (**time.Time)(ptr)
					tp = *(tpp)
				} else {
					tp = (*time.Time)(ptr)
				}

				if tp != nil {
					lt := tp.In(locale)
					str := lt.Format(timeFormat)
					stream.WriteString(str)
				} else {
					stream.WriteNil()
				}
			},
			isEmptyFunc: func(ptr unsafe.Pointer) bool {
				return (*time.Time)(ptr).IsZero()
			},
		}

		binding.Decoder = &funcDecoder{fun: func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if typeErr != nil {
				iter.Error = typeErr
				return
			}

			str := iter.ReadString()
			var t *time.Time
			if str != "" {
				var err error
				tmp, err := time.ParseInLocation(timeFormat, str, locale)
				if err != nil {
					iter.Error = err
					return
				}
				t = &tmp
			} else {
				t = nil
			}

			if isPtr {
				tpp := (**time.Time)(ptr)
				*tpp = t
			} else {
				tp := (*time.Time)(ptr)
				if tp != nil && t != nil {
					*tp = *t
				}
			}
		}}
	}
}
