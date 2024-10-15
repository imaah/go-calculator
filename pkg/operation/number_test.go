package operation_test

import (
	"testing"

	"github.com/imaah/go-calculator/pkg/operation"
)

func TestOpNumber_Eval(t *testing.T) {
	var num = operation.NewNumber(8)

	if num.Eval().GetNumber() != 8 {
		t.Logf("Should be %f but got %f", float64(8), num.Eval().GetNumber())
	}
}
