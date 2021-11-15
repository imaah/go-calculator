package regex

import (
	"emorisse.fr/go-calculator/pkg/operation/binary"
	"emorisse.fr/go-calculator/pkg/utils"
	"fmt"
	"regexp"
)

var BinaryExpressionRegex *regexp.Regexp

func init() {
	var binaryRegexStr = fmt.Sprintf(`^\(?(:[0-9]+|-?[0-9]+(?:.[0-9]+)?) *([%s]) *(:[0-9]+|-?[0-9]+(?:.[0-9]+)?)\)?$`,
		utils.SanitizeOperators(binary.KnownSymbols))
	BinaryExpressionRegex = regexp.MustCompile(binaryRegexStr)
}

func IsBinaryExpression(expression string) bool {
	return BinaryExpressionRegex.MatchString(expression)
}
