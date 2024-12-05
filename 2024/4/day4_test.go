package day4_test

import (
	"testing"

	day4 "github.com/molson82/adventofcode/2024/4"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day4.Part1(sample)
	assert.Equal(t, 18, ans)
}

func TestPart1(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	ans := day4.Part1(input)
	assert.NotEqual(t, 2541, ans, "too low")
	assert.NotEqual(t, 2547, ans, "too high")
	assert.Equal(t, 2545, ans)
}

func TestPart2Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day4.Part2(sample)
	assert.Equal(t, 9, ans)
}

// func TestPart2(t *testing.T) {
// 	input := utils.ReadInput("./input.txt")
// 	ans := day4.Part2(input)
// 	assert.Equal(t, 0, ans)
// }
