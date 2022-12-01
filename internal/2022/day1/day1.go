package day1

import (
	"sort"

	"github.com/molson82/adventofcode/pkg/utils"
)

func Part1(input []string) int {
	var max int
	var temp int
	for _, v := range input {
		if v == "" {
			if temp > max {
				max = temp
			}
			temp = 0
		} else {
			intValue := utils.Str_to_Int(v)
			temp += intValue
		}
	}
	return max
}

func Part2(input []string) int {
	sumList := []int{}
	var temp int
	for _, v := range input {
		if v == "" {
			sumList = append(sumList, temp)
			temp = 0
			continue
		}
		intValue := utils.Str_to_Int(v)
		temp += intValue
	}
	sumList = append(sumList, temp)

	sort.Ints(sumList)
	third := sumList[len(sumList)-3]
	second := sumList[len(sumList)-2]
	first := sumList[len(sumList)-1]
	return third + second + first
}
