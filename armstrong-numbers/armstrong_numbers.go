// Package armstrongnumbers provides a function to determine if a number is an Armstrong number.
package armstrongnumbers

import (
	"math"
	"strconv"
)

// IsNumber checks if the given integer n is an Armstrong number.
func IsNumber(n int) bool {
	str := strconv.Itoa(n)
	digits := len(str)

	sum := 0
	for _, r := range str {
		sum += int(math.Pow(float64(int(r-'0')), float64(digits)))
	}
	return sum == n
}
