package day9_test

import (
	"testing"

	day9 "github.com/molson82/adventofcode/2023/9"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day9.Part1(sample)
	assert.Equal(t, ans, 114)
}

func TestPart1(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	ans := day9.Part1(input)
	assert.Equal(t, ans, 1834108701)
}

func TestPart2Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day9.Part2(sample)
	assert.Equal(t, ans, 2)
}

func TestCalcSeq(t *testing.T) {
	type testList struct {
		seq []int
		ans int
	}
	tests := []testList{
		{seq: []int{0, 3, 6, 9, 12, 15}, ans: 18},
		{seq: []int{1, 3, 6, 10, 15, 21}, ans: 28},
		{seq: []int{10, 13, 16, 21, 30, 45}, ans: 68},
	}

	for _, x := range tests {
		assert.Equal(t, day9.CalcSequence(x.seq, [][]int{x.seq}), x.ans)
	}
}
