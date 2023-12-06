package day6_test

import (
	"testing"

	day6 "github.com/molson82/adventofcode/2023/6"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day6.Part1(sample)
	assert.Equal(t, ans, 288)
}
