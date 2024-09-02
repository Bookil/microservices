package helper

import "strings"

func IsContains(contain string, err error) bool {
	toLoweredErr := strings.ToLower(err.Error())
	isContains := strings.Contains(toLoweredErr, contain)
	return isContains
}
