package day8_test

import (
	"testing"

	day8 "github.com/molson82/adventofcode/2023/8"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day8.Part1(sample)
	assert.Equal(t, ans, 2)
}

func TestPart1Sample2(t *testing.T) {
	input := []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	ans := day8.Part1(input)
	assert.Equal(t, ans, 6)
}

func TestPart1(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	ans := day8.Part1(input)
	assert.Equal(t, ans, 12083)
}

func TestPart2Sample(t *testing.T) {
	input := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}
	ans := day8.Part2(input)
	assert.Equal(t, ans, 6)
}

func TestPart2(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	ans := day8.Part2(input)
	assert.NotEqual(t, ans, 103975, "too low")
}
