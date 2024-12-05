package day4

import (
	"fmt"
)

var directions = [][2]int{
	{0, 1},   // right
	{1, 0},   // down
	{1, 1},   // diagonal down-right
	{1, -1},  // diagonal down-left
	{0, -1},  // left
	{-1, 0},  // up
	{-1, -1}, // diagonal up-left
	{-1, 1},  // diagonal up-right
}

var xDirections = [][2]int{
	{1, 1},   // diagonal down-right
	{1, -1},  // diagonal down-left
	{-1, -1}, // diagonal up-left
	{-1, 1},  // diagonal up-right
}

func Part1(input []string) int {
	var ans int

	grid := parseGrid(input)
	ans = countXMAS(grid)
	fmt.Println("2024/4 part1 ans: ", ans)

	return ans
}

func Part2(input []string) int {
	var ans int

	grid := parseGrid(input)
	ans = countMAS(grid)
	fmt.Println("2024/4 part2 ans: ", ans)

	return ans
}

func isInBounds(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func countMAS(grid [][]rune) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'M' {
				for _, dir := range xDirections {
					if searchFromPosition(grid, i, j, dir[0], dir[1], "MAS") {
						// MAS found, check if other direction forms an X
						fmt.Println("mas found")
						if isXShape(grid, i+dir[0], j+dir[1], (dir[0])*-1, (dir[1])*-1) {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

// When the letter is A, verify the letters around it are M & S
func isXShape(grid [][]rune, x, y int, dx, dy int) bool {
	// Check bounds to ensure we don't go out of the grid
	nx, ny := x+dx, y+dy
	if !isInBounds(nx, ny, len(grid), len(grid[0])) {
		return false
	}

	if (grid[ny][nx] == 'M' && grid[(ny)*-1][(nx)*-1] == 'S') || (grid[ny][nx] == 'S' && grid[(ny)*-1][(nx)*-1] == 'M') {
		return true
	}

	return false
}

func searchFromPosition(grid [][]rune, x, y, dx, dy int, word string) bool {
	for i := 0; i < len(word); i++ {
		nx, ny := x+dx*i, y+dy*i
		if !isInBounds(nx, ny, len(grid), len(grid[0])) || grid[nx][ny] != rune(word[i]) {
			return false
		}
	}
	return true
}

func countXMAS(grid [][]rune) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'X' {
				for _, dir := range directions {
					if searchFromPosition(grid, i, j, dir[0], dir[1], "XMAS") {
						count++
					}
				}
			}
		}
	}
	return count
}

func parseGrid(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}
	return grid
}
