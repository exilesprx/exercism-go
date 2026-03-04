// Package bottlesong implements the solution to the "99 Bottles of Beer" exercise.
package bottlesong

import (
	"fmt"
	"strings"
)

// Recite returns the verses of the "99 Bottles of Beer" song, starting from startBottles and taking down takeDown bottles.
func Recite(startBottles, takeDown int) []string {
	phrases := []string{
		"%s green %s hanging on the wall,",
		"%s green %s hanging on the wall,",
		"And if one green bottle should accidentally fall,",
		"There'll be %s green %s hanging on the wall.",
	}
	result := []string{}
	for i := range takeDown {
		num := startBottles - i
		result = append(result,
			fmt.Sprintf(phrases[0], numToWord(num), numToBottle(num)),
			fmt.Sprintf(phrases[1], numToWord(num), numToBottle(num)),
			phrases[2],
			fmt.Sprintf(phrases[3], strings.ToLower(numToWord(num-1)), numToBottle(num-1)),
		)
		if i != takeDown-1 {
			result = append(result, "")
		}
	}
	return result
}

func numToWord(num int) string {
	switch num {
	case 0:
		return "No"
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	default:
		return ""
	}
}

func numToBottle(num int) string {
	if num == 1 {
		return "bottle"
	}
	return "bottles"
}
