package number

import (
	"emorisse.fr/go-calculator/operation"
	"fmt"
)

type opNumber struct {
	operation.Operation
	value float64
}

func (n opNumber) Eval() *operation.Result {
	return operation.NewNumberResult(n.value)
}

func (n opNumber) String() string {
	return fmt.Sprintf("%f", n.value)
}

//New Creates a new number operator
func New(value float64) operation.Operation {
	return opNumber{
		value: value,
	}
}
