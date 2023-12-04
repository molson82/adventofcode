package day4_test

import (
	"testing"

	day4 "github.com/molson82/adventofcode/2023/4"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day4.Part1(sample)
	assert.Equal(t, ans, 13)
}

func TestPart1(t *testing.T) {
	input := utils.ReadInput("./input.txt")
	ans := day4.Part1(input)
	assert.Equal(t, ans, 24706)
}

func TestParseCard(t *testing.T) {
	type input struct {
		cardString string
		card       day4.Card
	}

	tests := []input{
		{
			cardString: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			card:       day4.Card{Num: 1, WinningNumbers: []int{41, 48, 83, 86, 17}, YourNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53}},
		},
		{
			cardString: "Card  10: 82 32 48 60 17 85 97 22 26 87 | 33 49 81 29 70  8 74 45 97 68 36 78 83 11 14 90 93 99  1 59 95 30  4 26  2",
			card:       day4.Card{Num: 10, WinningNumbers: []int{82, 32, 48, 60, 17, 85, 97, 22, 26, 87}, YourNumbers: []int{33, 49, 81, 29, 70, 8, 74, 45, 97, 68, 36, 78, 83, 11, 14, 90, 93, 99, 1, 59, 95, 30, 4, 26, 2}},
		},
	}

	for _, q := range tests {
		assert.Equal(t, day4.ParseCard(q.cardString), q.card)
	}
}
