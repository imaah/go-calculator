package parser

import (
	"emorisse.fr/calcul/operators"
	"emorisse.fr/calcul/operators/binary"
	"emorisse.fr/calcul/operators/function"
	"emorisse.fr/calcul/operators/number"
	"emorisse.fr/calcul/operators/unary"
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
var functionExpressionRegex *regexp.Regexp
var findFunctionExpressionsRegex *regexp.Regexp
var subGroupRegex *regexp.Regexp

func init() {
	var innerParenthesisRegexStr = fmt.Sprintf("[+-]?(?:[a-z][a-z0-9]+)?\\(-?[0-9.a-z%s :]+\\)",
		strings.Replace(string(binary.KnownSymbols), "-", "\\-", 1))

	var binaryRegexStr = fmt.Sprintf("^\\(?(:[0-9]+|-?[0-9]+(?:.[0-9]+)?) *([%s]) *(:[0-9]+|-?[0-9]+(?:.[0-9]+)?)\\)?$",
		strings.Replace(string(binary.KnownSymbols), "-", "\\-", 1))

	var unaryRegexStr = fmt.Sprintf("^ *([%s]) *(?:\\(?(:[0-9]+|-?[0-9]+(?:.[0-9]+)?)\\)?)$",
		strings.Replace(string(unary.KnownSymbols), "-", "\\-", 1))

	regexInnerParenthesis = regexp.MustCompile(innerParenthesisRegexStr)
	binaryExpressionRegex = regexp.MustCompile(binaryRegexStr)
	unaryExpressionRegex = regexp.MustCompile(unaryRegexStr)
	functionExpressionRegex = regexp.MustCompile("^ *([a-z][a-z0-9]*) *\\((:[0-9]+|-?[0-9]+(?:.[0-9]+)?)\\)$")
	findFunctionExpressionsRegex = regexp.MustCompile(" *([a-z][a-z0-9]*) *\\(-?[0-9]+(?:.[0-9]+)?\\)")
	subGroupRegex = regexp.MustCompile("^:[0-9]+$")
}

func Parse(str string) (operators.Operation, error) {
	var group []byte
	var byteStr = adaptExpression([]byte(str))
	str = string(byteStr)

	var groups = make(groupMap)

	group = regexInnerParenthesis.Find(byteStr)

	for i := 0; group != nil; {
		i++
		var key = fmt.Sprintf(":%d", i)
		var part = string(group)
		groups[key] = part
		str = strings.Replace(str, part, fmt.Sprintf(":%d", i), 1)
		byteStr = []byte(str)
		group = regexInnerParenthesis.Find(byteStr)
	}

	groups = cleanMap(groups)

	var lastElem = getLastIndex(groups)
	return buildOperator(groups, lastElem)
}

func buildOperator(groups groupMap, index int) (operators.Operation, error) {
	var key = fmt.Sprintf(":%d", index)
	var elemStr = groups[key]
	var elem = []byte(elemStr)

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

func isFunctionExpression(expression []byte) bool {
	return functionExpressionRegex.Match(expression)
}

func getFunctionExpression(expression []byte, groups groupMap) (operators.Operation, error) {
	var value operators.Operation
	var err error
	var functionName = ""

	var regexGroups = functionExpressionRegex.FindAllSubmatch(expression, -1)[0]

	if len(regexGroups) == 3 {
		functionName = string(regexGroups[1])
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

func isUnaryExpression(expression []byte) bool {
	return unaryExpressionRegex.Match(expression)
}

func getUnaryExpression(expression []byte, groups groupMap) (operators.Operation, error) {
	var right operators.Operation
	var err error
	var symbol = ' '

	var regexGroups = unaryExpressionRegex.FindAllSubmatch(expression, -1)[0]

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

func isBinaryExpression(expression []byte) bool {
	return binaryExpressionRegex.Match(expression)
}

func getBinaryExpression(expression []byte, groups groupMap) (operators.Operation, error) {
	var left, right operators.Operation
	var err error
	var symbol = ' '

	var regexGroups = binaryExpressionRegex.FindAllSubmatch(expression, -1)[0]

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

func parseElement(expression []byte, groups groupMap) (operators.Operation, error) {
	var elem operators.Operation
	var err error

	if isSubGroup(expression) {
		var index = extractIndex(string(expression))

		elem, err = buildOperator(groups, index)
	} else {
		var exprStr = strings.Trim(string(expression), "()")
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

func isSubGroup(expression []byte) bool {
	return subGroupRegex.Match(expression)
}

func getLastIndex(groups groupMap) int {
	var max = 0

	for key := range groups {
		var curr = extractIndex(key)
		if curr > max {
			max = curr
		}
	}

	return max
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
		if regex.Match([]byte(val)) {
			delete(groups, key)

			var redirectTo = fmt.Sprintf(":%d", extractIndex(val))
			groups = removeInOther(key, redirectTo, groups)
		}
	}
	return groups
}

func removeInOther(remove, redirectTo string, groups groupMap) groupMap {
	for key, val := range groups {
		if strings.Contains(val, remove) {
			groups[key] = strings.Replace(val, remove, redirectTo, 1)
		}
	}
	return groups
}

func adaptExpression(expression []byte) []byte {
	expression = makeCalculusPriority(expression)
	expression = addParenthesisAroundFunctions(expression)
	return addParenthesisAround(expression)
}

func makeCalculusPriority(expression []byte) []byte {
	for _, oper := range binary.OperatorPriority {
		expression = makeSubGroups(expression, oper)
	}
	return expression
}

func makeSubGroups(expression []byte, operators []rune) []byte {
	var count = 0
	var baseRegex = "(?:[a-z][a-z0-9]+)?(?:-?[0-9.]+|\\(.*\\)) *[%s] (?:[a-z][a-z0-9]+)?(?:-?[0-9.]+|\\(.*\\))"
	var regexStr = fmt.Sprintf(baseRegex, string(operators))
	var regex, _ = regexp.Compile(regexStr)

	var str = string(expression)

	for _, ope := range operators {
		count += strings.Count(str, string(ope))
	}

	for i := 0; i < count; i++ {
		expression = regex.ReplaceAllFunc(expression, addParenthesisAround)
	}

	return expression
}

func addParenthesisAroundFunctions(expression []byte) []byte {
	return findFunctionExpressionsRegex.ReplaceAllFunc(expression, addParenthesisAround)
}

func addParenthesisAround(group []byte) []byte {
	return []byte("(" + strings.Trim(string(group), " ") + ")")
}
