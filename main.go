package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2022/day9"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("Answer: %v", day9.Part1(utils.ReadInput("internal/2022/day9/input.txt")))
}
