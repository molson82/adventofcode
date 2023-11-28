package utils

import "strconv"

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
