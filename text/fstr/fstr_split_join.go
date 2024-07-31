package fstr

import (
	"strings"
)

// Join concatenates the elements of `array` to create a single string. The separator string
// `sep` is placed between elements in the resulting string.
func Join(array []string, sep string) string {
	return strings.Join(array, sep)
}

// Split splits string `str` by a string `delimiter`, to an array.
func Split(str, delimiter string) []string {
	return strings.Split(str, delimiter)
}

// SplitAndTrim splits string `str` by a string `delimiter` to an array,
// and calls Trim to every element of this array. It ignores the elements
// which are empty after Trim.
func SplitAndTrim(str, delimiter string, characterMask ...string) []string {
	array := make([]string, 0)
	for _, v := range strings.Split(str, delimiter) {
		v = Trim(v, characterMask...)
		if v != "" {
			array = append(array, v)
		}
	}
	return array
}
