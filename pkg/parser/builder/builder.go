package builder

import (
	"emorisse.fr/go-calculator/pkg/operation"
	"emorisse.fr/go-calculator/pkg/operation/binary"
	"emorisse.fr/go-calculator/pkg/operation/function"
	"emorisse.fr/go-calculator/pkg/operation/number"
	"emorisse.fr/go-calculator/pkg/operation/unary"
	"emorisse.fr/go-calculator/pkg/parser/regex"
	"emorisse.fr/go-calculator/pkg/parser/types"
	"emorisse.fr/go-calculator/pkg/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func BuildOperator(groups types.GroupMap, index int) (operation.Operation, error) {
	var key = fmt.Sprintf(":%d", index)
	var elem = groups[key]

	if regex.IsBinaryExpression(elem) {
		return GetBinaryExpression(elem, groups)
	}

	if regex.IsUnaryExpression(elem) {
		return GetUnaryExpression(elem, groups)
	}

	if regex.IsFunctionExpression(elem) {
		return GetFunctionExpression(elem, groups)
	}

	return ParseElement(elem, groups)
}

func GetBinaryExpression(expression string, groups types.GroupMap) (operation.Operation, error) {
	var left, right operation.Operation
	var err error
	var symbol = ' '

	var regexGroups = regex.BinaryExpressionRegex.FindAllStringSubmatch(expression, -1)[0]

	if len(regexGroups) == 4 {
		left, err = ParseElement(regexGroups[1], groups)
		symbol = rune(regexGroups[2][0])
		right, err = ParseElement(regexGroups[3], groups)
	}

	if err != nil {
		return nil, err
	}

	if left != nil && right != nil && symbol != ' ' {
		return binary.New(symbol, left, right)
	}

	return nil, errors.New("IllegalExpression")
}

func GetUnaryExpression(expression string, groups types.GroupMap) (operation.Operation, error) {
	var right operation.Operation
	var err error
	var symbol = ' '

	var regexGroups = regex.UnaryExpressionRegex.FindAllStringSubmatch(expression, -1)[0]

	if len(regexGroups) == 3 {
		symbol = rune(regexGroups[1][0])
		right, err = ParseElement(regexGroups[2], groups)
	}

	if err != nil {
		return nil, err
	}

	if right != nil && symbol != ' ' {
		return unary.New(symbol, right)
	}

	return nil, errors.New("IllegalExpression")
}

func GetFunctionExpression(expression string, groups types.GroupMap) (operation.Operation, error) {
	var value operation.Operation
	var err error
	var functionName = ""

	var regexGroups = regex.FunctionExpressionRegex.FindAllStringSubmatch(expression, -1)[0]

	if len(regexGroups) == 3 {
		functionName = regexGroups[1]
		value, err = ParseElement(regexGroups[2], groups)
	}

	if err != nil {
		return nil, err
	}

	if value != nil && functionName != "" {
		return function.New(functionName, value)
	}

	return nil, errors.New("IllegalExpression")
}

func ParseElement(expression string, groups types.GroupMap) (operation.Operation, error) {
	var elem operation.Operation
	var err error

	fmt.Println("number")

	if regex.IsSubGroup(expression) {
		var index = utils.ExtractIndex(expression)
		elem, err = BuildOperator(groups, index)
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
