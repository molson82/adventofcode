package day1

import (
	"github.com/molson82/adventofcode/pkg/utils"
)

func Part1(depths []string) int {
	var ans int
	deps := utils.StrList_to_IntList(depths)

	for i := 1; i < len(deps); i++ {
		if deps[i] > deps[i-1] {
			ans++
		}
	}

	return ans
}

func Part2(depths []string) int {
	ans := -1

	deps := utils.StrList_to_IntList(depths)

	var prev int
	for i := 0; i < len(deps)-2; i++ {
		slidingSum := deps[i] + deps[i+1] + deps[i+2]
		if slidingSum > prev {
			ans++
		}
		prev = slidingSum
	}

	return ans
}
