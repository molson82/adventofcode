package utils

import (
	"strconv"
	"strings"
)

func Str_to_Int(str string) int {
	intVar, err := strconv.Atoi(str)
	CheckErr(err)
	return intVar
}

func StrList_to_IntList(strList []string) []int {
	var res []int
	for _, v := range strList {
		res = append(res, Str_to_Int(v))
	}
	return res
}

// Function to convert string to int slice
func String_to_IntList(s string) []int {
	strNumbers := strings.Fields(s)
	var numbers []int
	for _, str := range strNumbers {
		numbers = append(numbers, Str_to_Int(str))
	}
	return numbers
}
