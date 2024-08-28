package helper

import "strings"

func IsContains(contain string,err error )bool{
	return strings.Contains(contain,strings.ToLower(err.Error()))
}