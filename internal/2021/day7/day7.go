package day7

import (
	"math"
	"sort"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

func findMedian(arr []int) int {
	index := len(arr) / 2
	if len(arr)%2 != 0 {
		return arr[index]
	}
	return (arr[index-1] + arr[index]) / 2
}

func Part1(input []string) int {
	crabs := utils.StrList_to_IntList(strings.Split(input[0], ","))
	sort.Ints(crabs)

	var fuel int
	median := findMedian(crabs)
	for _, v := range crabs {
		fuel += int(math.Abs(float64(v) - float64(median)))
	}

	return fuel
}
