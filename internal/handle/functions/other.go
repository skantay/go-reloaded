package tool

import (
	"fmt"
	"regexp"
	"strings"
)

func Articles(text string) string {
	counter := 0
	text = regexp.MustCompile(`[Aa]\s*[aehuoi]`).ReplaceAllStringFunc(text, func(s string) string {
		if counter == 1 {
			return s
		}
		counter++
		word := regexp.MustCompile(`[Aa]\s*[aehuoi]`).ReplaceAllString(s, "$0")
		fmt.Println(word)
		return strings.ToUpper(word)
	})

	return text
}

func Quotes(text string) string {
	result := []rune(text)
	var first int
	var second int

	for i := 0; i < len(result); i++ {
		if result[i] == '\'' && i != len(result)-1 {
			if result[i+1] != ' ' {
				continue
			}
			first = i + 1

			for j := i + 1; j < len(result); j++ {
				if result[j] == '\'' {
					second = j - 1
					i = j
					break
				}
			}
			result = quotHelper(result, first, second)

		}
	}
	return string(result)
}

func quotHelper(result []rune, first, second int) []rune {
	res := result
	for {
		if res[first] == ' ' {
			res[first] = '\''
			res[first-1] = ' '
			first++
		} else {
			break
		}
	}
	for {
		if res[second] == ' ' {
			res[second] = '\''
			res[second+1] = ' '
			second--
		} else {
			break
		}
	}
	return res
}

func Punct(text string) string {
	result := []rune(text)
	for i := 0; i < len(result); i++ {
		if i == len(result)-1 {
			switch result[i] {
			case '.':
				helperPunct(result, i)
			case ',':
				helperPunct(result, i)
			case '!':
				helperPunct(result, i)
			case '?':
				helperPunct(result, i)
			case ':':
				helperPunct(result, i)
			case ';':
				helperPunct(result, i)
			}
			break
		}

		switch result[i] {
		case '.':
			helperPunct(result, i)
			if result[i+1] == '.' && result[i+2] == '.' {
				helperPunct(result, i)
				helperPunct(result, i)
			}
		case ',':
			helperPunct(result, i)
		case '!':
			helperPunct(result, i)

			if result[i+1] == '!' {
				helperPunct(result, i)
			}
		case '?':
			helperPunct(result, i)
		case ':':
			helperPunct(result, i)
		case ';':
			helperPunct(result, i)
		}
	}
	return string(result)
}

func helperPunct(result []rune, i int) []rune {
	for {
		if result[i-1] == ' ' {
			result[i-1] = result[i]
			result[i] = ' '
			i--
			continue
		}
		return result
	}
}
