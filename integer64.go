package collections

// HasInt64 returns if list of Integer64 contains search.
func HasInt64(list []int64, search int64) bool {
	for _, item := range list {
		if item == search {
			return true
		}
	}

	return false
}

// UniqueIntegers64 returns a list of Integer64 without any duplicates in it.
func UniqueIntegers64(list []int64) []int64 {
	index := map[int64]bool{}
	result := []int64{}
	for _, item := range list {
		if !index[item] {
			index[item] = true
			result = append(result, item)
		}
	}

	return result
}

// CompareInts64 returns true if the values of both list are equals.
func CompareInts64(lhs, rhs []int32) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for index := range lhs {
		if lhs[index] != rhs[index] {
			return false
		}
	}
	return true
}
