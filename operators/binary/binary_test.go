package binary

import (
	"emorisse.fr/calcul/operators/number"
	"testing"
)

func TestNew__WrongSymbol(t *testing.T) {
	var _, err = New('a', number.New(4), number.New(4))

	if err == nil {
		t.Logf("Shouldn't be nil but go %s", err)
		t.Fail()
	}

}

func TestNew__NilArgument(t *testing.T) {
	var _, err = New('-', nil, nil)

	if err == nil {
		t.Logf("Shouldn't be nil but got nil")
		t.Fail()
	}
}

func TestOpBinary_Eval(t *testing.T) {
	test('+', 10, 4, 14, t)
	test('-', 10, 4, 6, t)
	test('*', 10, 4, 40, t)
	test('/', 10, 4, 2.5, t)
	test('^', 3, 2, 9, t)
	test('%', 10, 3, 1, t)
}

func test(symbol rune, val1, val2, expected float64, t *testing.T) {
	var bin, err = New(symbol, number.New(val1), number.New(val2))

	if err != nil {
		t.Logf("Should be nil but got %s", err)
		t.Fail()
	}

	if bin.Eval().GetNumber() != expected {
		t.Logf("Should be %f but got %f", expected, bin.Eval().GetNumber())
	}
}
