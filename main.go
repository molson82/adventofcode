package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2021/day4"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("Part 2: %v", day4.Part2(utils.ReadInput("inputs/day4_input.txt")))
}
