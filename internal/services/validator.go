package services

import "regexp"

const (
	expRegExp = "^[\\(\\)\\+\\-\\*\\/\\s0-9]+$"
)

func IsValidExpression(expression string) bool {
	r := regexp.MustCompile(expRegExp)
	return r.MatchString(expression)
}
