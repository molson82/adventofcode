package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2021/day5"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("Answer: %v", day5.Part2(utils.ReadInput("inputs/day5_input.txt")))
}
