package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

type Pos struct {
	Row int
	Col int
}

type Board struct {
	TailVisits map[Pos]int
	Width      int
	Height     int
	Head       Pos
	Tail       Pos
}

func (b *Board) Print() {
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			if i == b.Head.Row && j == b.Head.Col {
				fmt.Print("H")
				continue
			}
			if (i == b.Tail.Row && j == b.Tail.Col) && !(b.Head.Row == b.Tail.Row && b.Head.Col == b.Tail.Col) {
				fmt.Print("T")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *Board) Move(dir string, num int) {
	switch dir {
	case "R":
		b.moveRight(num)
	case "U":
		b.moveUp(num)
	case "L":
		b.moveLeft(num)
	case "D":
		b.moveDown(num)
	}
}

func (b *Board) moveTail(dir string) {
	// check if tail is touching
	if (b.Tail.Col-1 == b.Head.Col && b.Tail.Row == b.Head.Row) ||
		(b.Tail.Col+1 == b.Head.Col && b.Tail.Row == b.Head.Row) ||
		(b.Tail.Row-1 == b.Head.Row && b.Tail.Col == b.Head.Col) ||
		(b.Tail.Row+1 == b.Head.Row && b.Tail.Col == b.Head.Col) {
		//fmt.Println("tail is touching")
		return
	}
	// chech diagonals
	if (b.Tail.Col+1 == b.Head.Col && b.Tail.Row-1 == b.Head.Row) ||
		(b.Tail.Col-1 == b.Head.Col && b.Tail.Row-1 == b.Head.Row) ||
		(b.Tail.Col-1 == b.Head.Col && b.Tail.Row+1 == b.Head.Row) ||
		(b.Tail.Col+1 == b.Head.Col && b.Tail.Row+1 == b.Head.Row) {
		//fmt.Println("tail is touching diagonal")
		return
	}
	// move tail
	switch dir {
	case "R":
		b.Tail.Col = b.Head.Col - 1
		b.Tail.Row = b.Head.Row
	case "U":
		b.Tail.Row = b.Head.Row + 1
		b.Tail.Col = b.Head.Col
	case "L":
		b.Tail.Col = b.Head.Col + 1
		b.Tail.Row = b.Head.Row
	case "D":
		b.Tail.Row = b.Head.Row - 1
		b.Tail.Col = b.Head.Col
	}
	// add to tail visits
	pos := Pos{Row: b.Tail.Row, Col: b.Tail.Col}
	b.TailVisits[pos] = b.TailVisits[pos] + 1
}

func (b *Board) moveRight(num int) {
	for i := 0; i < num; i++ {
		b.Head.Col++
		//b.Print()
		b.moveTail("R")
		//b.Print()
	}
}

func (b *Board) moveLeft(num int) {
	for i := 0; i < num; i++ {
		b.Head.Col--
		//b.Print()
		b.moveTail("L")
		//b.Print()
	}
}

func (b *Board) moveUp(num int) {
	for i := 0; i < num; i++ {
		b.Head.Row--
		//b.Print()
		b.moveTail("U")
		//b.Print()
	}
}

func (b *Board) moveDown(num int) {
	for i := 0; i < num; i++ {
		b.Head.Row++
		//b.Print()
		b.moveTail("D")
		//b.Print()
	}
}

func Part1(input []string) int {
	head := Pos{Row: 4, Col: 0}
	tail := Pos{Row: 4, Col: 0}
	board := Board{Width: 6, Height: 5, Head: head, Tail: tail, TailVisits: make(map[Pos]int)}

	for _, v := range input {
		l := strings.Split(v, " ")
		dir := l[0]
		n := l[1]
		num, err := strconv.Atoi(n)
		utils.CheckErr(err)

		board.Move(dir, num)
	}
	return len(board.TailVisits)
}

func Part2(input []string) int {
	return 0
}
