package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2021/day3"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("Part 2: %v", day3.Part2(utils.ReadInput("inputs/day3_input.txt")))
}
