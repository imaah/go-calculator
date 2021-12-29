package utils

import (
	"fmt"
	"regexp"
)

var OnlySubGroupRegex = regexp.MustCompile(`^ *{[0-9]+} *$`)
var SubGroupRegex = regexp.MustCompile(` *{[0-9]+} *`)

func IsSubGroup(str string) bool {
	return OnlySubGroupRegex.MatchString(str)
}

func GetSubGroupValue(str string) (uint, error) {
	var val uint
	var n, err = fmt.Sscanf(str, "{%d}", &val)
	if n < 1 {
		return 0, err
	}
	return val, nil
}
