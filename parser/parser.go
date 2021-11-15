package parser

import (
	"emorisse.fr/go-calculator/operation"
	"emorisse.fr/go-calculator/operation/binary"
	"emorisse.fr/go-calculator/operation/function"
	"emorisse.fr/go-calculator/operation/number"
	"emorisse.fr/go-calculator/operation/unary"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type groupMap map[string]string

var regexInnerParenthesis *regexp.Regexp
var binaryExpressionRegex *regexp.Regexp
var unaryExpressionRegex *regexp.Regexp
var findUnaryExpressionRegex *regexp.Regexp
var functionExpressionRegex *regexp.Regexp
var findFunctionExpressionsRegex *regexp.Regexp
var subGroupRegex *regexp.Regexp

func init() {
	var innerParenthesisRegexStr = fmt.Sprintf("[+-]?(?:[a-z][a-z0-9]+)?\\(-?[0-9.a-z%s :]+\\)",
		sanitizeOperators(binary.KnownSymbols))

	var binaryRegexStr = fmt.Sprintf("^\\(?(:[0-9]+|-?[0-9]+(?:.[0-9]+)?) *([%s]) *(:[0-9]+|-?[0-9]+(?:.[0-9]+)?)\\)?$",
		sanitizeOperators(binary.KnownSymbols))

	var unaryRegexStr = fmt.Sprintf("([%s])(?:(:[0-9]+|-?[0-9]+(?:.[0-9]+)?))",
		sanitizeOperators(unary.KnownSymbols))

	regexInnerParenthesis = regexp.MustCompile(innerParenthesisRegexStr)
	binaryExpressionRegex = regexp.MustCompile(binaryRegexStr)
	unaryExpressionRegex = regexp.MustCompile("^" + unaryRegexStr + "$")
	findUnaryExpressionRegex = regexp.MustCompile(unaryRegexStr)
	functionExpressionRegex = regexp.MustCompile("^ *([a-z][a-z0-9]*) *\\((:[0-9]+|-?[0-9]+(?:.[0-9]+)?)\\)$")
	findFunctionExpressionsRegex = regexp.MustCompile("([a-z][a-z0-9]*) *\\(-?[0-9]+(?:.[0-9]+)?\\)")
	subGroupRegex = regexp.MustCompile("^:[0-9]+$")
}

func Parse(str string) (operation.Operation, error) {
	var groups = make(groupMap)
	var lastIndex = 1

	str, lastIndex, groups = processFunctions(str, lastIndex, groups)
	str, lastIndex, groups = processParenthesis(str, lastIndex, groups)

	for k, group := range groups {
		groups[k], lastIndex, groups = processBinary(group, lastIndex, groups)
	}
	str, lastIndex, groups = processBinary(str, lastIndex, groups)

	for k, group := range groups {
		groups[k], lastIndex, groups = processUnary(group, lastIndex, groups)
	}
	str, lastIndex, groups = processUnary(str, lastIndex, groups)

	var key = fmt.Sprintf(":%d", lastIndex)
	groups[key] = str

	groups = cleanMap(groups)
//	fmt.Println(groups)
	return buildOperator(groups, lastIndex)
}

func processFunctions(str string, lastIndex int, groups groupMap) (string, int, groupMap) {
	var functions = findFunctionExpressionsRegex.FindAllString(str, -1)

	for i, group := range functions {
		key := fmt.Sprintf(":%d", i+lastIndex)
		groups[key] = group
		str = strings.Replace(str, group, key, 1)
	}

	return str, lastIndex + len(functions), groups
}

func processParenthesis(str string, lastIndex int, groups groupMap) (string, int, groupMap) {
	var group = regexInnerParenthesis.FindString(str)
	var i int

	for i = 0; group != ""; i++ {
		key := fmt.Sprintf(":%d", i+lastIndex)
		groups[key] = group
		str = strings.Replace(str, group, key, 1)
		lastIndex++
		group = regexInnerParenthesis.FindString(str)
	}

	return str, lastIndex + i, groups
}

func processUnary(str string, lastIndex int, groups groupMap) (string, int, groupMap) {
	if isUnaryExpression(str) || isBinaryExpression(str) {
		return str, lastIndex, groups
	}

	var una = findUnaryExpressionRegex.FindAllString(str, -1)

	for i, group := range una {
		key := fmt.Sprintf(":%d", i+lastIndex)
		groups[key] = group
		str = strings.Replace(str, group, key, 1)
	}

	return str, lastIndex + len(una), groups
}

func processBinary(str string, lastIndex int, groups groupMap) (string, int, groupMap) {
	for _, operators := range binary.OperatorPriority {
		str, lastIndex, groups = processPriority(str, lastIndex, groups, operators)
	}

	return str, lastIndex, groups
}

func processPriority(str string, lastIndex int, groups groupMap, operators []rune) (string, int, groupMap) {
	regexStr := fmt.Sprintf("-?(:\\d+|\\d+(?:\\.\\d*)?) *[%s] *-?(:\\d+|-?\\d+(?:\\.\\d*)?)",
		sanitizeOperators(operators))
	var regex = regexp.MustCompile(regexStr)
	var fullMatch = regexp.MustCompile("^" + regexStr + "$")

	if fullMatch.MatchString(str) {
		return str, lastIndex, groups
	}

	var group = regex.FindString(str)
	var i int

	for i = 0; group != ""; i++ {
		key := fmt.Sprintf(":%d", i+lastIndex)
		groups[key] = group
		str = strings.Replace(str, group, key, 1)
		group = regex.FindString(str)
	}

	return str, lastIndex + i, groups
}

func buildOperator(groups groupMap, index int) (operation.Operation, error) {
	var key = fmt.Sprintf(":%d", index)
	var elem = groups[key]

	if isBinaryExpression(elem) {
		return getBinaryExpression(elem, groups)
	}

	if isUnaryExpression(elem) {
		return getUnaryExpression(elem, groups)
	}

	if isFunctionExpression(elem) {
		return getFunctionExpression(elem, groups)
	}

	return parseElement(elem, groups)
}

func isFunctionExpression(expression string) bool {
	return functionExpressionRegex.MatchString(expression)
}

func getFunctionExpression(expression string, groups groupMap) (operation.Operation, error) {
	var value operation.Operation
	var err error
	var functionName = ""

	var regexGroups = functionExpressionRegex.FindAllStringSubmatch(expression, -1)[0]

	if len(regexGroups) == 3 {
		functionName = regexGroups[1]
		value, err = parseElement(regexGroups[2], groups)
	}

	if err != nil {
		return nil, err
	}

	if value != nil && functionName != "" {
		return function.New(functionName, value)
	}

	return nil, errors.New("IllegalExpression")
}

func isUnaryExpression(expression string) bool {
	return unaryExpressionRegex.MatchString(expression)
}

func getUnaryExpression(expression string, groups groupMap) (operation.Operation, error) {
	var right operation.Operation
	var err error
	var symbol = ' '

	var regexGroups = unaryExpressionRegex.FindAllStringSubmatch(expression, -1)[0]

	if len(regexGroups) == 3 {
		symbol = rune(regexGroups[1][0])
		right, err = parseElement(regexGroups[2], groups)
	}

	if err != nil {
		return nil, err
	}

	if right != nil && symbol != ' ' {
		return unary.New(symbol, right)
	}

	return nil, errors.New("IllegalExpression")
}

func isBinaryExpression(expression string) bool {
	return binaryExpressionRegex.MatchString(expression)
}

func getBinaryExpression(expression string, groups groupMap) (operation.Operation, error) {
	var left, right operation.Operation
	var err error
	var symbol = ' '

	var regexGroups = binaryExpressionRegex.FindAllStringSubmatch(expression, -1)[0]

	if len(regexGroups) == 4 {
		left, err = parseElement(regexGroups[1], groups)
		symbol = rune(regexGroups[2][0])
		right, err = parseElement(regexGroups[3], groups)
	}

	if err != nil {
		return nil, err
	}

	if left != nil && right != nil && symbol != ' ' {
		return binary.New(symbol, left, right)
	}

	return nil, errors.New("IllegalExpression")
}

func parseElement(expression string, groups groupMap) (operation.Operation, error) {
	var elem operation.Operation
	var err error

	if isSubGroup(expression) {
		var index = extractIndex(expression)

		elem, err = buildOperator(groups, index)
	} else {
		var exprStr = strings.Trim(expression, "()")
		var num float64
		num, err = strconv.ParseFloat(exprStr, 64)

		if err == nil {
			elem = number.New(num)
		}
	}

	if elem == nil && err == nil {
		err = errors.New("IllegalExpression")
	}

	if err != nil {
		return nil, err
	}

	return elem, nil
}

func isSubGroup(expression string) bool {
	return subGroupRegex.MatchString(expression)
}

func extractIndex(key string) int {
	var index int
	_, err := fmt.Sscanf(key, ":%d", &index)

	if err != nil {
		_, err = fmt.Sscanf(key, "(:%d)", &index)

		if err != nil {
			return 0
		}
	}

	return index
}

func cleanMap(groups groupMap) groupMap {
	var regex, _ = regexp.Compile("^\\( *:\\d+ *\\)$")
	for key, val := range groups {
			if regex.MatchString(val) {
					delete(groups, key)

					var redirectTo = fmt.Sprintf(":%d", extractIndex(val))
					groups = removeInOther(key, redirectTo, groups)
			}
	}
	return groups
}

func removeInOther(remove, redirectTo string, groups groupMap) groupMap {
	var regexStr = fmt.Sprintf(" *%s *", remove)
	var regexStr2 = fmt.Sprintf(" *%s[0-9]+ *", remove)
//      fmt.Println(regexStr, "->", redirectTo)
	var regex = regexp.MustCompile(regexStr)
	var regex2 = regexp.MustCompile(regexStr2)
	for key, val := range groups {
			if regex.MatchString(val) && !regex2.MatchString(val) {
//                      fmt.Println(val, remove)
					groups[key] = regex.ReplaceAllString(val, redirectTo)
			}
	}
	return groups
}

func sanitizeOperators(operators []rune) string {
	var strBuilder strings.Builder

	for _, ope := range operators {
		strBuilder.WriteString("\\")
		strBuilder.WriteRune(ope)
	}

	return strBuilder.String()
}
