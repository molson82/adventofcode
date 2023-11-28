package day6

import (
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

func Solution(input []string, numOfDays int) int {
	fish := utils.StrList_to_IntList(strings.Split(input[0], ","))

	fishMap := make(map[int]int, 8)
	for _, v := range fish {
		fishMap[v]++
	}

	for i := 1; i <= numOfDays; i++ {
		tempMap := make(map[int]int)
		for k, v := range fishMap {
			tempMap[k] = v
		}

		for i := 0; i <= 6; i++ {
			fishMap[i] = tempMap[i+1]
		}
		fishMap[6] = tempMap[7] + tempMap[0]
		fishMap[7] = tempMap[8]
		fishMap[8] = tempMap[0]
	}
	var totalFish int

	for _, v := range fishMap {
		totalFish += v
	}

	return totalFish
}
