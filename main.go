package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2021/day6"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("Answer: %v", day6.Part1(utils.ReadInput("inputs/day6_input.txt"), 80))
}
