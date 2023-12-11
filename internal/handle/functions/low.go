package tool

import (
	"regexp"
	"strings"
)

func Low(text string) string {
	counter := 0
	text = regexp.MustCompile(`([a-zA-Z0-9]+)[\s]*\([\s\n]*([Ll][Oo][Ww])\s*\)|\([\s\n]*([Ll][Oo][Ww])\s*\)`).ReplaceAllStringFunc(text, func(s string) string {
		if counter == 1 {
			return s
		}
		counter++
		word := regexp.MustCompile(`([a-zA-Z0-9]+)[\s]*\([\s\n]*([Ll][Oo][Ww])\s*\)|\([\s\n]*([Ll][Oo][Ww])\s*\)`).ReplaceAllString(s, "$1")

		return strings.ToLower(word)
	})

	return text
}
