package isbn

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if chars := utf8.RuneCountInString(isbn); chars != 10 {
		return false
	}

	sum := 0
	for i, r := range isbn {
		if i == 9 && r == 'X' {
			sum += 10
			continue
		}

		if unicode.IsLetter(r) {
			return false
		}

		sum += int(r-'0') * (10 - i)
	}

	return sum%11 == 0
}
