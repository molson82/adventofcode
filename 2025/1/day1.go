package day1

import (
	"fmt"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

func Part1(input []string) int {
	position := 50
	totalPasses := 0

	for _, v := range input {
		direction := v[0]
		amount := utils.Str_to_Int(v[1:])
		if strings.ToUpper(string(direction)) == "L" {
			position = rotateLeft(position, amount)
		} else {
			position = rotateRight(position, amount)
		}
		// Count when we land exactly on 0
		if position == 0 {
			totalPasses++
		}
	}
	fmt.Println("2025/1 part1 ans: ", totalPasses)

	return totalPasses
}

func Part2(input []string) int {
	position := 50
	totalPasses := 0

	for _, v := range input {
		direction := v[0]
		amount := utils.Str_to_Int(v[1:])
		var passes int
		if strings.ToUpper(string(direction)) == "L" {
			position, passes = rotateLeftWithCount(position, amount)
		} else {
			position, passes = rotateRightWithCount(position, amount)
		}
		totalPasses += passes
	}
	fmt.Println("2025/1 part2 ans: ", totalPasses)

	return totalPasses
}

func rotateLeft(position int, amount int) int {
	newPos := (position - amount) % 100
	if newPos < 0 {
		newPos += 100
	}
	return newPos
}

func rotateRight(position int, amount int) int {
	return (position + amount) % 100
}

func rotateLeftWithCount(position int, amount int) (int, int) {
	newPos := rotateLeft(position, amount)

	var passes int
	if position == 0 {
		// Starting at 0, we only pass through it again after full rotations
		passes = amount / 100
	} else if amount >= position {
		// We pass 0 at steps: position, position + 100, position + 200, etc.
		passes = (amount-position)/100 + 1
	}

	return newPos, passes
}

func rotateRightWithCount(position int, amount int) (int, int) {
	newPos := rotateRight(position, amount)

	var passes int
	if position == 0 {
		// Starting at 0, we only pass through it again after full rotations
		passes = amount / 100
	} else {
		// We pass 0 at steps: (100 - position), (100 - position) + 100, etc.
		stepsToZero := 100 - position
		if amount >= stepsToZero {
			passes = (amount-stepsToZero)/100 + 1
		}
	}

	return newPos, passes
}
