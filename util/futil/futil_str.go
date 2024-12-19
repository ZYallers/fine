package futil

import (
	"github.com/ZYallers/fine/internal/util/utils"
)

// IsLetterUpper checks whether the given byte b is in upper case.
func IsLetterUpper(b byte) bool {
	return utils.IsLetterUpper(b)
}

// IsLetterLower checks whether the given byte b is in lower case.
func IsLetterLower(b byte) bool {
	return utils.IsLetterLower(b)
}

// IsLetter checks whether the given byte b is a letter.
func IsLetter(b byte) bool {
	return utils.IsLetterUpper(b) || utils.IsLetterLower(b)
}

// IsNumeric checks whether the given string s is numeric.
// Note that float string like "123.456" is also numeric.
func IsNumeric(s string) bool {
	return utils.IsNumeric(s)
}

// UcFirst returns a copy of the string s with the first letter mapped to its upper case.
func UcFirst(s string) string {
	return utils.UcFirst(s)
}

// ReplaceByMap returns a copy of `origin`,
// which is replaced by a map in unordered way, case-sensitively.
func ReplaceByMap(origin string, replaces map[string]string) string {
	return utils.ReplaceByMap(origin, replaces)
}

// RemoveSymbols removes all symbols from string and lefts only numbers and letters.
func RemoveSymbols(s string) string {
	return utils.RemoveSymbols(s)
}

// EqualFoldWithoutChars checks string `s1` and `s2` equal case-insensitively,
// with/without chars '-'/'_'/'.'/' '.
func EqualFoldWithoutChars(s1, s2 string) bool {
	return utils.EqualFoldWithoutChars(s1, s2)
}

// SplitAndTrim splits string `str` by a string `delimiter` to an array,
// and calls Trim to every element of this array. It ignores the elements
// which are empty after Trim.
func SplitAndTrim(str, delimiter string, characterMask ...string) []string {
	return utils.SplitAndTrim(str, delimiter, characterMask...)
}

// Trim strips whitespace (or other characters) from the beginning and end of a string.
// The optional parameter `characterMask` specifies the additional stripped characters.
func Trim(str string, characterMask ...string) string {
	return utils.Trim(str, characterMask...)
}

// FormatCmdKey formats string `s` as command key using uniformed format.
func FormatCmdKey(s string) string {
	return utils.FormatCmdKey(s)
}

// FormatEnvKey formats string `s` as environment key using uniformed format.
func FormatEnvKey(s string) string {
	return utils.FormatEnvKey(s)
}

// StripSlashes un-quotes a quoted string by AddSlashes.
func StripSlashes(str string) string {
	return utils.StripSlashes(str)
}
