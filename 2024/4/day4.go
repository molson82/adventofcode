package day4

import (
	"fmt"
)

type Letter struct {
	L string
	X int
	Y int
}

type Row []Letter

type Grid []Row

var xLocations [][]int

const XMAS = "XMAS"

func Part1(input []string) int {
	var ans int

	letterGrid := createGrid(input)
	fmt.Println(xLocations)
	for _, v := range xLocations {
		if letterGrid.checkHorizontal(v[0], v[1], XMAS) {
			ans++
		}
		if letterGrid.checkVertical(v[0], v[1], XMAS) {
			ans++
		}
		if letterGrid.checkDiagonal(v[0], v[1], XMAS) {
			ans++
		}
	}

	fmt.Println("2024/4 part1 ans: ", ans)

	return ans
}

func Part2(input []string) int {
	var ans int
	// code here...
	fmt.Println("2024/4 part2 ans: ", ans)

	return ans
}

func createGrid(input []string) Grid {
	var newGrid Grid

	for i, r := range input {
		newRow := Row{}
		for j, l := range r {
			newLetter := Letter{
				L: string(l),
				X: j,
				Y: i,
			}
			// add to xlocations
			if string(l) == "X" {
				xLocations = append(xLocations, []int{j, i})
			}
			newRow = append(newRow, newLetter)
		}
		newGrid = append(newGrid, newRow)
	}

	return newGrid
}

func (g Grid) checkHorizontal(x, y int, w string) bool {
	// check forward
	if x+len(w) > len(g[y]) {
		return false
	}
	var matchf string
	for i := range w {
		matchf += g[y][x+i].L
	}
	if matchf == w {
		fmt.Printf("match horz forward: x: %v | y: %v\n", x, y)
		return true
	}

	// check backward
	if x-len(w) < 0 {
		return false
	}
	var matchb string
	for i := range w {
		matchb += g[y][x-i].L
	}
	if matchb == w {
		fmt.Printf("match horz backward: x: %v | y: %v\n", x, y)
		return true
	}

	return false
}

func (g Grid) checkVertical(x, y int, w string) bool {
	return true
}

func (g Grid) checkDiagonal(x, y int, w string) bool {
	return true
}
