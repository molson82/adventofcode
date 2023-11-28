package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

func Part1(input []string) int {
	x := 1
	cycle := 1
	queue := []int{}
	for _, v := range input {
		if v == "noop" {
			queue = append(queue, 0)
			continue
		}
		// get number from input
		n := strings.Split(v, "addx ")[1]
		num, err := strconv.Atoi(n)
		utils.CheckErr(err)
		// add number to queue
		queue = append(queue, 0)
		queue = append(queue, num)
	}
	for len(queue) != 0 {
		fmt.Printf("start cycle: %d | x: %d\n", cycle, x)
		var t int
		t, queue = queue[0], queue[1:]
		x = x + t
		cycle++
		fmt.Printf("end cycle: %d | x: %d\n", cycle, x)
	}
	return x
}

func Part2(input []string) int {
	return 0
}
