package day2

import (
	"fmt"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

/*
*
Report is "safe" if BOTH are true:
- The levels are either all increasing or all decreasing.
- Any two adjacent levels differ by at least 1 and at most 3.
*
*/
func Part1(input []string) int {
	var ans int

	numInput := convertInput(input)
	for _, r := range numInput {
		isSafe := isReportSafe(r, 1, 3)
		if isSafe {
			ans++
		}
	}

	fmt.Println("2024/2 part1 ans: ", ans)

	return ans
}

func Part2(input []string) int {
	var ans int
	// code here...
	fmt.Println("2024/2 part2 ans: ", ans)

	return ans
}

func isReportSafe(report []int, min, max int) bool {
	return isIncreasing(report, min, max) || isDecreasing(report, min, max)
}

func isIncreasing(report []int, min, max int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < min || diff > max {
			return false
		}
	}
	return true
}

func isDecreasing(report []int, min, max int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if diff < min || diff > max {
			return false
		}
	}
	return true
}

func convertInput(input []string) [][]int {
	var result [][]int

	for _, l := range input {
		numbers := strings.Fields(l)
		numList := utils.StrList_to_IntList(numbers)
		result = append(result, numList)
	}

	return result
}
