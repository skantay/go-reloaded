package tool

import (
	"fmt"
	"regexp"
	"strings"
)

func Cap(text string) string {
	counter := 0
	text = regexp.MustCompile(`([a-zA-Z0-9]+)[\s]*\([\s\n]*([Cc][Aa][Pp])\s*\)|\([\s\n]*([Cc][Aa][Pp])\s*\)`).ReplaceAllStringFunc(text, func(s string) string {
		if counter == 1 {
			return s
		}
		counter++
		word := regexp.MustCompile(`([a-zA-Z0-9]+)[\s]*\([\s\n]*([Cc][Aa][Pp])\s*\)|\([\s\n]*([Cc][Aa][Pp])\s*\)`).ReplaceAllString(s, "$1")
		fmt.Println(s)

		word = strings.ToLower(word)
		return strings.Title(word)
	})

	return text
}

// func CapCount(text string) string {
// 	re := regexp.MustCompile(`(.*?\(cap,\s*\d+\))`)

// 	result := re.ReplaceAllStringFunc(text, func(s string) string {
// 		matches := re.FindStringSubmatch(s)

// 		wordsBuffer := strings.Split(matches[1], " ")
// 		words := wordsBuffer[:len(wordsBuffer)-2]

// 		number, err := CustomAtoi(matches[1])
// 		if err != nil {
// 			return ""
// 		}

// 		if number > len(words) {
// 			number = len(words)
// 		}
// 		for i := 0; i < number; i++ {
// 			words[len(words)-1-i] = strings.Title(words[len(words)-1-i])
// 		}
// 		return strings.Join(words, " ")
// 	})
// 	return result
// }
