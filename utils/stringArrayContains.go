package utils

// StringArrayContains checks whether or not the given string array contains the given string and returns the index if it does
func StringArrayContains(array []string, element string) (int, bool) {
	for index, str := range array {
		if str == element {
			return index, true
		}
	}
	return -1, false
}
