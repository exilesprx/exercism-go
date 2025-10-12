package lsproduct

import (
	"fmt"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, fmt.Errorf("span must not be negative")
	}

	if span > len(digits) {
		return 0, fmt.Errorf("span must be smaller than string length")
	}

	var lsp int64 = 0
	for i := 0; i < len(digits); i++ {
		if i+span > len(digits) {
			break
		}

		var product int64 = 0
		for j, r := range digits[i : i+span] {
			if !unicode.IsDigit(r) {
				return 0, fmt.Errorf("digits input must only contain digits")
			}

			value := int64(r - '0')
			if j == 0 {
				product = value
				continue
			}
			product = value * product
		}

		if product > lsp {
			lsp = product
		}
	}

	return lsp, nil
}
