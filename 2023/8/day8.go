package day8

import (
	"fmt"
	"strings"
)

func Part1(lines []string) int {
	var ans int
	seq := strings.Split(lines[0], "")

	network := make(map[string][]string)
	for i := 2; i < len(lines); i++ {
		l := lines[i]
		spl := strings.Split(l, "= ")
		key := strings.Fields(spl[0])[0]
		valSpl := strings.Split(spl[1], ", ")
		leftVal := valSpl[0][1:]
		rightVal := valSpl[1][:len(valSpl[1])-1]

		network[key] = []string{leftVal, rightVal}
	}

	// go over seq
	stepKey := "AAA"
	for i := 0; i < len(seq); i++ {
		ans++
		if stepKey == "ZZZ" {
			break
		}
		s := seq[i]
		currNode := network[stepKey]
		if s == "R" {
			stepKey = currNode[1]
		} else {
			stepKey = currNode[0]
		}
		// reset loop
		if i == len(seq)-1 {
			i = -1
		}
	}
	fmt.Println("2023/8 part1 ans: ", ans-1)
	return ans - 1
}

func Part2(lines []string) int {
	var ans int
	seq := strings.Split(lines[0], "")
	fmt.Println("seq: ", seq)

	network := make(map[string][]string)
	for i := 2; i < len(lines); i++ {
		l := lines[i]
		spl := strings.Split(l, "= ")
		key := strings.Fields(spl[0])[0]
		valSpl := strings.Split(spl[1], ", ")
		leftVal := valSpl[0][1:]
		rightVal := valSpl[1][:len(valSpl[1])-1]

		network[key] = []string{leftVal, rightVal}
	}

	fmt.Println("2023/8 part2 ans: ", ans)
	return ans
}
