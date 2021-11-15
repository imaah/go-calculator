package parser

import (
	"emorisse.fr/go-calculator/pkg/operation"
	"emorisse.fr/go-calculator/pkg/parser/builder"
	"emorisse.fr/go-calculator/pkg/parser/preprocessor"
	"emorisse.fr/go-calculator/pkg/parser/types"
	"emorisse.fr/go-calculator/pkg/utils"
	"fmt"
)

func Parse(str string) (operation.Operation, error) {
	var groups = make(types.GroupMap)
	var lastIndex = 1

	str, lastIndex, groups = preprocessor.ProcessFunctions(str, lastIndex, groups)
	str, lastIndex, groups = preprocessor.ProcessParenthesis(str, lastIndex, groups)

	for k, group := range groups {
		groups[k], lastIndex, groups = preprocessor.ProcessBinary(group, lastIndex, groups)
	}
	str, lastIndex, groups = preprocessor.ProcessBinary(str, lastIndex, groups)

	for k, group := range groups {
		groups[k], lastIndex, groups = preprocessor.ProcessUnary(group, lastIndex, groups)
	}
	str, lastIndex, groups = preprocessor.ProcessUnary(str, lastIndex, groups)

	var key = fmt.Sprintf(":%d", lastIndex)
	groups[key] = str

	groups, lastIndex = utils.CleanMap(groups)
	return builder.BuildOperator(groups, lastIndex)
}
