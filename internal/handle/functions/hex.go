package tool

import (
	"regexp"
	"strconv"
	"strings"
)

func Hex(text string) string {
	sliceOfRegex := []string{
		`([a-fA-F0-9]+)[\s\W]*\([\s\n]*([Hh][Ee][Xx])\s*\)`,
		`\([\s\n]*([Hh][Ee][Xx])\s*\)`,
	}
	regex := strings.Join(sliceOfRegex, "|")
	counter := 0
	text = regexp.MustCompile(regex).ReplaceAllStringFunc(text, func(s string) string {
		if counter == 1 {
			return s
		}
		counter++
		word := regexp.MustCompile(regex).ReplaceAllString(s, "$1")
		dec, err := strconv.ParseInt(word, 16, 64)
		if err == nil {
			return strconv.Itoa(int(dec))
		}

		return word
	})

	return text
}
