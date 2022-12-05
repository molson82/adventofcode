package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2022/day3"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("Answer: %v", day3.Part2(utils.ReadInput("internal/2022/day3/input.txt")))
}
