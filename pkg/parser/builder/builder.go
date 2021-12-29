package builder

import (
	"emorisse.fr/go-calculator/pkg/operation"
	"emorisse.fr/go-calculator/pkg/operation/binary"
	"emorisse.fr/go-calculator/pkg/operation/function"
	"emorisse.fr/go-calculator/pkg/operation/number"
	"emorisse.fr/go-calculator/pkg/operation/unary"
	"emorisse.fr/go-calculator/pkg/parser/regex"
	"emorisse.fr/go-calculator/pkg/utils"
	"strconv"
	"strings"
)

func BuildOperator(m map[uint]string, index uint) (operation.Operation, error) {
	var elem = m[index]

	if utils.IsSubGroup(elem) {
		var val, _ = utils.GetSubGroupValue(strings.TrimSpace(elem))
		return BuildOperator(m, val)
	}

	if regex.IsFunctionExpression(elem) {
		return BuildFunction(m, index)
	}

	if regex.IsBinaryExpression(elem) {
		return BuildBinary(m, index)
	}

	var num, err = strconv.ParseFloat(strings.TrimSpace(elem), 64)

	if err != nil {
		return nil, err
	}

	return number.New(num), nil
}

func BuildFunction(m map[uint]string, index uint) (operation.Operation, error) {
	var elem = m[index]
	var funName = strings.TrimSpace(utils.SubGroupRegex.ReplaceAllString(elem, ""))
	var val, _ = utils.GetSubGroupValue(utils.SubGroupRegex.FindString(elem))

	var op, err = BuildOperator(m, val)

	if err != nil {
		return nil, err
	}

	return function.New(funName, op)
}

func BuildBinary(m map[uint]string, index uint) (operation.Operation, error) {
	var elem = m[index]
	var op = rune(strings.TrimSpace(utils.SubGroupRegex.ReplaceAllString(elem, ""))[0])
	var subs = utils.SubGroupRegex.FindAllString(elem, -1)
	var leftVal uint

	leftVal, _ = utils.GetSubGroupValue(subs[0])

	for utils.IsSubGroup(m[leftVal]) {
		leftVal, _ = utils.GetSubGroupValue(m[leftVal])
	}

	var rightVal, _ = utils.GetSubGroupValue(subs[1])
	var left, right operation.Operation
	var err error

	right, err = BuildOperator(m, rightVal)

	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(m[leftVal])) == 0 {
		return unary.New(op, right)
	}

	left, err = BuildOperator(m, leftVal)

	if err != nil {
		return nil, err
	}

	return binary.New(op, left, right)
}
