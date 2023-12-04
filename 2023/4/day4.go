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

var part2Pile []Card
var stackedPile []Card

func Part2(lines []string) int {
	for _, l := range lines {
		newCard := ParseCard(l)
		newCard.CalcScore()
		part2Pile = append(part2Pile, newCard)
		stackedPile = append(stackedPile, newCard)
	}
	ProcessStack(part2Pile)
	PrintOrder(stackedPile)
	fmt.Println("\n2023/4 part2 ans: ", len(stackedPile))
	return len(stackedPile)
}

func ProcessStack(cards []Card) {
	for i := 0; i < 2; i++ {
		c := stackedPile[i]
		fmt.Println("\ncard: ", c.Num)
		fmt.Printf("matches: %v\n", c.NumOfMatches)
		if c.NumOfMatches != 0 {
			cardSection := GetCardsFromStack(c.Num, c.Num+c.NumOfMatches-1, part2Pile)
			PrintOrder(cardSection)
			// find first index of first card from GetCardsFromStack
			var index int
			for j, q := range stackedPile {
				if q.Num == cardSection[0].Num {
					index = j
				}
			}
			fmt.Println("index: ", index)
			stackedPile = insertCards(stackedPile, cardSection, index+1)
			fmt.Println()
			PrintOrder(stackedPile)
		}
	}
}

func insertCards(destination []Card, cardsToInsert []Card, startIndex int) []Card {
	if startIndex < 0 || startIndex > len(destination) {
		return nil
	}

	// Create a new slice to hold the combined result
	result := make([]Card, len(destination)+len(cardsToInsert))

	// Copy elements from the destination slice up to the start index
	copy(result, destination[:startIndex])

	// Insert the new cards
	copy(result[startIndex:], cardsToInsert)

	// Copy the remaining elements from the original slice
	copy(result[startIndex+len(cardsToInsert):], destination[startIndex:])

	return result
}

func GetCardsFromStack(start, end int, pile []Card) []Card {
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
