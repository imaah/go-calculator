package preprocessor

import (
	"errors"
	"fmt"
	"strings"
)

func ProcessParenthesis(str string, m map[uint]string) (uint, error) {
	var stack = make([]int, 0)
	var num = uint(1)
	var i = 0

	for i < len(str) {
		var c = str[i]

		if c == '(' {
			stack = append(stack, i)
		}
		if c == ')' {
			if len(stack) > 0 {
				var start = stack[len(stack)-1]
				var end = i
				stack = stack[0 : len(stack)-1]

				var sub = str[start : end+1]
				var prelen = len(str)
				str = strings.Replace(str, sub, fmt.Sprintf("{%d}", num), 1)

				i -= prelen - len(str)

				m[num] = sub[1 : len(sub)-1]
				num += 1
			} else {
				return 0, errors.New("ParseError: missing left parenthesis")
			}
		}

		i++
	}

	m[0] = str

	if len(stack) != 0 {
		return 0, errors.New("ParseError: missing right parenthesis")
	}

	return num, nil
}

func ProcessToken(str string, token rune, cIdx uint, minIdx uint, m map[uint]string) (uint, error) {
	var i = 0
	var num = minIdx

	for i < len(str) {
		var c = str[i]

		if rune(c) == token {
			var leftSub = str[:i]
			var rightSub = str[i+1:]
			var prelen = len(str)

			m[num] = leftSub
			str = strings.Replace(str, leftSub, fmt.Sprintf("{%d}", num), 1)
			num += 1

			m[num] = rightSub
			str = strings.Replace(str, rightSub, fmt.Sprintf("{%d}", num), 1)
			num += 1

			i -= max(prelen-len(str)-1, 0)
		}
		i++
	}

	m[cIdx] = str

	return num, nil
}

func max(a, b int) int {
	if a < b {
		return a
	}
	return b
}
