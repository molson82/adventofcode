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

func computePart2(num int, arr []int) int {
	var res int
	for _, v := range arr {
		res += int(math.Abs(float64(v)-float64(num)) * (math.Abs(float64(v)-float64(num)) + 1) / 2)
	}

	return res
}

func Part2(input []string) int {
	crabs := utils.StrList_to_IntList(strings.Split(input[0], ","))
	sort.Ints(crabs)

	minFuel := 99999999
	min := crabs[0]
	max := crabs[len(crabs)-1]

	for i := min; i < max; i++ {
		val := computePart2(i, crabs)
		if val < minFuel {
			minFuel = val
		}
	}

	return minFuel
}
