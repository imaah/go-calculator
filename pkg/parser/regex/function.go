package regex

import (
	"regexp"
)

var FunctionRegex = regexp.MustCompile("^ *[a-zA-Z][0-9a-zA-Z]+{[0-9]+} *$")

func IsFunctionExpression(expression string) bool {
	return FunctionRegex.MatchString(expression)
}
