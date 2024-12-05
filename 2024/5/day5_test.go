package day5_test

import (
	"testing"

	day5 "github.com/molson82/adventofcode/2024/5"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day5.Part1(sample)
	assert.NotEqual(t, 0, ans)
}

// func TestPart1(t *testing.T) {
// 	input := utils.ReadInput("./input.txt")
// 	ans := day5.Part1(input)
// 	assert.Equal(t, 0, ans)
// }
//
// func TestPart2Sample(t *testing.T) {
// 	sample := utils.ReadInput("./sample.txt")
// 	ans := day5.Part2(sample)
// 	assert.Equal(t, 0, ans)
// }
//
// func TestPart2(t *testing.T) {
// 	input := utils.ReadInput("./input.txt")
// 	ans := day5.Part2(input)
// 	assert.Equal(t, 0, ans)
// }
