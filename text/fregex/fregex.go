package fregex

import (
	"fmt"
	"regexp"
)

// MatchString return strings that matched `pattern`.
func MatchString(pattern string, src string) ([]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindStringSubmatch(src), nil
	} else {
		return nil, err
	}
}

// Replace replaces all matched `pattern` in bytes `src` with bytes `replace`.
func Replace(pattern string, replace, src []byte) ([]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.ReplaceAll(src, replace), nil
	} else {
		return nil, err
	}
}

// ReplaceString replace all matched `pattern` in string `src` with string `replace`.
func ReplaceString(pattern, replace, src string) (string, error) {
	r, e := Replace(pattern, []byte(replace), []byte(src))
	return string(r), e
}

// MatchAllString return all strings that matched `pattern`.
func MatchAllString(pattern string, src string) ([][]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindAllStringSubmatch(src, -1), nil
	} else {
		return nil, err
	}
}

// IsMatchString checks whether given string `src` matches `pattern`.
func IsMatchString(pattern string, src string) bool {
	return IsMatch(pattern, []byte(src))
}

// IsMatch checks whether given bytes `src` matches `pattern`.
func IsMatch(pattern string, src []byte) bool {
	if r, err := getRegexp(pattern); err == nil {
		return r.Match(src)
	}
	return false
}

// ReplaceStringFuncMatch replace all matched `pattern` in string `src`
// with custom replacement function `replaceFunc`.
// The parameter `match` type for `replaceFunc` is []string,
// which is the result contains all sub-patterns of `pattern` using MatchString function.
func ReplaceStringFuncMatch(pattern string, src string, replaceFunc func(match []string) string) (string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return string(r.ReplaceAllFunc([]byte(src), func(bytes []byte) []byte {
			match, _ := MatchString(pattern, string(bytes))
			return []byte(replaceFunc(match))
		})), nil
	} else {
		return "", err
	}
}

func getRegexp(pattern string) (regex *regexp.Regexp, err error) {
	if regex, err = regexp.Compile(pattern); err != nil {
		err = fmt.Errorf(`regexp.Compile failed for pattern "%s": %s`, pattern, err)
		return
	}
	return
}
