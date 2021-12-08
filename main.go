package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2021/day8"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("Answer: %v", day8.Part1(utils.ReadInput("inputs/day8_input.txt")))
}
