package parser

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"unicode"
)

var alpha = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}
var hexNums = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'a', 'b', 'c', 'd', 'e', 'f'}
var decNums = hexNums[:10]
var operators = []byte{'+', '-', '*', '/', '%', '^'}
var unaryOperators = operators[:2]

func tokenizeNumberDec(str string, i int) (float64, int, error) {
	n := 0.0
	decimal := 0

	for i < len(str) {
		v := str[i]
		if '0' <= v && v <= '9' {
			if decimal != 0 {
				n = n + (float64(str[i]-'0') * math.Pow10(-decimal))
				decimal++
			} else {
				n = n*10 + float64(str[i]-'0')
			}
		} else if v == '.' {
			decimal = 1

		} else if slices.Contains(hexNums, v) {
			return 0, 0, fmt.Errorf("cannot use %c in a decimal number", v)
		} else {
			break
		}
		i++
	}

	return n, i, nil
}

func tokenizeNumberOct(str string, i int) (float64, int, error) {
	n := 0

	for i < len(str) {
		v := str[i]
		if '0' <= v && v <= '7' {
			n = n<<3 + int(v-'0')
		} else if v == '.' {
			return 0, 0, fmt.Errorf("cannot use decimal in an octal number")
		} else if slices.Contains(hexNums, v) {
			return 0, 0, fmt.Errorf("cannot use %c in an octal number", v)
		} else {
			break
		}
		i++
	}

	return float64(n), i, nil

}
func tokenizeNumberHex(str string, i int) (float64, int, error) {
	n := 0

	for i < len(str) {
		v := str[i]
		if '0' <= v && v <= '9' {
			n = n<<4 + int(v-'0')
		} else if 'A' <= v && v <= 'F' {
			n = n<<4 + int(v-'A') + 10
		} else if 'a' <= v && v <= 'f' {
			n = n<<4 + int(v-'a') + 10
		} else if v == '.' {
			return 0, 0, fmt.Errorf("cannot use decimal in an hexadecimal number")
		} else {
			break
		}
		i++
	}

	return float64(n), i, nil

}

func tokenizeNumberBin(str string, i int) (float64, int, error) {
	n := 0

	for i < len(str) {
		v := str[i]
		if v == '0' {
			n = n << 1
		} else if v == '1' {
			n = n<<1 + 1
		} else if v == '.' {
			return 0, 0, fmt.Errorf("cannot use decimal in a binary number")
		} else if slices.Contains(hexNums, v) {
			return 0, 0, fmt.Errorf("cannot use %c in a binary number", v)
		} else {
			break
		}
		i++
	}

	return float64(n), i, nil

}

func tokenizeNumber(str string, i int) (Token, int, error) {
	var err error
	n := 0.0

	negate := false
	if slices.Contains(unaryOperators, str[i]) {
		negate = str[i] == '-'
		i++

		if i >= len(str) {
			return nil, 0, fmt.Errorf("got '-' at the end of the string")
		}
	}

	if str[i] == '0' && len(str) > i+1 {
		i++
		switch str[i] {
		case 'x':
			// hex
			n, i, err = tokenizeNumberHex(str, i+1)
		case 'b':
			// bin
			n, i, err = tokenizeNumberDec(str, i+1)
		case 'o':
			// octal
			n, i, err = tokenizeNumberOct(str, i+1)
		default:
			// decimal
			n, i, err = tokenizeNumberDec(str, i+1)
		}
	} else {
		n, i, err = tokenizeNumberDec(str, i)
	}

	if err != nil {
		return nil, 0, err
	}

	if !negate {
		return Number(n), i, nil
	}

	n = -n

	return Number(n), i, nil
}

func tokenizeFunction(str string, i int) (Token, int, error) {
	function := strings.Builder{}

loop:
	for i < len(str) {
		v := str[i]
		switch {
		case slices.Contains(alpha, v) || slices.Contains(decNums, v):
			function.WriteRune(rune(v))
		default:
			break loop
		}
		i++
	}

	return Function(function.String()), i, nil
}

func tokenizeOperator(str string, i int, prev Token) (Token, int, error) {
	v := str[i]

	if prev == nil {
		return nil, i, nil
	} else if _, ok := prev.(LParen); ok {
		return nil, i, nil
	} else if _, ok := prev.(Binary); ok {
		return nil, i, nil
	}

	switch {
	case slices.Contains(operators, v):
		return Binary(str[i]), i + 1, nil
	}

	return nil, 0, fmt.Errorf("invalid operator: %c", str[i])
}

func tokenize(str string, i int, prev Token) (Token, int, error) {
	v := str[i]

	switch {
	case slices.Contains(operators, v):
		tok, i, err := tokenizeOperator(str, i, prev)

		if err != nil {
			return nil, i, fmt.Errorf("failed to tokenize operator: %w", err)
		}

		if tok != nil {
			return tok, i, nil
		}

		fallthrough
	case slices.Contains(decNums, v) || slices.Contains(unaryOperators, v):
		tok, i, err := tokenizeNumber(str, i)

		if err != nil {
			return nil, i, fmt.Errorf("failed to tokenize number: %w", err)
		}

		return tok, i, nil
	case slices.Contains(alpha, v):
		tok, i, err := tokenizeFunction(str, i)

		if err != nil {
			return nil, i, fmt.Errorf("failed to tokenize operator: %w", err)
		}

		return tok, i, nil
	case v == '(':
		return LParen{}, i + 1, nil
	case v == ')':
		return RParen{}, i + 1, nil
	case unicode.IsSpace(rune(v)):
		return nil, i + 1, nil
	default:
		return nil, 0, fmt.Errorf("invalid character")
	}
}

func Tokenize(str string) ([]Token, error) {
	tokens := make([]Token, 0)
	var tok Token
	var err error

	for i := 0; i < len(str); {
		prev := Token(nil)
		if len(tokens) > 0 {
			prev = tokens[len(tokens)-1]
		}
		tok, i, err = tokenize(str, i, prev)

		if err != nil {
			return nil, fmt.Errorf("failed to tokenize at index %d: %w", i, err)
		}

		if tok != nil {
			tokens = append(tokens, tok)
		}
	}

	if len(tokens) == 0 {
		return nil, fmt.Errorf("syntax error")
	}

	return tokens, nil
}
