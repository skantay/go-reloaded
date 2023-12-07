package tool

import (
	"regexp"
	"strings"
)

func Up(text string) string {
	re := regexp.MustCompile(`(\b[a-zA-Z0-9_!]+,?)\s*\(up\)`)

	result := re.ReplaceAllStringFunc(text, func(s string) string {
		matches := re.FindStringSubmatch(s)
		text := strings.ToUpper(matches[1])
		return text
	})

	return result
}
