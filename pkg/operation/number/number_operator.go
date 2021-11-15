package number

import (
	"emorisse.fr/go-calculator/pkg/operation"
	"fmt"
)

type OpNumber struct {
	operation.Operation
	value float64
}

func (n OpNumber) Eval() *operation.Result {
	return operation.NewNumberResult(n.value)
}

func (n OpNumber) String() string {
	return fmt.Sprintf("%f", n.value)
}

//New Creates a new number operator
func New(value float64) operation.Operation {
	return OpNumber{
		value: value,
	}
}
