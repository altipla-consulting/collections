package collections

// HasString returns true if list contains search.
func HasString(list []string, search string) bool {
	for _, item := range list {
		if item == search {
			return true
		}
	}

	return false
}

// CompareStrings returns true if the values of both lhs & rhs are equals.
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

// DifferenceStrings returns the items present in a and not in b.
func DifferenceStrings(a, b []string) []string {
	result := []string{}

	for _, item := range a {
		if !HasString(b, item) {
			result = append(result, item)
		}
	}

	return result
}

// RemoveString return list without remove.
func RemoveString(list []string, remove string) []string {
	result := []string{}
	for _, item := range list {
		if item != remove {
			result = append(result, item)
		}
	}

	return result
}
