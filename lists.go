package collections

func HasString(list []string, search string) bool {
	for _, item := range list {
		if item == search {
			return true
		}
	}

	return false
}

func UniqueString(list []string) []string {
	result := []string{}
	seen := map[string]bool{}
	for _, item := range list {
		if seen[item] {
			continue
		}
		seen[item] = true

		result = append(result, item)
	}

	return result
}
