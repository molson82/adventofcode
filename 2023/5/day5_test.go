package day5_test

import (
	"log"
	"os"
	"testing"

	day5 "github.com/molson82/adventofcode/2023/5"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}
	defer file.Close()

	ans, err := day5.PartOne(file, os.Stdout)
	if err != nil {
		log.Fatalf("could not solve: %v", err)
	}
	assert.Equal(t, ans, 35)
}

func TestPart1(t *testing.T) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}
	defer file.Close()

	ans, err := day5.PartOne(file, os.Stdout)
	if err != nil {
		log.Fatalf("could not solve: %v", err)
	}
	assert.Equal(t, ans, 322500873)
}

func TestPart2(t *testing.T) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}
	defer file.Close()

	ans, err := day5.PartTwo(file, os.Stdout)
	if err != nil {
		log.Fatalf("could not solve: %v", err)
	}
	assert.Equal(t, ans, 108956227)
}
