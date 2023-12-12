package tool

import (
	"fmt"
	"regexp"
	"strconv"
)

func Hex(text string) string {
	counter := 0
	text = regexp.MustCompile(`([a-fA-F0-9]+)[\s\W]*\([\s\n]*([Hh][Ee][Xx])\s*\)|\([\s\n]*([Hh][Ee][Xx])\s*\)`).ReplaceAllStringFunc(text, func(s string) string {
		if counter == 1 {
			return s
		}
		counter++
		word := regexp.MustCompile(`([a-fA-F0-9]+)[\s\W]*\([\s\n]*([Hh][Ee][Xx])\s*\)|\([\s\n]*([Hh][Ee][Xx])\s*\)`).ReplaceAllString(s, "$1")
		dec, err := strconv.ParseInt(word, 16, 64)
		if err == nil {
			return fmt.Sprintf("%d", dec)
		}

		return word
	})

	return text
}
