package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	f := func(c rune) bool {
		return c != '\'' && (unicode.IsSpace(c) || unicode.IsPunct(c) || unicode.IsSymbol(c))
	}

	frequency := Frequency{}
	words := strings.FieldsFunc(phrase, f)
	for _, word := range words {
		word = strings.Trim(word, "'\"")
		if word == "" {
			continue
		}
		frequency[strings.ToLower(word)]++
	}
	return frequency
}
