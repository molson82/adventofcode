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
	NumOfMatches   int
}

func Part1(lines []string) int {
	var ans int
	var pile []Card
	for _, l := range lines {
		newCard := ParseCard(l)
		newCard.CalcScore()
		pile = append(pile, newCard)
	}

	for _, c := range pile {
		ans += c.Points
	}

	fmt.Println("2023/4 part1 ans: ", ans)
	return ans
}

var pile []Card

func Part2(lines []string) int {
	pile = []Card{}
	var ans int
	stack := make(map[int]int)

	for _, l := range lines {
		newCard := ParseCard(l)
		newCard.CalcScore()
		pile = append(pile, newCard)
		stack[newCard.Num] = 1
	}
	ProcessStack(stack)
	for _, v := range stack {
		ans += v
	}
	fmt.Println("\n2023/4 part2 ans: ", ans)
	return ans
}

func ProcessStack(stack map[int]int) {
	card := 1
	for card < len(stack) {
		s := stack[card]
		for i := 0; i < s; i++ {
			copies := GetCardsFromStack(card, card+pile[card-1].NumOfMatches-1, pile)
			// fmt.Printf("\nCard %v | num: %v NumOfMatches: %v | Copies: ", card, pile[card-1].Num, pile[card-1].NumOfMatches)
			// PrintOrder(copies)
			for _, c := range copies {
				stack[c.Num] = stack[c.Num] + 1
			}
		}
		card++
	}
}

func GetCardsFromStack(start, end int, pile []Card) []Card {
	// fmt.Println("len: ", len(pile))
	return pile[start : end+1]
}

func PrintOrder(cards []Card) {
	for _, c := range cards {
		fmt.Printf("%v, ", c.Num)
	}
}

func (c *Card) CalcScore() {
	var score int
	for _, wn := range c.WinningNumbers {
		for _, yn := range c.YourNumbers {
			if wn == yn {
				c.NumOfMatches++
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
