package utils

//RuneArrayContains returns whether the needle is inside the array or not
func RuneArrayContains(arr []rune, needle rune) bool {
	for _, elem := range arr {
		if elem == needle {
			return true
		}
	}
	return false
}

//StringArrayContains returns whether the needle is inside the array or not
func StringArrayContains(arr []string, needle string) bool {
	for _, elem := range arr {
		if elem == needle {
			return true
		}
	}
	return false
}
