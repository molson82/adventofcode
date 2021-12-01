package day1

import "github.com/molson82/adventofcode/pkg/utils"

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
