package preprocessor

import (
	"emorisse.fr/go-calculator/pkg/operation/binary"
	"emorisse.fr/go-calculator/pkg/parser/regex"
	"emorisse.fr/go-calculator/pkg/parser/types"
	"emorisse.fr/go-calculator/pkg/utils"
	"fmt"
	"regexp"
	"strings"
)

func ProcessBinary(str string, lastIndex int, groups types.GroupMap) (string, int, types.GroupMap) {
	for _, operators := range binary.OperatorPriority {
		str, lastIndex, groups = processPriority(str, lastIndex, groups, operators)
	}

	return str, lastIndex, groups
}

func ProcessUnary(str string, lastIndex int, groups types.GroupMap) (string, int, types.GroupMap) {
	if regex.IsUnaryExpression(str) || regex.IsBinaryExpression(str) {
		return str, lastIndex, groups
	}

	var una = regex.FindUnaryExpressionRegex.FindAllString(str, -1)

	for i, group := range una {
		key := fmt.Sprintf(":%d", i+lastIndex)
		groups[key] = group
		str = strings.Replace(str, group, key, 1)
	}

	return str, lastIndex + len(una), groups
}

func ProcessFunctions(str string, lastIndex int, groups types.GroupMap) (string, int, types.GroupMap) {
	var functions = regex.FindFunctionExpressionsRegex.FindAllString(str, -1)

	for i, group := range functions {
		key := fmt.Sprintf(":%d", i+lastIndex)
		groups[key] = group
		str = strings.Replace(str, group, key, 1)
	}

	return str, lastIndex + len(functions), groups
}

func ProcessParenthesis(str string, lastIndex int, groups types.GroupMap) (string, int, types.GroupMap) {
	var group = regex.InnerParenthesisRegex.FindString(str)
	var i int

	for i = 0; group != ""; i++ {
		key := fmt.Sprintf(":%d", i+lastIndex)
		groups[key] = group
		str = strings.Replace(str, group, key, 1)
		lastIndex++
		group = regex.InnerParenthesisRegex.FindString(str)
	}

	return str, lastIndex + i, groups
}

func processPriority(str string, lastIndex int, groups types.GroupMap, operators []rune) (string, int, types.GroupMap) {
	regexStr := fmt.Sprintf("-?(:\\d+|\\d+(?:\\.\\d*)?) *[%s] *-?(:\\d+|-?\\d+(?:\\.\\d*)?)",
		utils.SanitizeOperators(operators))
	var regexBinaryOperator = regexp.MustCompile(regexStr)
	var fullMatch = regexp.MustCompile("^" + regexStr + "$")

	if fullMatch.MatchString(str) {
		return str, lastIndex, groups
	}

	var group = regexBinaryOperator.FindString(str)
	var i int

	for i = 0; group != ""; i++ {
		key := fmt.Sprintf(":%d", i+lastIndex)
		groups[key] = group
		str = strings.Replace(str, group, key, 1)
		group = regexBinaryOperator.FindString(str)
	}

	return str, lastIndex + i, groups
}
