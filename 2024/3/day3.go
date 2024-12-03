package day3

import (
	"fmt"
	"regexp"
	"sort"
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

var enabled = true

func Part2(input []string) int {
	var ans int
	for _, l := range input {
		nums := extractNumbersPart2(l)
		for _, n := range nums {
			mul := n[0] * n[1]
			ans += mul
		}
	}
	fmt.Println("2024/3 part2 ans: ", ans)

	return ans
}

func extractNumbersPart2(input string) [][]int {
	// Define the regex patterns
	mulPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`
	mulRe := regexp.MustCompile(mulPattern)
	doRe := regexp.MustCompile(doPattern)
	dontRe := regexp.MustCompile(dontPattern)
	// Find all matches for mul, do, and don't
	mulMatches := mulRe.FindAllStringSubmatchIndex(input, -1)
	doMatches := doRe.FindAllStringIndex(input, -1)
	dontMatches := dontRe.FindAllStringIndex(input, -1)

	// Combine all matches into a single slice and sort by position
	type match struct {
		start, end int
		kind       string
		submatches []int
	}

	var allMatches []match
	for _, m := range mulMatches {
		allMatches = append(allMatches, match{m[0], m[1], "mul", m[2:]})
	}
	for _, m := range doMatches {
		allMatches = append(allMatches, match{m[0], m[1], "do", nil})
	}
	for _, m := range dontMatches {
		allMatches = append(allMatches, match{m[0], m[1], "dont", nil})
	}

	// Sort matches by start position
	sort.Slice(allMatches, func(i, j int) bool {
		return allMatches[i].start < allMatches[j].start
	})

	// Process matches and keep track of the current state
	var results [][]int
	for _, m := range allMatches {
		switch m.kind {
		case "do":
			enabled = true
		case "dont":
			enabled = false
		case "mul":
			if enabled {
				num1, err1 := strconv.Atoi(input[m.submatches[0]:m.submatches[1]])
				num2, err2 := strconv.Atoi(input[m.submatches[2]:m.submatches[3]])
				if err1 == nil && err2 == nil {
					results = append(results, []int{num1, num2})
				}
			}
		}
	}
	return results
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
