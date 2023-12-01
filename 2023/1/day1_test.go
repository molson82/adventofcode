package day1_test

import (
	"testing"

	day1 "github.com/molson82/adventofcode/2023/1"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	expect := day1.Part1(sample)
	assert.Equal(t, expect, 142)
}

func TestPart1(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	expect := day1.Part1(input)
	assert.Equal(t, expect, 55172)
}

func TestPart2Sample(t *testing.T) {
	sample := utils.ReadInput("./samplePart2.txt")
	expect := day1.Part2(sample)
	assert.Equal(t, expect, 281)
}

func TestPart2(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	expect := day1.Part2(input)
	assert.NotEqual(t, expect, 54953)
	assert.Equal(t, expect, 54925)
}

func TestGetLineNumV2(t *testing.T) {
	var testInput = [][]string{
		{"pknhg9jql51one4", "94"},
		{"1threetwo6rfvsgjbhjlcfour", "14"},
		{"four7fiveseven3fourdqvsnkkdrbrhnfsx", "44"},
		{"7rcqcbqsfz766", "76"},
		{"4hdzzzggfzs18eight2twoqmnlcxzxs", "42"},
		{"x8", "88"},
		{"1two3fourxxx5xxxsixseven", "17"},
		{"xpkfldsggg3eightbqtbrfthd", "38"},
		{"eighthree", "83"},
		{"sevenine", "79"},
	}

	for _, v := range testInput {
		e := day1.GetLineNumV2(v[0])
		assert.Equal(t, e, v[1])
	}
}
