package binary

import (
	"emorisse.fr/calcul/operators"
	"emorisse.fr/calcul/utils"
	"errors"
)

//KnownSymbols all symbols that can be used for a binary operator
var KnownSymbols = []rune{'-', '+', '*', '/', '^', '%'}

var OperatorPriority = [][]rune{
	{'^'},
	{'*', '/', '%'},
	{'-', '+'},
}

type opBinary struct {
	operators.Operation
	Left, Right operators.Operation
	Symbol      rune
}

func (b *opBinary) Eval() *operators.OperationResult {
	switch b.Symbol {
	case '+':
		return add(b.Left.Eval(), b.Right.Eval())
	case '-':
		return subtract(b.Left.Eval(), b.Right.Eval())
	case '/':
		return divide(b.Left.Eval(), b.Right.Eval())
	case '*':
		return multiply(b.Left.Eval(), b.Right.Eval())
		//case '^':
		//	return math.Pow(b.Left.Eval(), b.Right.Eval())
		//case '%':
		//	return math.Mod(b.Left.Eval(), b.Right.Eval())
	}
	return nil
}

func multiply(left, right *operators.OperationResult) *operators.OperationResult {
	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() * right.GetNumber()
		return operators.NewNumberResult(result)
	}

	return operators.NewStringResult(left.GetString() + " * " + right.GetString())
}

func divide(left, right *operators.OperationResult) *operators.OperationResult {
	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() / right.GetNumber()
		return operators.NewNumberResult(result)
	}

	return operators.NewStringResult(left.GetString() + " / " + right.GetString())
}

func add(left, right *operators.OperationResult) *operators.OperationResult {
	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() + right.GetNumber()
		return operators.NewNumberResult(result)
	}

	return operators.NewStringResult(left.GetString() + " + " + right.GetString())
}

func subtract(left, right *operators.OperationResult) *operators.OperationResult {
	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() - right.GetNumber()
		return operators.NewNumberResult(result)
	}

	return operators.NewStringResult(left.GetString() + " - " + right.GetString())
}

//New Creates a new binary operator
func New(symbol rune, left, right operators.Operation) (operators.Operation, error) {
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
