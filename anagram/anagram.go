package anagram

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	for _, candidate := range candidates {
		if strings.EqualFold(candidate, subject) || utf8.RuneCountInString(candidate) != utf8.RuneCountInString(subject) {
			continue
		}

		letters := map[rune]int{}
		matches := map[rune]int{}
		for _, letter := range subject {
			letters[unicode.ToLower(letter)]++
		}

		isAnagram := true
		for _, letter := range candidate {
			if value, ok := letters[unicode.ToLower(letter)]; !ok || value == matches[unicode.ToLower(letter)] {
				isAnagram = false
				break
			}

			matches[unicode.ToLower(letter)]++
		}

		if isAnagram {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
