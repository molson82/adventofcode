package main

import (
	"log"

	"github.com/molson82/adventofcode/internal/2022/day7"
	"github.com/molson82/adventofcode/pkg/utils"
)

func main() {
	log.Println("Advent of Code")

	log.Printf("Answer: %v", day7.Part1(utils.ReadInput("internal/2022/day7/input.txt")))
}
