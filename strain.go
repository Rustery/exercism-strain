package strain

type Ints []int
type Lists [][]int
type Strings []string

func keep[T int | []int | string ](collection []T, filter func(T) bool ) []T {
	var result []T
	for _, v := range collection {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func discard[T int | []int | string ](collection []T, filter func(T) bool ) []T {
	var result []T
	for _, v := range collection {
		if !filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func (i Ints) Keep(filter func(int) bool) Ints {
	return keep(i, filter)
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return discard(i, filter)
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	return keep(l, filter)
}

func (s Strings) Keep(filter func(string) bool) Strings {
	return keep(s, filter)
}
