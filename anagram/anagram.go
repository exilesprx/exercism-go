package anagram

import (
	"sort"
	"strings"
	"unicode/utf8"
)

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	for _, candidate := range candidates {
		if strings.EqualFold(candidate, subject) || utf8.RuneCountInString(candidate) != utf8.RuneCountInString(subject) {
			continue
		}

		if sortString(strings.ToLower(candidate)) == sortString(strings.ToLower(subject)) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}
