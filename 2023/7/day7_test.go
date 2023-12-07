package day7_test

import (
	"testing"

	day7 "github.com/molson82/adventofcode/2023/7"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day7.Part1(sample)
	assert.Equal(t, ans, 6440)
}
