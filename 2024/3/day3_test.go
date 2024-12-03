package day3_test

import (
	"testing"

	day3 "github.com/molson82/adventofcode/2024/3"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day3.Part1(sample)
	assert.NotEqual(t, 0, ans)
}

// func TestPart1(t *testing.T) {
// 	input := utils.ReadInput("./input.txt")
// 	ans := day3.Part1(input)
// 	assert.NotEqual(t, 0, ans)
// }
//
// func TestPart2Sample(t *testing.T) {
// 	sample := utils.ReadInput("./sample.txt")
// 	ans := day3.Part2(sample)
// 	assert.NotEqual(t, 0, ans)
// }
//
// func TestPart2(t *testing.T) {
// 	input := utils.ReadInput("./input.txt")
// 	ans := day3.Part2(input)
// 	assert.NotEqual(t, 0, ans)
// }
