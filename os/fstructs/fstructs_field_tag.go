package fstructs

import (
	"strings"

	"github.com/ZYallers/fine/util/ftag"
)

// TagJsonName returns the `json` tag name string of the field.
func (f *Field) TagJsonName() string {
	if jsonTag := f.Tag(ftag.Json); jsonTag != "" {
		return strings.Split(jsonTag, ",")[0]
	}
	return ""
}

// TagDefault returns the most commonly used tag `default/d` value of the field.
func (f *Field) TagDefault() string {
	v := f.Tag(ftag.Default)
	if v == "" {
		v = f.Tag(ftag.DefaultShort)
	}
	return v
}

// TagParam returns the most commonly used tag `param/p` value of the field.
func (f *Field) TagParam() string {
	v := f.Tag(ftag.Param)
	if v == "" {
		v = f.Tag(ftag.ParamShort)
	}
	return v
}

// TagValid returns the most commonly used tag `valid/v` value of the field.
func (f *Field) TagValid() string {
	v := f.Tag(ftag.Valid)
	if v == "" {
		v = f.Tag(ftag.ValidShort)
	}
	return v
}

// TagDescription returns the most commonly used tag `description/des/dc` value of the field.
func (f *Field) TagDescription() string {
	v := f.Tag(ftag.Description)
	if v == "" {
		v = f.Tag(ftag.DescriptionShort)
	}
	if v == "" {
		v = f.Tag(ftag.DescriptionShort2)
	}
	return v
}

// TagSummary returns the most commonly used tag `summary/sum/sm` value of the field.
func (f *Field) TagSummary() string {
	v := f.Tag(ftag.Summary)
	if v == "" {
		v = f.Tag(ftag.SummaryShort)
	}
	if v == "" {
		v = f.Tag(ftag.SummaryShort2)
	}
	return v
}

// TagAdditional returns the most commonly used tag `additional/ad` value of the field.
func (f *Field) TagAdditional() string {
	v := f.Tag(ftag.Additional)
	if v == "" {
		v = f.Tag(ftag.AdditionalShort)
	}
	return v
}

// TagExample returns the most commonly used tag `example/eg` value of the field.
func (f *Field) TagExample() string {
	v := f.Tag(ftag.Example)
	if v == "" {
		v = f.Tag(ftag.ExampleShort)
	}
	return v
}

// TagIn returns the most commonly used tag `in` value of the field.
func (f *Field) TagIn() string {
	v := f.Tag(ftag.In)
	return v
}

// TagPriorityName checks and returns tag name that matches the name item in `ftag.StructTagPriority`.
// It or else returns attribute field Name if it doesn't have a tag name by `ftag.StructsTagPriority`.
func (f *Field) TagPriorityName() string {
	name := f.Name()
	for _, tagName := range ftag.StructTagPriority {
		if tagValue := f.Tag(tagName); tagValue != "" {
			name = tagValue
			break
		}
	}
	return name
}
