package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

func Part1(input []string) int {
	var ans int
	// Input is just one long line. Split line up by "," to get actual list of ranges
	ranges := strings.Split(input[0], ",")

	for _, r := range ranges {
		t := strings.Split(r, "-")
		num1 := utils.Str_to_Int(t[0])
		num2 := utils.Str_to_Int(t[1])

		for i := num1; i <= num2; i++ {
			if !isValid(i) {
				ans += i
			}
		}
	}

	fmt.Println("2025/2 part1 ans: ", ans)
	return ans
}

func Part2(input []string) int {
	var ans int
	// Input is just one long line. Split line up by "," to get actual list of ranges
	ranges := strings.Split(input[0], ",")

	for _, r := range ranges {
		t := strings.Split(r, "-")
		num1 := utils.Str_to_Int(t[0])
		num2 := utils.Str_to_Int(t[1])

		for i := num1; i <= num2; i++ {
			if !isValidPart2(i) {
				ans += i
			}
		}
	}

	fmt.Println("2025/2 part2 ans: ", ans)

	return ans
}

func isValid(n int) bool {
	s := strconv.Itoa(n)

	if len(s)%2 != 0 {
		return true
	}

	mid := len(s) / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]

	return firstHalf != secondHalf
}

func isValidPart2(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	// Try all possible pattern lengths from 1 to length/2
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		// Length must be divisible by pattern length for it to repeat evenly
		if length%patternLen != 0 {
			continue
		}

		pattern := s[:patternLen]
		repetitions := length / patternLen

		// Build the repeated string and compare
		if strings.Repeat(pattern, repetitions) == s {
			return false
		}
	}

	return true
}
