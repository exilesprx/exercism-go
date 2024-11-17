package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("Cannot be less than or equal to 0")
	}

	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}

		steps++
	}

	return steps, nil
}
