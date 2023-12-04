package day4

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/molson82/adventofcode/pkg/utils"
)

type Card struct {
	Num            int
	WinningNumbers []int
	YourNumbers    []int
	Points         int
}

func Part1(lines []string) int {
	var ans int
	var pile []Card
	for _, l := range lines {
		newCard := ParseCard(l)
		newCard.calcScore()
		pile = append(pile, newCard)
	}

	for _, c := range pile {
		ans += c.Points
	}

	fmt.Println("2023/4 part1 ans: ", ans)
	return ans
}

func Part2() {
	var ans int
	// code here...
	fmt.Println("2023/4 part2 ans: ", ans)
}

func (c *Card) calcScore() {
	var score int
	for _, wn := range c.WinningNumbers {
		for _, yn := range c.YourNumbers {
			if wn == yn {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}
	c.Points = score
}

func ParseCard(input string) Card {
	re := regexp.MustCompile(`Card\s+(\d+): (.+?) \| (.+)`)
	matches := re.FindStringSubmatch(input)

	cardNum, err := strconv.Atoi(matches[1])
	utils.CheckErr(err)

	winningNumbers := utils.String_to_IntList(matches[2])
	yourNumbers := utils.String_to_IntList(matches[3])

	return Card{
		Num:            cardNum,
		WinningNumbers: winningNumbers,
		YourNumbers:    yourNumbers,
		Points:         0,
	}
}
