package day1_test

import (
	"testing"

	day1 "github.com/molson82/adventofcode/2025/1"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day1.Part1(sample)
	assert.NotEqual(t, 0, ans)
	assert.Equal(t, 3, ans)
}

func TestPart1(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	ans := day1.Part1(input)
	assert.Equal(t, 1180, ans)
}

func TestPart2Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day1.Part2(sample)
	assert.Equal(t, 6, ans)
}

func TestPart2(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	ans := day1.Part2(input)
	assert.Equal(t, 6892, ans)
}
