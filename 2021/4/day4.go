package day4

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

type square struct {
	value  int
	marked bool
}

type board [][]*square

func (b board) print() {
	for _, v := range b {
		fmt.Printf("\n")
		for _, j := range v {
			fmt.Printf("%v ", j)
		}
	}
}

func (b board) drawNumber(num int) {
	for _, v := range b {
		for _, j := range v {
			if j.value == num {
				j.marked = true
			}
		}
	}
}

func checkRow(row []*square) bool {
	for _, v := range row {
		if !v.marked {
			return false
		}
	}
	return true
}

func transpose(a board) board {
	newArr := make([][]*square, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr
}

func (b board) checkBoard() bool {
	bCopy := b
	for _, v := range bCopy {
		if checkRow(v) {
			return true
		}
	}

	// transpose and check cols as rows
	bCopy = transpose(bCopy)
	for _, v := range bCopy {
		if checkRow(v) {
			return true
		}
	}

	return false
}

func (b board) calculateScore() int {
	var score int
	for _, v := range b {
		for _, j := range v {
			if !j.marked {
				score += j.value
			}
		}
	}

	return score
}

func createSquareRow(r []int) []*square {
	var res []*square
	for _, v := range r {
		s := square{value: v, marked: false}
		res = append(res, &s)
	}
	return res
}

func Part1(input []string) int {
	var boardList []board
	var b board
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			// end of board, add to board list
			boardList = append(boardList, b)
			b = board{}
		} else {
			// still creating the board
			whitespaces := regexp.MustCompile(`\s+`)
			str := whitespaces.ReplaceAllString(strings.TrimSpace(input[i]), " ")
			b = append(b, createSquareRow(utils.StrList_to_IntList(strings.Split(str, " "))))
		}
	}
	boardList = append(boardList, b)

	drawOrder := utils.StrList_to_IntList(strings.Split(input[0], ","))
	for _, v := range drawOrder {
		for _, j := range boardList {
			j.drawNumber(v)
			if j.checkBoard() {
				return j.calculateScore() * v
			}
		}
	}

	return 0
}

func Part2(input []string) int {
	var boardList []board
	var b board
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			// end of board, add to board list
			boardList = append(boardList, b)
			b = board{}
		} else {
			// still creating the board
			whitespaces := regexp.MustCompile(`\s+`)
			str := whitespaces.ReplaceAllString(strings.TrimSpace(input[i]), " ")
			b = append(b, createSquareRow(utils.StrList_to_IntList(strings.Split(str, " "))))
		}
	}
	boardList = append(boardList, b)

	var lastScore int
	drawOrder := utils.StrList_to_IntList(strings.Split(input[0], ","))
	for _, v := range drawOrder {
		for _, j := range boardList {
			j.drawNumber(v)
			for i, k := range boardList {
				if k.checkBoard() {
					if i+1 > len(boardList) {
						boardList = boardList[:len(boardList)-1]
					} else {
						boardList = append(boardList[:i], boardList[i+1:]...)
					}
					lastScore = j.calculateScore() * v
					break
				}
			}
		}
	}

	return lastScore
}
