package day3

import (
	"fmt"
	"strconv"
)

func Part1(input []string) int {
	var ans int

	for _, bank := range input {
		ans += maxNum(2, bank)
	}

	fmt.Println("2025/3 part1 ans: ", ans)

	return ans
}

func Part2(input []string) int {
	var ans int

	for _, bank := range input {
		ans += maxNum(12, bank)
	}

	fmt.Println("2025/3 part2 ans: ", ans)

	return ans
}

func maxNum(size int, b string) int {
	result := ""
	startPos := 0

	for i := 0; i < size; i++ {
		endPos := len(b) - (size - i - 1)

		maxDigit := byte('0')
		maxPos := startPos

		for j := startPos; j < endPos; j++ {
			if b[j] > maxDigit {
				maxDigit = b[j]
				maxPos = j
			}
		}

		result += string(maxDigit)
		startPos = maxPos + 1
	}

	num, _ := strconv.Atoi(result)
	return num
}
