// Package lineup provides a function to format a message for customers in a line.
package lineup

import (
	"fmt"
)

// Format takes a customer's name and their position in line, and returns a formatted message.
func Format(name string, number int) string {
	return fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, suffix(number))
}

// suffix returns the appropriate ordinal suffix for a given number.
func suffix(number int) string {
	if number%100 >= 11 && number%100 <= 13 {
		return "th"
	}
	if number%10 == 1 {
		return "st"
	}
	if number%10 == 2 {
		return "nd"
	}
	if number%10 == 3 {
		return "rd"
	}
	return "th"
}
