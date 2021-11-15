package utils

import "strings"

func SanitizeOperators(operators []rune) string {
	var strBuilder strings.Builder

	for _, ope := range operators {
		strBuilder.WriteString("\\")
		strBuilder.WriteRune(ope)
	}

	return strBuilder.String()
}
