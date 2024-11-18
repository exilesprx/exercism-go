package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	letters := map[rune]bool{}
	for _, letter := range word {
		letter = unicode.ToLower(letter)
		if !unicode.IsLetter(letter) {
			continue
		}
		if letters[letter] {
			return false
		}
		letters[letter] = true
	}

	return true
}
