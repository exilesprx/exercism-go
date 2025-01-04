package romannumerals

import (
	"errors"
)

type RomanNumeral struct {
	letters string
	value   int
}

var ruleSet = []RomanNumeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("invalid number")
	}

	romanNumerals := ""
	for input > 0 {
		for _, r := range ruleSet {
			if input >= r.value {
				input = input - r.value
				romanNumerals += r.letters
				break
			}
		}
	}

	return romanNumerals, nil
}
