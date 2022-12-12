package day9

import "fmt"

type Pos struct {
	Row int
	Col int
}

type Board struct {
	Width  int
	Height int
	Head   Pos
	Tail   Pos
}

func (b *Board) Print() {
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			if i == b.Head.Row && j == b.Head.Col {
				fmt.Print("H")
				continue
			}
			if (i == b.Tail.Row && j == b.Tail.Row) && (b.Head.Row != b.Tail.Row && b.Head.Col != b.Tail.Col) {
				fmt.Print("T")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func Part1(input []string) int {
	head := Pos{Row: 4, Col: 0}
	tail := Pos{Row: 4, Col: 0}
	board := Board{Width: 6, Height: 5, Head: head, Tail: tail}
	board.Print()
	return 0
}

func Part2(input []string) int {
	return 0
}
