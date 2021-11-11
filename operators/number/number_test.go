package number

import "testing"

func TestOpNumber_Eval(t *testing.T) {
	var num = New(8)

	if num.Eval() != 8 {
		t.Logf("Should be %f but got %f", float64(8), num.Eval())
	}
}
