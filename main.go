package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2021/day7"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("Answer: %v", day7.Part2(utils.ReadInput("inputs/day7_input.txt")))
}
