package collections

// HasString returns if list contains search.
func HasString(list []string, search string) bool {
	for _, item := range list {
		if item == search {
			return true
		}
	}

	return false
}

// CompareStrings returns if the values of two list are equals.
func CompareStrings(lhs, rhs []string) bool {
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

// UniqueStrings returns list without any duplicates in it.
func UniqueStrings(list []string) []string {
	present := map[string]bool{}
	var result []string
	for _, item := range list {
		if !present[item] {
			present[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Difference returns the items present in a and not in b.
func Difference(a, b []string) []string {
	result := []string{}

	for _, item := range a {
		if !HasString(b, item) {
			result = append(result, item)
		}
	}

	return result
}

// RemoveString return the same list without remove.
func RemoveString(list []string, remove string) []string {
	result := []string{}
	for _, item := range list {
		if item != remove {
			result = append(result, item)
		}
	}

	return result
}

// HasInt64 returns if list of Integer64 contains search.
func HasInt64(list []int64, search int64) bool {
	for _, item := range list {
		if item == search {
			return true
		}
	}

	return false
}

// UniqueIntegers64 returns list of Integer64 without any duplicates in it.
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

// HasInt32 returns if list of Integer32 contains search
func HasInt32(list []int32, search int32) bool {
	for _, item := range list {
		if item == search {
			return true
		}
	}

	return false
}

// CompareInts returns if the values of two list are equals.
func CompareInts(lhs, rhs []int32) bool {
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
