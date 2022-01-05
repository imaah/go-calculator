package builder

import (
	"emorisse.fr/go-calculator/pkg/errors"
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
	return BuildOperatorStr(m, elem)
}

func BuildOperatorStr(m map[uint]string, elem string) (operation.Operation, error) {
	if strings.TrimSpace(elem) == "" {
		return nil, errors.EmptyStringError
	}

	if utils.IsSubGroup(elem) {
		var val, _ = utils.GetSubGroupValue(strings.TrimSpace(elem))
		return BuildOperator(m, val)
	}

	if regex.IsFunctionExpression(elem) {
		return BuildFunction(m, elem)
	}

	if regex.IsBinaryExpression(elem) {
		return BuildBinary(m, elem)
	}

	var num, err = strconv.ParseFloat(strings.TrimSpace(elem), 64)

	if err != nil {
		return BuildBinary(m, elem)
	}

	return number.New(num), nil
}

func BuildFunction(m map[uint]string, elem string) (operation.Operation, error) {
	var funName = strings.TrimSpace(utils.SubGroupRegex.ReplaceAllString(elem, ""))
	var val, _ = utils.GetSubGroupValue(utils.SubGroupRegex.FindString(elem))

	var op, err = BuildOperator(m, val)

	if err != nil {
		return nil, err
	}

	return function.New(funName, op)
}

func BuildBinary(m map[uint]string, elem string) (operation.Operation, error) {
	if strings.TrimSpace(elem) == "" {
		return nil, errors.EmptyStringError
	}
	var op = rune(strings.TrimSpace(utils.SubGroupRegex.ReplaceAllString(elem, ""))[0])
	var subs = utils.SubGroupRegex.FindAllString(elem, -1)

	if len(subs) < 2 {
		return nil, errors.InvalidString
	}

	var left, right operation.Operation
	var err error

	right, err = buildSub(m, subs[1])

	if err != nil {
		return nil, err
	}

	left, err = buildSub(m, subs[0])

	if err != nil {
		if err == errors.EmptyStringError {
			return unary.New(op, right)
		}
		return nil, err
	}

	return binary.New(op, left, right)
}

func buildSub(m map[uint]string, elem string) (operation.Operation, error) {
	if strings.TrimSpace(elem) == "" {
		return nil, errors.EmptyStringError
	}
	if utils.IsSubGroup(elem) {
		var val, _ = utils.GetSubGroupValue(elem)

		for utils.IsSubGroup(m[val]) {
			var tmpVal, err = utils.GetSubGroupValue(m[val])

			if err != nil {
				return BuildOperatorStr(m, m[val])
			}
			val = tmpVal
		}

		return BuildOperator(m, val)
	}
	return BuildOperatorStr(m, elem)
}
