package tool

import (
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
		modify += last

		return modify
	})
}

func Quotes(text string) string {
	regex := `'\s*(.*?)\s*'`

	return regexp.MustCompile(regex).ReplaceAllStringFunc(text, func(match string) string {
		trimmed := strings.TrimRight(match, " \n'")
		trimmed = strings.TrimLeft(trimmed, " \n'")

		result := make([]rune, 1, len(trimmed)+5)

		runes := []rune(trimmed)
		result = append(result, runes...)
		result[0] = '\''
		result = append(result, '\'')

		return string(result)
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
