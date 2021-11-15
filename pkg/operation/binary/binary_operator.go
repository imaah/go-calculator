package binary

import (
	"emorisse.fr/go-calculator/pkg/operation"
	"emorisse.fr/go-calculator/pkg/utils"
	"errors"
	"fmt"
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
		return b.add(b.Right.Eval())
	case '-':
		return b.subtract(b.Right.Eval())
	case '/':
		return b.divide(b.Right.Eval())
	case '*':
		return b.multiply(b.Right.Eval())
	case '^':
		return b.pow(b.Right.Eval())
	case '%':
		return b.modulo(b.Right.Eval())
	}
	return nil
}

func (b opBinary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.Left.String(), b.Symbol, b.Right.String())
}

func (b *opBinary) multiply(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() * right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " * " + right.GetString())
}

func (b *opBinary) divide(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() / right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " / " + right.GetString())
}

func (b *opBinary) add(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() + right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " + " + right.GetString())
}

func (b *opBinary) subtract(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() - right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " - " + right.GetString())
}

func (b *opBinary) pow(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = math.Pow(left.GetNumber(), right.GetNumber())
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " ^ " + right.GetString())
}

func (b *opBinary) modulo(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

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
