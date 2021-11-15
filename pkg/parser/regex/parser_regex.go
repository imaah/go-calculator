package regex

import (
	"emorisse.fr/go-calculator/pkg/operation/binary"
	"emorisse.fr/go-calculator/pkg/utils"
	"fmt"
	"regexp"
)

var InnerParenthesisRegex *regexp.Regexp
var SubGroupRegex *regexp.Regexp

func init() {
	var innerParenthesisRegexStr = fmt.Sprintf(`[+-]?(?:[a-z][a-z0-9]+)?\(-?[0-9.a-z%s :]+\)`,
		utils.SanitizeOperators(binary.KnownSymbols))

	InnerParenthesisRegex = regexp.MustCompile(innerParenthesisRegexStr)
	SubGroupRegex = regexp.MustCompile(`^ *(?::\d+|\( *:\d+ *\)) *$`)
}

func IsSubGroup(expression string) bool {
	return SubGroupRegex.MatchString(expression)
}
