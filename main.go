package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2022/day1"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("2022 | Day1 Answer: %v", day1.Part2(utils.ReadInput("internal/2022/day1/input.txt")))
}
