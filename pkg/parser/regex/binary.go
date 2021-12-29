package regex

import (
	"regexp"
)

var BinaryRegex = regexp.MustCompile("^ *{[0-9]+} *[+*/-] *{[0-9]+} *$")

func IsBinaryExpression(expression string) bool {
	return BinaryRegex.MatchString(expression)
}
