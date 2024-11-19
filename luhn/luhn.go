package luhn

import (
	"strconv"
	"unicode/utf8"
)

func Valid(id string) bool {
	sum := 0
	digitCount := 0
	for i := utf8.RuneCountInString(id) - 1; i >= 0; i-- {
		if id[i] == ' ' {
			continue
		}

		num, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}

		if digitCount++; digitCount%2 != 0 {
			sum += num
			continue
		}

		doubled := 2 * num
		if doubled > 9 {
			sum += doubled - 9
      continue
		}

		sum += doubled
	}

	return digitCount > 1 && sum%10 == 0
}
