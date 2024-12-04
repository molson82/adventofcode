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
var confirmedLocations [][]int

const XMAS = "XMAS"

func Part1(input []string) int {
	var ans int

	letterGrid := createGrid(input)
	for _, v := range xLocations {
		ans += letterGrid.checkHorizontal(v[0], v[1], XMAS)
		ans += letterGrid.checkVertical(v[0], v[1], XMAS)
		ans += letterGrid.checkDiagonal(v[0], v[1], XMAS)
	}
	// letterGrid.printConfirmedLocations()

	fmt.Println("2024/4 part1 ans: ", ans)

	return ans
}

func Part2(input []string) int {
	var ans int
	// code here...
	fmt.Println("2024/4 part2 ans: ", ans)

	return ans
}

func (g Grid) printConfirmedLocations() {
	fmt.Printf("\n")
	for y, r := range g {
		for x := range r {
			var printx bool
			// check if x,y is in confirmedLocations
			for _, v := range confirmedLocations {
				if v[0] == x && v[1] == y {
					fmt.Print("X")
					printx = true
				}
			}
			if !printx {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
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

func (g Grid) checkHorizontal(x, y int, w string) int {
	var matchf, matchb string
	var total int
	wl := len(w) - 1

	// check forward
	if x+len(w) <= len(g[y]) {
		for i := range w {
			matchf += g[y][x+i].L
		}
		if matchf == w {
			// fmt.Printf("match horz forward: x: %v | y: %v\n", x, y)
			addToConfirmedLocations(x, y)
			total++
		}
	}

	// check backward
	if x-wl >= 0 {
		for i := range w {
			matchb += g[y][x-i].L
		}
		if matchb == w {
			// fmt.Printf("match horz backward: x: %v | y: %v\n", x, y)
			addToConfirmedLocations(x, y)
			total++
		}
	}

	return total
}

func (g Grid) checkVertical(x, y int, w string) int {
	var matchu, matchd string
	var total int
	wl := len(w) - 1

	// check down
	if y+len(w) <= len(g) {
		for i := range w {
			matchu += g[y+i][x].L
		}
		if matchu == w {
			// fmt.Printf("match vert down: x: %v | y: %v\n", x, y)
			addToConfirmedLocations(x, y)
			total++
		}
	}

	// check up
	if y-wl >= 0 {
		for i := range w {
			matchd += g[y-i][x].L
		}
		if matchd == w {
			// fmt.Printf("match vert up: x: %v | y: %v\n", x, y)
			addToConfirmedLocations(x, y)
			total++
		}
	}

	return total
}

func (g Grid) checkDiagonal(x, y int, w string) int {
	var matchur, matchul, matchdr, matchdl string
	var total int
	wl := len(w) - 1

	// check up right
	if y-wl >= 0 && x+len(w) <= len(g[y]) {
		for i := range w {
			matchur += g[y-i][x+i].L
		}
		if matchur == w {
			// fmt.Printf("match up right: x: %v | y: %v\n", x, y)
			addToConfirmedLocations(x, y)
			total++
		}
	}

	// check up left
	if y-wl >= 0 && x-wl >= 0 {
		for i := range w {
			matchul += g[y-i][x-i].L
		}
		if matchul == w {
			// fmt.Printf("match up left: x: %v | y: %v\n", x, y)
			addToConfirmedLocations(x, y)
			total++
		}
	}

	// check down right
	if y+len(w) <= len(g) && x+len(w) <= len(g[y]) {
		for i := range w {
			matchdr += g[y+i][x+i].L
		}
		if matchdr == w {
			// fmt.Printf("match down right: x: %v | y: %v\n", x, y)
			addToConfirmedLocations(x, y)
			total++
		}
	}

	// check down left
	if y+len(w) <= len(g) && x-wl >= 0 {
		for i := range w {
			matchdl += g[y+i][x-i].L
		}
		if matchdl == w {
			// fmt.Printf("match down left: x: %v | y: %v\n", x, y)
			addToConfirmedLocations(x, y)
			total++
		}
	}

	return total
}

func addToConfirmedLocations(x, y int) {
	// Check if the x, y pair already exists in confirmedLocations
	for _, loc := range confirmedLocations {
		if loc[0] == x && loc[1] == y {
			return
		}
	}
	// Pair does not exist, so add it to confirmedLocations
	confirmedLocations = append(confirmedLocations, []int{x, y})
}
