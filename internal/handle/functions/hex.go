package tool

import (
	"fmt"
	"regexp"
	"strconv"
)

func Hex(text string) string {
	counter := 0
	text = regexp.MustCompile(`([a-zA-Z0-9]+)[\s\W]*\(([Hh][Ee][Xx])\)`).ReplaceAllStringFunc(text, func(s string) string {
		if counter == 1 {
			return s
		}
		counter++
		word := regexp.MustCompile(`([a-zA-Z0-9]+)[\s\W]*\(([Hh][Ee][Xx])\)`).ReplaceAllString(s, "$1")
		dec, err := strconv.ParseInt(word, 16, 64)
		if err == nil {
			return fmt.Sprintf("%d", dec)
		}

		return word
	})

	return text
}
