package tool

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Low(text, numFlag string) string {
	lowFirst := `(\n|.)*\([\s\n]*([Ll][Oo][Ww])[\s\n]*,[\s\n]*\d*[\s\n]*\)`
	lowSecond := `([a-zA-Z0-9]+)[\s]*\([\s\n]*([Ll][Oo][Ww])\s*\)|\([\s\n]*([Ll][Oo][Ww])\s*\)`

	var done bool

	regex := lowFirst

	num, err := strconv.Atoi(numFlag)
	if err != nil {
		regex = lowSecond
	}

	text = regexp.MustCompile(regex).ReplaceAllStringFunc(text, func(s string) string {
		if done {
			return s
		}
		done = true

		if regex == lowSecond {
			word := regexp.MustCompile(regex).ReplaceAllString(s, "$1")
			return strings.ToLower(word)
		}
		word := regexp.MustCompile(regex).ReplaceAllString(s, "$0")
		indexWord := strings.LastIndex(word, "(")
		word = helperLow(word, indexWord, num)
		return word
	})

	return text
}

func helperLow(text string, indexWord, num int) string {
	trimmed := strings.TrimRight(string(text[:indexWord]), " ")
	for {
		if text[indexWord] != ' ' && text[indexWord] != '\n' {
			break
		}
		trimmed = strings.TrimRight(string(trimmed), "\n")
		trimmed = strings.TrimRight(string(trimmed), " ")
	}
	result := []rune(trimmed)
	var isWord bool
	var foundLetter bool
	for i := len(result) - 1; i >= 0; i-- {

		if num == 0 {
			break
		}

		if unicode.IsLetter(result[i]) {
			foundLetter = true
			isWord = true
			result[i] = unicode.ToLower(result[i])
			continue
		}
		
		if isWord && foundLetter {
			isWord = false
			num--
		}
	}
	return string(result)
}
