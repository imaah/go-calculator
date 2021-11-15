package function

import (
	"emorisse.fr/go-calculator/pkg/operation/number"
	"math"
	"testing"
)

func TestNew__WrongFunction(t *testing.T) {
	var _, err = New("a", number.New(1))

	if err == nil {
		t.Logf("Shouldn't be nil but go %s", err)
		t.Fail()
	}

}

func TestNew__NilArgument(t *testing.T) {
	var _, err = New("cos", nil)

	if err == nil {
		t.Logf("Shouldn't be nil")
		t.Fail()
	}
}

func TestOpFunction_Eval(t *testing.T) {
	test("cos", 0, math.Cos(0), t)
	test("sin", 0, math.Sin(0), t)
	test("tan", 0, math.Tan(0), t)
}

func TestNewUsingTempFunc(t *testing.T) {
	testFunc(func(val float64) float64 { return double(double(val)) }, 4, 16, t)
	testFunc(func(val float64) float64 { return val / 4 }, 4, 1, t)
	testFunc(func(val float64) float64 { return val }, 4, 4, t)
	testFunc(double, 4, 8, t)
}

func TestRegisterFunction__NilArgument(t *testing.T) {
	var err = RegisterFunction("double", nil)

	if err == nil {
		t.Logf("Shouldn't be nil")
		t.Fail()
	}
}

func TestRegisterFunction__Duplicate(t *testing.T) {
	var err = RegisterFunction("sin", double)

	if err == nil {
		t.Logf("Shouldn't be nil")
		t.Fail()
	}
}

func TestRegisterFunction(t *testing.T) {
	var err = RegisterFunction("double", double)

	if err != nil {
		t.Logf("Should be nil but got %s", err)
		t.Fail()
	}

	test("double", 4, 8, t)
}

func double(val float64) float64 {
	return val * 2
}

func testFunc(function OpFunc, val, expected float64, t *testing.T) {
	var bin = NewUsingTempFunc(function, number.New(val))

	if bin.Eval().GetNumber() != expected {
		t.Logf("Should be %f but got %f", expected, bin.Eval().GetNumber())
	}
}

func test(function string, val, expected float64, t *testing.T) {
	var bin, err = New(function, number.New(val))

	if err != nil {
		t.Logf("Should be nil but got %s", err)
		t.Fail()
	}

	if bin.Eval().GetNumber() != expected {
		t.Logf("Should be %f but got %f", expected, bin.Eval().GetNumber())
	}
}
