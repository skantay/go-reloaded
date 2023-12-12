package tool

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Up(text, numFlag string) string {
	upFirst := `(\n|.)*\([\s\n]*([Uu][Pp])[\s\n]*,[\s\n]*\d*[\s\n]*\)`
	upSecond  := `([a-zA-Z0-9]+)[\s]*\([\s\n]*([Uu][Pp])\s*\)|\([\s\n]*([Uu][Pp])\s*\)`
	
	var done bool

	regex := upFirst

	num, err := strconv.Atoi(numFlag)
	if err != nil {
		regex = upSecond
	}

	text = regexp.MustCompile(regex).ReplaceAllStringFunc(text, func(s string) string {
		if done {
			return s
		}
		done = true

		if regex == upSecond {
			word := regexp.MustCompile(regex).ReplaceAllString(s, "$1")
			return strings.ToUpper(word)
		}

		word := regexp.MustCompile(regex).ReplaceAllString(s, "$0")
		indexWord := strings.LastIndex(word, "(")
		word = helperUp(word, indexWord, num)
		return word
	})
	return text
}

func helperUp(text string, indexWord, num int) string {
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
			result[i] = unicode.ToUpper(result[i])
			continue
		}

		if isWord && foundLetter {
			isWord = false
			num--
		}
	}
	return string(result)
}
