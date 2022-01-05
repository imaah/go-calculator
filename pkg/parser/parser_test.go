package parser

import (
	"math"
	"testing"
)

var workingSamples = map[string]float64{
	"((2 + 5) * cos(4)) / (sin(24) - (4 * 12))": ((2 + 5) * math.Cos(4)) / (math.Sin(24) - (4 * 12)),
	"1+2":                 3,
	"50 / 4":              12.5,
	"cos(12)":             math.Cos(12),
	"-cos(4)":             -math.Cos(4),
	"(125 - 4) * -sin(8)": (125 - 4) * -math.Sin(8),
	"8":                   8,
	"-8":                  -8,
}

var failingSamples = [...]string{
	"",
	"bob(18)",
	"10 $ 4",
	"-",
	"1 - *",
	"((2 + 5) * cos(/)) / (sin(24) - (+ * 12))",
	"hello world!",
	"ww",
	"w",
}

func TestParse__passes(t *testing.T) {
	for sample := range workingSamples {
		var _, err = Parse(sample)

		if err != nil {
			t.Logf("%s should pass but doesn't (%s)", sample, err)
			t.Fail()
		}
	}
}

func TestParse__fails(t *testing.T) {
	for _, sample := range failingSamples {
		var _, err = Parse(sample)

		if err == nil {
			t.Logf("%s should fail but doesn't", sample)
			t.Fail()
		}
	}
}

func TestParse__result(t *testing.T) {
	for sample, result := range workingSamples {
		var calc, _ = Parse(sample)

		if calc.Eval().GetNumber() != result {
			t.Logf("Got %s but expected %f", calc.Eval().GetString(), result)
			t.Fail()
		}
	}
}
