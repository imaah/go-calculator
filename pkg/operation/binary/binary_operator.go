package binary

import (
	"emorisse.fr/go-calculator/pkg/operation"
	"emorisse.fr/go-calculator/pkg/utils"
	"errors"
	"fmt"
	"math"
)

//OperatorPriority all symbols that can be used for a binary operator ordered by priority
var OperatorPriority = []rune{'^', '*', '/', '%', '-', '+'}

type OpBinary struct {
	operation.Operation
	Left, Right operation.Operation
	Symbol      rune
}

func (b *OpBinary) Eval() *operation.Result {
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

func (b OpBinary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.Left.String(), b.Symbol, b.Right.String())
}

func (b *OpBinary) multiply(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() * right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " * " + right.GetString())
}

func (b *OpBinary) divide(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() / right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " / " + right.GetString())
}

func (b *OpBinary) add(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() + right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " + " + right.GetString())
}

func (b *OpBinary) subtract(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = left.GetNumber() - right.GetNumber()
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " - " + right.GetString())
}

func (b *OpBinary) pow(right *operation.Result) *operation.Result {
	var left = b.Left.Eval()

	if left.IsNumber() && right.IsNumber() {
		var result = math.Pow(left.GetNumber(), right.GetNumber())
		return operation.NewNumberResult(result)
	}

	return operation.NewStringResult(left.GetString() + " ^ " + right.GetString())
}

func (b *OpBinary) modulo(right *operation.Result) *operation.Result {
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
		return nil, errors.New("ArgumentIsNil (binary)")
	}

	if utils.RuneArrayContains(OperatorPriority, symbol) {
		operator := OpBinary{
			Right:  right,
			Left:   left,
			Symbol: symbol,
		}
		return &operator, nil
	}
	return nil, errors.New(fmt.Sprintf("InvalidBinaryOperator (%c)", symbol))
}
