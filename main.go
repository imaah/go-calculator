package main

import (
	"emorisse.fr/calcul/operators/binary"
	"emorisse.fr/calcul/operators/function"
	"emorisse.fr/calcul/operators/number"
	"fmt"
)

func main() {
	// ((2 + 5) * cos(4)) / (sin(24) - (4 * 12))
	tpf, _ := binary.New('+', number.New(2), number.New(5))
	c4, _ := function.New("cos", number.New(4))
	tpftc4, _ := binary.New('*', tpf, c4)
	s24, _ := function.New("sin", number.New(24))
	ftt, _ := binary.New('*', number.New(4), number.New(12))
	s24mftt, _ := binary.New('-', s24, ftt)

	div, _ := binary.New('/', tpftc4, s24mftt)

	fmt.Println(div.Eval())
}
