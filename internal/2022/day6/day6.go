package day6

import (
	"fmt"
	"strings"
)

func Part1(input string) int {
	var seg string
OUTER:
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			s := string(input[j])
			if strings.Contains(seg, s) {
				seg = ""
				continue OUTER
			}
			seg = fmt.Sprintf("%s%s", seg, s)
			if len(seg) == 4 {
				return j + 1
			}
		}
	}

	return 0
}

func Part2(input []string) int {
	return 0
}
