package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func Part1(input []string) int {
	var ans int

	for _, l := range input {
		nums := extractNumbersPart1(l)
		for _, n := range nums {
			mul := n[0] * n[1]
			ans += mul
		}
	}

	fmt.Println("2024/3 part1 ans: ", ans)

	return ans
}

func Part2(input []string) int {
	var ans int
	// code here...
	fmt.Println("2024/3 part2 ans: ", ans)

	return ans
}

func extractNumbersPart1(input string) [][]int {
	// Define the regex pattern
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)
	// Extract the numbers from each match
	var results [][]int
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		results = append(results, []int{num1, num2})
	}

	return results
}
