// Package atbashcipher implements the Atbash cipher, a simple substitution cipher that replaces each letter with its reverse counterpart in the alphabet.
package atbashcipher

import (
	"strings"
	"unicode"
)

// Atbash takes a string and returns its Atbash cipher equivalent. The function should ignore capitalization and non-alphabetic characters, and the output should be in lowercase.
func Atbash(s string) string {
	str := strings.Builder{}
	spaces := 0
	for _, r := range s {
		if unicode.IsPunct(r) || unicode.IsSpace(r) {
			continue
		}
		if str.Len() != 0 && (str.Len()-spaces)%5 == 0 {
			str.WriteRune(' ')
			spaces++
		}
		if unicode.IsLetter(r) {
			str.WriteRune('z' - (unicode.ToLower(r) - 'a'))
			continue
		}
		str.WriteRune(r)
	}
	return str.String()
}
