package unary

import (
	"emorisse.fr/go-calculator/pkg/operation/number"
	"testing"
)

func TestNew__NilArgument(t *testing.T) {
	var _, err = New('-', nil)

	if err == nil {
		t.Logf("Shouldn't be nil but got nil")
		t.Fail()
	}
}

func TestNew__WrongSymbol(t *testing.T) {
	var _, err = New('a', number.New(4))

	if err == nil {
		t.Logf("Shouldn't be nil but go %s", err)
		t.Fail()
	}
}

func TestOpUnary_Eval(t *testing.T) {
	test('-', 4, -4, t)
	test('+', 4, 4, t)
}

func test(symbol rune, value, expected float64, t *testing.T) {
	var una, err = New(symbol, number.New(value))

	if err != nil {
		t.Logf("Should be nil but got %s", err)
		t.Fail()
	}

	if una.Eval().GetNumber() != expected {
		t.Logf("Should be %f but got %f", expected, una.Eval().GetNumber())
	}
}
