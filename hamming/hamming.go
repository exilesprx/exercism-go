package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("String must be of equal length")
	}

	distance := 0
	for i, r := range a {
		if r != rune(b[i]) {
			distance += 1
		}
	}
	return distance, nil
}
