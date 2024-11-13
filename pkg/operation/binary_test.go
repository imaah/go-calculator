package operation_test

import (
	"testing"

	"github.com/imaah/go-calculator/pkg/operation"
)

func TestNew_Binary_WrongSymbol(t *testing.T) {
	_, err := operation.NewBinary('a', operation.NewNumber(4), operation.NewNumber(4))

	if err == nil {
		t.Logf("Shouldn't be nil but go %s", err)
		t.Fail()
	}

}

func TestNew_Binary_NilArgument(t *testing.T) {
	_, err := operation.NewBinary('-', nil, nil)

	if err == nil {
		t.Logf("Shouldn't be nil but got nil")
		t.Fail()
	}
}

func TestOpBinary_Eval(t *testing.T) {
	testBinary('+', 10, 4, 14, t)
	testBinary('-', 10, 4, 6, t)
	testBinary('*', 10, 4, 40, t)
	testBinary('/', 10, 4, 2.5, t)
	testBinary('^', 3, 2, 9, t)
	testBinary('%', 10, 3, 1, t)
}

func testBinary(symbol rune, val1, val2, expected float64, t *testing.T) {
	bin, err := operation.NewBinary(symbol, operation.NewNumber(val1), operation.NewNumber(val2))

	if err != nil {
		t.Logf("Should be nil but got %s", err)
		t.Fail()
	}

	if bin.Eval() != expected {
		t.Logf("Should be %f but got %f", expected, bin.Eval())
	}
}
