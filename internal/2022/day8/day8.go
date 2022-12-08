package day8

import (
	"fmt"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

var width int
var height int

func Part1(input []string) int {
	var treeGrid [][]int
	for _, v := range input {
		row := strings.Split(v, "")
		rowInt := utils.StrList_to_IntList(row)
		treeGrid = append(treeGrid, rowInt)
	}
	width = len(treeGrid[0])
	height = len(treeGrid)
	visible := (2 * width) + (2 * height) - 4

OUTER:
	for row, v := range treeGrid {
	INNER:
		for col, t := range v {
			if row == 0 || row == len(treeGrid)-1 {
				continue OUTER
			}
			if col == 0 || col == len(v)-1 {
				continue INNER
			}
			// only evaluating the inner trees
			if checkVerticleUp(t, row, col, treeGrid) {
				//fmt.Println("vert up passed: ", t)
				visible++
			} else if checkVerticleDown(t, row, col, treeGrid) {
				//fmt.Println("vert down passed: ", t)
				visible++
			} else if checkHorizRight(t, row, col, treeGrid) {
				//fmt.Println("passed right: ", t)
				visible++
			} else if checkHorizLeft(t, row, col, treeGrid) {
				//fmt.Println("passed left: ", t)
				visible++
			}

		}
	}

	return visible
}

func Part2(input []string) int {

	return 0
}

func checkHorizLeft(tree, row, col int, grid [][]int) bool {
	for i := col - 1; i >= 0; i-- {
		//fmt.Println("horiz right: ", row, ",", col, " | ", tree, " vs ", grid[row][i])
		if tree > grid[row][i] {
			//fmt.Println("tree big")
			continue
		}
		return false
	}
	return true
}

func checkHorizRight(tree, row, col int, grid [][]int) bool {
	for i := col + 1; i < width; i++ {
		//fmt.Println("horiz right: ", row, ",", col, " | ", tree, " vs ", grid[row][i])
		if tree > grid[row][i] {
			//fmt.Println("tree big")
			continue
		}
		return false
	}
	return true
}

func checkVerticleUp(tree, row, col int, grid [][]int) bool {
	for i := row - 1; i >= 0; i-- {
		//fmt.Println("vert up: ", row, ",", col, " | ", tree, " vs ", grid[i][col])
		if tree > grid[i][col] {
			//fmt.Println("tree big")
			continue
		}
		return false
	}
	return true
}

func checkVerticleDown(tree, row, col int, grid [][]int) bool {
	for i := row + 1; i < height; i++ {
		//fmt.Println("vert down: ", row, ",", col, " | ", tree, " vs ", grid[i][col])
		if tree > grid[i][col] {
			//fmt.Println("tree big")
			continue
		}
		return false
	}
	return true
}

func printGrid(grid [][]int) {
	for _, v := range grid {
		fmt.Printf("%v\n", v)
	}
	fmt.Println()
}
