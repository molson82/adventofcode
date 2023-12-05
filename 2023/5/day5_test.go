package day5_test

import (
	"testing"

	day5 "github.com/molson82/adventofcode/2023/5"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day5.Part1(sample)
	assert.Equal(t, ans, 35)
}
