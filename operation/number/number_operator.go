package number

import "emorisse.fr/go-calculator/operation"

type opNumber struct {
	operation.Operation
	value float64
}

func (n opNumber) Eval() *operation.Result {
	return operation.NewNumberResult(n.value)
}

//New Creates a new number operator
func New(value float64) operation.Operation {
	return opNumber{
		value: value,
	}
}
