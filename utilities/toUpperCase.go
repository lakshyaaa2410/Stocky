package utilities

import "strings"

func UpperCaseFirstLetter(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
