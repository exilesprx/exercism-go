package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	alphabet := map[rune]bool{}
	for _, r := range strings.ToUpper(input) {
		if unicode.IsLetter(r) {
			alphabet[r] = true
		}
	}

	return len(alphabet) == 26
}
