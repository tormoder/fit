package profile

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var camelRegex = regexp.MustCompile("[0-9A-Za-z]+")

func toCamelCase(s string) string {
	chunks := camelRegex.FindAllString(s, -1)
	for i, val := range chunks {
		chunks[i] = strings.Title(val)
	}
	return strings.Join(chunks, "")
}

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func trimFloat(val string) string {
	const floatSuffix = ".0"
	return strings.TrimSuffix(val, floatSuffix)
}
