package profile

import (
	"regexp"
	"strings"
)

var camelRegex = regexp.MustCompile("[0-9A-Za-z]+")

func toCamelCase(s string) string {
	chunks := camelRegex.FindAllString(s, -1)
	for i, val := range chunks {
		chunks[i] = strings.Title(val)
	}
	return strings.Join(chunks, "")
}
