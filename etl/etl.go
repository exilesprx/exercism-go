package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	mapping := map[string]int{}
	for points, letters := range in {
		for _, letter := range letters {
			mapping[strings.ToLower(letter)] = points
		}
	}

	return mapping
}
