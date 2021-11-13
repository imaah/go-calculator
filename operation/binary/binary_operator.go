package binary

import (
	"emorisse.fr/go-calculator/operation"
	"emorisse.fr/go-calculator/utils"
	"errors"
	"math"
)

//KnownSymbols all symbols that can be used for a binary operator
var KnownSymbols = []rune{'-', '+', '*', '/', '^', '%'}

var OperatorPriority = [][]rune{
	{'^'},
	{'*', '/', '%'},
	{'-', '+'},
}

type opBinary struct {
	operation.Operation
	Left, Right operation.Operation
	Symbol      rune
}

func (b *opBinary) Eval() *operation.Result {
	switch b.Symbol {
	case '+':
		return add(b.Left.Eval(), b.Right.Eval())
	case '-':
		return subtract(b.Left.Eval(), b.Right.Eval())
	case '/':
		return divide(b.Left.Eval(), b.Right.Eval())
	case '*':
		return multiply(b.Left.Eval(), b.Right.Eval())
	case '^':
		return pow(b.Left.Eval(), b.Right.Eval())
	case '%':
		return modulo(b.Left.Eval(), b.Right.Eval())
	}
	return nil
}

func multiply(left, right *operation.Result) *operation.Result {
	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() * right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " * " + right.GetString())
}

func divide(left, right *operation.Result) *operation.Result {
	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() / right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " / " + right.GetString())
}

func add(left, right *operation.Result) *operation.Result {
	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() + right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " + " + right.GetString())
}

func subtract(left, right *operation.Result) *operation.Result {
	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() - right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " - " + right.GetString())
}

func pow(left, right *operation.Result) *operation.Result {
	if left.IsNumber() && right.IsNumber() {
		var result = math.Pow(left.GetNumber(), right.GetNumber())
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " ^ " + right.GetString())
}

func modulo(left, right *operation.Result) *operation.Result {
	if left.IsNumber() && right.IsNumber() {
		var result = math.Mod(left.GetNumber(), right.GetNumber())
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " % " + right.GetString())
}

//New Creates a new binary operator
func New(symbol rune, left, right operation.Operation) (operation.Operation, error) {
	if left == nil || right == nil {
		return nil, errors.New("ArgumentIsNil")
	}

	if utils.RuneArrayContains(KnownSymbols, symbol) {
		operator := opBinary{
			Right:  right,
			Left:   left,
			Symbol: symbol,
		}
		return &operator, nil
	}
	return nil, errors.New("InvalidOperator")
}
