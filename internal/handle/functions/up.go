package tool

import (
	"regexp"
	"strings"
)

func Up(text string) string {
	counter := 0
	text = regexp.MustCompile(`([a-zA-Z0-9]+)[\s]*\([\s\n]*([Uu][Pp])\s*\)|\([\s\n]*([Uu][Pp])\s*\)`).ReplaceAllStringFunc(text, func(s string) string {
		if counter == 1 {
			return s
		}
		counter++
		word := regexp.MustCompile(`([a-zA-Z0-9]+)[\s]*\([\s\n]*([Uu][Pp])\s*\)|\([\s\n]*([Uu][Pp])\s*\)`).ReplaceAllString(s, "$1")
		return strings.ToUpper(word)
	})

	return text
}
