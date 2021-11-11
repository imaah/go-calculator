package utils

func RuneArrayContains(arr []rune, needle rune) bool {
	for _, elem := range arr {
		if elem == needle {
			return true
		}
	}
	return false
}
