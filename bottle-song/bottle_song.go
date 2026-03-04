// Package bottlesong implements the solution to the "99 Bottles of Beer" exercise.
package bottlesong

import (
	"fmt"
	"strings"
)

// Recite returns the verses of the "99 Bottles of Beer" song, starting from startBottles and taking down takeDown bottles.
func Recite(startBottles, takeDown int) []string {
	result := []string{}
	for i := range takeDown {
		num := startBottles - i
		result = append(result, verse(num)...)
		if i != takeDown-1 {
			result = append(result, "")
		}
	}
	return result
}

func verse(num int) []string {
	return []string{
		fmt.Sprintf("%s green bottle%s hanging on the wall,", numToWord(num), plural(num)),
		fmt.Sprintf("%s green bottle%s hanging on the wall,", numToWord(num), plural(num)),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.", strings.ToLower(numToWord(num-1)), plural(num-1)),
	}
}

func numToWord(num int) string {
	words := []string{"No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	return words[num]
}

func plural(num int) string {
	if num != 1 {
		return "s"
	}
	return ""
}
