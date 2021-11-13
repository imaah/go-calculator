package number

import "emorisse.fr/calcul/operators"

type opNumber struct {
	operators.Operation
	value float64
}

func (n opNumber) Eval() *operators.OperationResult {
	return operators.NewNumberResult(n.value)
}

//New Creates a new number operator
func New(value float64) operators.Operation {
	return opNumber{
		value: value,
	}
}
