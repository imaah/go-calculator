package parser

import (
	"emorisse.fr/calcul/operators"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var regexInnerParenthesis *regexp.Regexp

func init() {
	regexInnerParenthesis, _ = regexp.Compile("[a-z]*\\([0-9.a-z+\\-*/ :]+\\)")
}

func Parse(str string) (operators.Operation, error) {
	var group []byte
	fmt.Println(str)
	var byteStr = makeCalculusPriority([]byte(str))
	str = string(byteStr)
	fmt.Println(string(byteStr))

	var groups = make(map[string]string)
	group = regexInnerParenthesis.Find(byteStr)

	for i := 0; group != nil; {
		i++
		var key = fmt.Sprintf(":%d", i)
		var part = string(group)
		groups[key] = part
		str = strings.Replace(str, part, fmt.Sprintf(":%d", i), 1)
		byteStr = []byte(str)
		group = regexInnerParenthesis.Find(byteStr)
	}

	groups = cleanMap(groups)
	var lastElem = len(groups)
	return buildOperator(groups, lastElem)
}

func buildOperator(groups map[string]string, index int) (operators.Operation, error) {
	var key = fmt.Sprintf(":%d", index)
	var elem = groups[key]
	fmt.Println(elem)

	//TODO : check if it is a binary expression
	//TODO : check if it is a unary expression
	//TODO : check if it is a function expression
	//TODO : check if it is a number expression

	//TODO : find the other operations inside the current expression then build it !
	//TODO : build the operation according to the expression found.

	return nil, errors.New("UnimplementedMethod")
}

func cleanMap(groups map[string]string) map[string]string {
	var regex, _ = regexp.Compile("\\(:\\d+\\)")
	for key, val := range groups {
		if regex.Match([]byte(val)) {
			delete(groups, key)
		}
	}
	return groups
}

func makeCalculusPriority(bytes []byte) []byte {
	bytes = makeSubGroups(bytes, []string{"*", "/", "^"})
	bytes = makeSubGroups(bytes, []string{"+", "-"})
	return bytes
}

func makeSubGroups(bytes []byte, operators []string) []byte {
	var count = 0
	var baseRegex = "(?:[0-9.]+|\\(.*\\)) *[%s] (?:[0-9.]+|\\(.*\\))"
	var regexStr = fmt.Sprintf(baseRegex, strings.Join(operators, ""))
	var regex, _ = regexp.Compile(regexStr)

	var str = string(bytes)

	for _, ope := range operators {
		count += strings.Count(str, ope)
	}

	for i := 0; i < count; i++ {
		bytes = regex.ReplaceAllFunc(bytes, addParenthesisAround)
	}

	return bytes
}

func addParenthesisAround(group []byte) []byte {
	return []byte("(" + string(group) + ")")
}
