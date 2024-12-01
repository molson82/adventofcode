package day1

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	var ans int
	// code here...
	lists := inputLists(input)
	l1 := lists[0]
	l2 := lists[1]
	sort.Ints(l1)
	sort.Ints(l2)

	// calc distance
	for i := 0; i < len(l1); i++ {
		ans += int(math.Abs(float64(l1[i]) - float64(l2[i])))
	}

	fmt.Println("2024/1 part1 ans: ", ans)
	return ans
}

func Part2(input []string) int {
	var ans int
	// code here...
	lists := inputLists(input)
	l1 := lists[0]
	l2 := lists[1]

	for _, v := range l1 {
		occ := countOccurrences(l2, v)
		temp := v * occ
		ans += temp
	}

	fmt.Println("2024/1 part2 ans: ", ans)
	return ans
}

func countOccurrences(ns []int, num int) int {
	count := 0
	for _, value := range ns {
		if value == num {
			count++
		}
	}
	return count
}

func inputLists(input []string) [][]int {
	var l1 []int
	var l2 []int
	var result [][]int

	for _, line := range input {
		numbers := strings.Fields(line)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		l1 = append(l1, num1)
		l2 = append(l2, num2)
	}
	result = append(result, l1)
	result = append(result, l2)

	return result
}
