// Package listops implements a number of list operations on lists of integers without using built-in functions
package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := s.Length() - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	lst := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			lst = lst.Append(IntList{v})
		}
	}
	return lst
}

func (s IntList) Length() int {
	i := 0
	for range s {
		i++
	}
	return i
}

func (s IntList) Map(fn func(int) int) IntList {
	lst := make(IntList, s.Length())
	for i, v := range s {
		lst[i] = fn(v)
	}
	return lst
}

func (s IntList) Reverse() IntList {
	lstLen := s.Length()
	lst := make(IntList, lstLen)
	for i, v := range s {
		lst[lstLen-1-i] = v
	}
	return lst
}

func (s IntList) Append(lst IntList) IntList {
	l := make(IntList, s.Length()+lst.Length())
	i := 0
	for _, v := range s {
		l[i] = v
		i++
	}
	for _, v := range lst {
		l[i] = v
		i++
	}
	return l
}

func (s IntList) Concat(lists []IntList) IntList {
	lst := s.Append(IntList{})
	for _, v := range lists {
		lst = lst.Append(v)
	}
	return lst
}
