package tool

import (
	"fmt"
	"regexp"
	"strconv"
)

func Bin(text string) string {
	counter := 0
	text = regexp.MustCompile(`([0-1]+)[\s\W]*\([\s\n]*([Bb][Ii][Nn])\s*\)|\([\s\n]*([Bb][Ii][Nn])\s*\)`).ReplaceAllStringFunc(text, func(s string) string {
		if counter == 1 {
			return s
		}
		counter++
		word := regexp.MustCompile(`([0-1]+)[\s\W]*\([\s\n]*([Bb][Ii][Nn])\s*\)|\([\s\n]*([Bb][Ii][Nn])\s*\)`).ReplaceAllString(s, "$1")
		dec, err := strconv.ParseInt(word, 2, 64)
		if err == nil {
			return fmt.Sprintf("%d", dec)
		}
		return word
	})

	return text
}
