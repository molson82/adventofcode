package day2

import (
	"strings"
)

func Part1(input []string) int {
	// A | X = Rock
	// B | Y = Paper
	// C | Z = Scissors
	// shape: 1,2,3
	// outcome: 0,3,6

	var score int
	for _, v := range input {
		opp := strings.Split(v, " ")[0]
		you := strings.Split(v, " ")[1]

		var choice int
		switch you {
		case "X":
			choice = 1
		case "Y":
			choice = 2
		case "Z":
			choice = 3
		}
		score += play(opp, you) + choice
	}

	return score
}

func Part2(input []string) int {
	// X = lose
	// Y = draw
	// Z = win

	var score int
	for _, v := range input {
		opp := strings.Split(v, " ")[0]
		you := strings.Split(v, " ")[1]

		var choice int
		switch you {
		case "X":
			choice = 0
		case "Y":
			choice = 3
		case "Z":
			choice = 6
		}

		var move int
		switch opp {
		case "A":
			if choice == 0 {
				move = 3
			} else if choice == 3 {
				move = 1
			} else if choice == 6 {
				move = 2
			}
		case "B":
			if choice == 0 {
				move = 1
			} else if choice == 3 {
				move = 2
			} else if choice == 6 {
				move = 3
			}
		case "C":
			if choice == 0 {
				move = 2
			} else if choice == 3 {
				move = 3
			} else if choice == 6 {
				move = 1
			}
		}

		score += choice + move
	}

	return score
}

// 0 = you loose
// 3 = draw
// 6 = you win
func play(opp, you string) int {
	if opp == "A" && you == "Z" {
		return 0
	} else if opp == "A" && you == "X" {
		return 3
	} else if opp == "A" && you == "Y" {
		return 6
	} else if opp == "B" && you == "X" {
		return 0
	} else if opp == "B" && you == "Y" {
		return 3
	} else if opp == "B" && you == "Z" {
		return 6
	} else if opp == "C" && you == "Y" {
		return 0
	} else if opp == "C" && you == "Z" {
		return 3
	}
	return 6
}
