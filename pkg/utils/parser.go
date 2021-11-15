package utils

import (
	"emorisse.fr/go-calculator/pkg/parser/types"
	"fmt"
	"regexp"
)

var SubGroupRegex = regexp.MustCompile(`^ *(?::\d+|\( *:\d+ *\)) *$`)

func ExtractIndex(key string) int {
	var index int
	_, err := fmt.Sscanf(key, ":%d", &index)

	if err != nil {
		_, err = fmt.Sscanf(key, "(:%d)", &index)

		if err != nil {
			return 0
		}
	}

	return index
}

func CleanMap(groups types.GroupMap) (types.GroupMap, int) {
	var max = 0

	for key, val := range groups {
		var tmpMax = max
		if id := ExtractIndex(key); id > max {
			tmpMax = id
		}
		if SubGroupRegex.MatchString(val) {
			tmpMax = max
			delete(groups, key)

			var redirectTo = fmt.Sprintf(":%d", ExtractIndex(val))
			groups = removeInOther(key, redirectTo, groups)
		}

		max = tmpMax
	}
	return groups, max
}

func removeInOther(remove, redirectTo string, groups types.GroupMap) types.GroupMap {
	var regexStr = fmt.Sprintf(" *%s *", remove)
	var regexStr2 = fmt.Sprintf(" *%s[0-9]+ *", remove)
	//TODO : Find a better fix.
	var regex1 = regexp.MustCompile(regexStr)
	var regex2 = regexp.MustCompile(regexStr2)

	for key, val := range groups {
		if regex1.MatchString(val) && !regex2.MatchString(val) {
			groups[key] = regex1.ReplaceAllString(val, redirectTo)
		}
	}
	return groups
}
