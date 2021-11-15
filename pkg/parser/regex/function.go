package regex

import (
	"fmt"
	"regexp"
)

var FunctionExpressionRegex *regexp.Regexp
var FindFunctionExpressionsRegex *regexp.Regexp

func init() {
	FunctionExpressionRegex = regexp.MustCompile(`^ *([a-z][a-z0-9]*) *\((:[0-9]+|-?[0-9]+(?:[0-9]+)?)\)$`)
	FindFunctionExpressionsRegex = regexp.MustCompile(`([a-z][a-z0-9]*) *\(-?[0-9]+(?:[0-9]+)?\)`)
}

func IsFunctionExpression(expression string) bool {
	fmt.Println(expression)
	return FunctionExpressionRegex.MatchString(expression)
}
