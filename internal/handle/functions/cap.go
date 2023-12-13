package tool

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Cap(text, numFlag string) string {
	capFirst := `(\n|.)*\([\s\n]*([Cc][Aa][Pp])[\s\n]*,[\s\n]*\d*[\s\n]*\)`
	capSecond := `([a-zA-Z0-9]+)[\s]*\([\s\n]*([Cc][Aa][Pp])\s*\)|\([\s\n]*([Cc][Aa][Pp])\s*\)`

	var done bool

	regex := capFirst

	num, err := strconv.Atoi(numFlag)
	if err != nil {
		regex = capSecond
	}

	text = regexp.MustCompile(regex).ReplaceAllStringFunc(text, func(match string) string {
		if done {
			return match
		}
		done = true

		if regex == capSecond {
			word := regexp.MustCompile(regex).ReplaceAllString(match, "$1")
			word = strings.ToLower(word)

			return strings.ToUpper(word)
		}

		word := regexp.MustCompile(regex).ReplaceAllString(match, "$0")
		indexWord := strings.LastIndex(word, "(")
		word = helperCap(word, indexWord, num)

		return word
	})

	return text
}

func helperCap(text string, indexWord, num int) string {
	trimmed := strings.TrimRight(text[:indexWord], " \n")
	result := []rune(trimmed)

	var isWord, foundLetter bool

	for i := len(result) - 1; i >= 0; i-- {
		if num == 0 {
			break
		}

		if unicode.IsLetter(result[i]) {
			foundLetter = true
			isWord = true
			result[i] = unicode.ToLower(result[i])

			continue
		} else if unicode.IsLetter(result[i+1]) {
			result[i+1] = unicode.ToUpper(result[i+1])
		}

		if isWord && foundLetter {
			isWord = false
			num--
		}
	}

	if num > len(strings.Split(text, " ")) {
		result[0] = unicode.ToUpper(result[0])
	}

	return string(result)
}
