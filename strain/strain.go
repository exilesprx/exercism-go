package strain

// Keep returns a new slice containing all elements that satisfy the predicate.
func Keep[T any](list []T, predicate func(T) bool) []T {
	nlist := []T{}
	for _, t := range list {
		if predicate(t) {
			nlist = append(nlist, t)
		}
	}

	return nlist
}

// Discard returns a new slice containing all elements that do not satisfy the predicate.
func Discard[T any](list []T, predicate func(T) bool) []T {
	nlist := []T{}
	for _, t := range list {
		if !predicate(t) {
			nlist = append(nlist, t)
		}
	}

	return nlist
}
