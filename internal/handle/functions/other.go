package tool

import (
	"fmt"
	"regexp"
	"strings"
)

func Articles(s string) string {
	regex := `\b[Aa]\b\s*[aAeEhHuUoOiI]`
	re := regexp.MustCompile(regex)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		last := match[len(match)-2:]
		modify := match[:len(match)-2]
		modify = strings.ReplaceAll(modify, "a", "an")
		modify = strings.ReplaceAll(modify, "A", "An")
		modify += string(last)
		return modify
	})
}

func Quotes(text string) string {
	regex := `'\s*(.*?)\s*'`
	return regexp.MustCompile(regex).ReplaceAllStringFunc(text, func(match string) string {
		var from int = len(match)

		var to int = 0

		for i, v := range match {
			if v != ' ' && v != '\'' && v != '\n' {
				to = i
			}
		}

		for i := len(match) - 1; i >= 0; i-- {
			v := match[i]
			if v != ' ' && v != '\'' && v != '\n' {
				from = i
			}
		}
		if from >= to {
			return strings.ReplaceAll(match, " ", "")
		}
		word := match[from:to+1]
		result := fmt.Sprintf("'%s'", word)

		return result
	})
}

func Punct(text string) string {
	sliceOfRegex := []string{
		`[a-zA-Z'](\s+)(\.{3})`,
		`[a-zA-Z'](\s\n+)(\!{2})`,
		`[a-zA-Z'](\s+)([.,!?:;])`,
	}
	regex := strings.Join(sliceOfRegex, "|")

	re := regexp.MustCompile(regex)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		return strings.ReplaceAll(match, " ", "")
	})
}
