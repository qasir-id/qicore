package util

import (
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

var IntegerToBool = map[int]bool{
	0: false,
	1: true,
}

//convert integer to bool
func ItoB(input int) bool {
	return IntegerToBool[input]
}

//convert bool to integer
func BtoI(input bool) int {
	if input {
		return 1
	}
	return 0
}

//string to array string
func Explode(s string, separator string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, separator)
}

//covert array string to array int64
func ExplodeInt64(s string, separator string) []int64 {
	var integers []int64
	for _, v := range Explode(s, separator) {
		val, _ := strconv.Atoi(v)
		integers = append(integers, int64(val))
	}
	return integers
}

//covert array string to array int32
func ExplodeInt32(s string, separator string) []int32 {
	var integers []int32
	for _, v := range Explode(s, separator) {
		val, _ := strconv.Atoi(v)
		integers = append(integers, int32(val))
	}
	return integers
}
