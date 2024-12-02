package day2_test

import (
	"testing"

	day2 "github.com/molson82/adventofcode/2024/2"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day2.Part1(sample)
	assert.Equal(t, 2, ans)
}

// func TestPart1(t *testing.T) {
// 	input := utils.ReadInput("./input.txt")
// 	ans := day2.Part1(input)
// 	assert.Equal(t, 0, ans)
// }
//
// func TestPart2Sample(t *testing.T) {
// 	sample := utils.ReadInput("./sample.txt")
// 	ans := day2.Part2(sample)
// 	assert.Equal(t, 0, ans)
// }
//
// func TestPart2(t *testing.T) {
// 	input := utils.ReadInput("./input.txt")
// 	ans := day2.Part2(input)
// 	assert.Equal(t, 0, ans)
// }
