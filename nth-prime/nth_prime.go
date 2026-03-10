// Package nthprime implements the Nth function, which returns the nth prime number.
package nthprime

import (
	"fmt"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("cannot calculate the %dth prime number", n)
	}

	prime := 0
	for i, count := 2, 0; count < n; i++ {
		if isPrime(i) {
			prime = i
			count++
		}
	}

	return prime, nil
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i <= int(math.Sqrt(float64(n))); i++ {
		// only check odds
		if i%2 == 0 {
			continue
		}
		// if current odd or the next odd step
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 3 // next odd step
	}

	return true
}
