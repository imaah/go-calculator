package regex

import (
	"emorisse.fr/go-calculator/pkg/operation/unary"
	"emorisse.fr/go-calculator/pkg/utils"
	"fmt"
	"regexp"
)

var UnaryExpressionRegex *regexp.Regexp
var FindUnaryExpressionRegex *regexp.Regexp

func init() {
	var unaryRegexStr = fmt.Sprintf(`([%s])(?:(:[0-9]+|-?[0-9]+(?:.[0-9]+)?))`,
		utils.SanitizeOperators(unary.KnownSymbols))

	FindUnaryExpressionRegex = regexp.MustCompile(unaryRegexStr)
	UnaryExpressionRegex = regexp.MustCompile(`^` + unaryRegexStr + `$`)
}

func IsUnaryExpression(expression string) bool {
	return UnaryExpressionRegex.MatchString(expression)
}
