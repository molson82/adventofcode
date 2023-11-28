package day8

import "strings"

func Part1(input []string) int {
	var count int
	for _, v := range input {
		output := strings.Split(strings.Split(v, " | ")[1], " ")
		for _, j := range output {
			if len(j) == 2 || len(j) == 4 || len(j) == 3 || len(j) == 7 {
				count++
			}
		}

	}

	return count
}

func Part2(input []string) int {

	return 0
}
