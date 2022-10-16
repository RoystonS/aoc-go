package aoccommon

// Converts a map from key to value into a map from value to keys
func PivotMap[K comparable, V comparable](counts map[K]V) map[V][]K {
	result := map[V][]K{}

	for r, count := range counts {
		arr := result[count]
		arr = append(arr, r)
		result[count] = arr
	}

	return result
}
