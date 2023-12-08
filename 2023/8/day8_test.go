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
