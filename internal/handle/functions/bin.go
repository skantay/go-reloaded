package tool

import (
	"regexp"
	"strconv"
	"strings"
)

func Bin(text string) string {
	sliceOfRegex := []string{
		`([0-1]+)[\s\W]*\([\s\n]*([Bb][Ii][Nn])\s*\)`,
		`\([\s\n]*([Bb][Ii][Nn])\s*\)`,
	}
	regex := strings.Join(sliceOfRegex, "|")

	counter := 0
	text = regexp.MustCompile(regex).ReplaceAllStringFunc(text, func(s string) string {
		if counter == 1 {
			return s
		}
		counter++
		word := regexp.MustCompile(regex).ReplaceAllString(s, "$1")
		dec, err := strconv.ParseInt(word, 2, 64)
		if err == nil {
			return strconv.Itoa(int(dec))
		}

		return word
	})

	return text
}
