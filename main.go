package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2021/day1"
	"github.com/molson82/adventofcode/internal/2021/day2"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	// attempt 1: 1615 = too low
	// attempt 2: 1616 = correct
	log.Printf("2021 Day 1 Part 1: %v", day1.Part1(utils.ReadInput("inputs/day1_input.txt")))
	// attempt 1: 1645 = correct
	log.Printf("2021 Day 1 Part 2: %v", day1.Part2(utils.ReadInput("inputs/day1_input.txt")))

	// attempt 1: 2039912 = correct
	log.Printf("2021 Day 2 Part 1: %v", day2.Part1(utils.ReadInput("inputs/day2_input.txt")))
}
