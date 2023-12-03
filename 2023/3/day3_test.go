package day3_test

import (
	"testing"

	day3 "github.com/molson82/adventofcode/2023/3"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day3.Part1(sample)
	assert.Equal(t, ans, 4361)
}

func TestPart1(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	ans := day3.Part1(input)
	assert.NotEqual(t, ans, 336471, "too low")
	assert.Equal(t, ans, 549908)
}

func TestPart2Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day3.Part2(sample)
	assert.Equal(t, ans, 467835)
}

func TestPart2(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	ans := day3.Part2(input)
	assert.NotEqual(t, ans, 31488236, "too low")
	assert.NotEqual(t, ans, 3846746, "too low")
}

func TestGetNumber(t *testing.T) {
	type input struct {
		line  string
		index int
		num   int
	}
	inputs := []input{
		{line: "..35..633.", index: 2, num: 35},
		{line: "..35..633.", index: 3, num: 35},
		{line: "..35..633.", index: 6, num: 633},
		{line: "..35..633.", index: 7, num: 633},
		{line: "..35..633.", index: 8, num: 633},
		{line: "617*.....", index: 0, num: 617},
	}
	for _, i := range inputs {
		_, exp, err := day3.GetNumber(i.index, i.line)
		assert.Equal(t, err, nil)
		assert.Equal(t, exp, i.num)
	}
}
