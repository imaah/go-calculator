package parser

import (
	"emorisse.fr/calcul/operators"
	"emorisse.fr/calcul/operators/binary"
	"emorisse.fr/calcul/operators/number"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var regexInnerParenthesis *regexp.Regexp
var binaryExpressionRegex *regexp.Regexp
var subGroupRegex *regexp.Regexp

var binarySymbols string

func init() {
	regexInnerParenthesis = regexp.MustCompile("[a-z]*\\([0-9.a-z+\\-*/ :]+\\)")
	subGroupRegex = regexp.MustCompile("^:[0-9]+$")

	for _, symbol := range binary.KnownSymbols {
		binarySymbols += string(symbol)
	}

	var binaryRegexStr = fmt.Sprintf("^\\(?(:[0-9]+|[0-9]+(?:.[0-9]+)?) *([%s]) *(:[0-9]+|[0-9]+(?:.[0-9]+)?)\\)?$",
		binarySymbols)

	binaryExpressionRegex = regexp.MustCompile(binaryRegexStr)
}

func Parse(str string) (operators.Operation, error) {
	var group []byte
	fmt.Println(str)
	var byteStr = makeCalculusPriority([]byte(str))
	str = string(byteStr)
	fmt.Println(string(byteStr))

	var groups = make(map[string]string)
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
	var lastElem = len(groups)
	return buildOperator(groups, lastElem)
}

func buildOperator(groups map[string]string, index int) (operators.Operation, error) {
	var key = fmt.Sprintf(":%d", index)
	var elemStr = groups[key]
	var elem = []byte(elemStr)
	fmt.Println(elemStr)

	if isBinaryExpression(elem) {
		return getBinaryExpression(elem, groups)
	}

	//if isUnaryExpression(elem) {
	//	return getUnaryExpression(elem);
	//}
	//TODO : check if it is a unary expression
	//TODO : check if it is a function expression
	//TODO : check if it is a number expression

	//TODO : find the other operations inside the current expression then build it !
	//TODO : build the operation according to the expression found.

	return nil, errors.New("UnimplementedMethod")
}

func isBinaryExpression(expression []byte) bool {
	return binaryExpressionRegex.Match(expression)
}

func getBinaryExpression(expression []byte, groups map[string]string) (operators.Operation, error) {
	var left, right operators.Operation
	var err error
	var symbol = ' '

	var regexGroups = binaryExpressionRegex.FindAllSubmatch(expression, -1)[0]

	for _, group := range regexGroups {
		fmt.Println(string(group))
	}

	if len(regexGroups) == 4 {
		left, err = parseElement(regexGroups[1], groups)
		symbol = rune(regexGroups[2][0])
		right, err = parseElement(regexGroups[3], groups)
	}

	if err != nil {
		return nil, err
	}
	fmt.Println("hehe")
	fmt.Println(left, right, symbol)

	if left != nil && right != nil && symbol != ' ' {
		return binary.New(symbol, left, right)
	}

	return nil, errors.New("IllegalExpression")
}

func parseElement(expression []byte, groups map[string]string) (operators.Operation, error) {
	var elem operators.Operation
	var err error

	if isSubGroup(expression) {
		var index int
		_, _ = fmt.Sscanf(string(expression), ":%d", &index)

		elem, err = buildOperator(groups, index)
	} else {
		var num float64
		_, err = fmt.Sscanf(string(expression), "%f", &num)

		fmt.Println(num)

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

func cleanMap(groups map[string]string) map[string]string {
	var regex, _ = regexp.Compile("\\(:\\d+\\)")
	for key, val := range groups {
		if regex.Match([]byte(val)) {
			delete(groups, key)
		}
	}
	return groups
}

func makeCalculusPriority(bytes []byte) []byte {
	bytes = makeSubGroups(bytes, []string{"*", "/", "^"})
	bytes = makeSubGroups(bytes, []string{"+", "-"})
	return bytes
}

func makeSubGroups(bytes []byte, operators []string) []byte {
	var count = 0
	var baseRegex = "(?:[0-9.]+|\\(.*\\)) *[%s] (?:[0-9.]+|\\(.*\\))"
	var regexStr = fmt.Sprintf(baseRegex, strings.Join(operators, ""))
	var regex, _ = regexp.Compile(regexStr)

	var str = string(bytes)

	for _, ope := range operators {
		count += strings.Count(str, ope)
	}

	for i := 0; i < count; i++ {
		bytes = regex.ReplaceAllFunc(bytes, addParenthesisAround)
	}

	return bytes
}

func addParenthesisAround(group []byte) []byte {
	return []byte("(" + string(group) + ")")
}
