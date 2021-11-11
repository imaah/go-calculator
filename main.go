package main

import (
	"emorisse.fr/calcul/operators/binary"
	"emorisse.fr/calcul/operators/function"
	"emorisse.fr/calcul/operators/number"
	"emorisse.fr/calcul/operators/unary"
	"fmt"
)

func main() {
	u, _ := unary.New('-', number.New(10))
	op, _ := binary.New('+', number.New(20), u)

	fmt.Println(op.Eval())

	cos, _ := function.New("cos", number.New(0))
	fmt.Println(cos.Eval())
}
