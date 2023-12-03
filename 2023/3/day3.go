package day3

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/molson82/adventofcode/pkg/utils"
)

type Point struct {
	X int
	Y int
}

var (
	ans    int
	top    int
	right  int
	left   int
	bottom int
)

func Part1(lines []string) int {
	initializeBounds(lines)
	validNumbers := make(map[Point]int)

	for i, l := range lines {
		for j, c := range l {
			if !unicode.IsDigit(c) && c != '.' {
				processCursor(i, j, l, lines, validNumbers)
			}
		}
	}

	for _, v := range validNumbers {
		ans += v
	}
	fmt.Println("\n2023/3 part1 ans: ", ans)
	return ans
}

func Part2(lines []string) int {
	totalGearRatio := 0

	for i, line := range lines {
		for j, c := range line {
			if c == '*' {
				gearRatio, isGear := calculateGearRatio(i, j, lines)
				if isGear {
					totalGearRatio += gearRatio
				}
			}
		}
	}

	fmt.Println("\n2023/3 part2 ans: ", totalGearRatio)
	return totalGearRatio
}

func calculateGearRatio(x, y int, lines []string) (int, bool) {
	numbers := make([]int, 0, 2)

	// New function to check and add number from a given direction
	addNumberIfPresent := func(dx, dy int) {
		nx, ny := x+dx, y+dy
		if nx >= 0 && nx < len(lines) && ny >= 0 && ny < len(lines[nx]) {
			if start, number, err := getNumber(nx, ny, lines); err == nil {
				fmt.Println("num: ", number)
				if start == ny || (dx != 0 && start <= y && y <= ny) { // Correctly identify the starting point of the number
					fmt.Println("num: ", number)
					numbers = append(numbers, number)
				}
			}
		}
	}

	// Check all four directions
	addNumberIfPresent(-1, 0) // Up
	addNumberIfPresent(1, 0)  // Down
	addNumberIfPresent(0, -1) // Left
	addNumberIfPresent(0, 1)  // Right

	if len(numbers) == 2 {
		return numbers[0] * numbers[1], true
	}
	return 0, false
}

// Updated GetNumber function
func getNumber(x, y int, lines []string) (int, int, error) {
	line := lines[x]
	if y < 0 || y >= len(line) || !unicode.IsDigit(rune(line[y])) {
		return 0, 0, fmt.Errorf("not a number")
	}

	start, end := y, y
	for start > 0 && unicode.IsDigit(rune(line[start-1])) {
		start--
	}
	for end < len(line)-1 && unicode.IsDigit(rune(line[end+1])) {
		end++
	}

	numberStr := line[start : end+1]
	number, err := strconv.Atoi(numberStr)
	return start, number, err
}

func initializeBounds(lines []string) {
	top = 0
	right = len(lines[0]) - 1
	left = 0
	bottom = len(lines) - 1
	ans = 0
}

func processCursor(i, j int, line string, lines []string, validNumbers map[Point]int) {
	checkAndAddNumber := func(x, y int, line string) {
		start, num, err := GetNumber(y, line)
		utils.CheckErr(err)
		validNumbers[Point{x, start}] = num
	}

	if i >= top && i != bottom {
		if CheckBelow(i, j, lines) {
			checkAndAddNumber(i+1, j, lines[i+1])
		}
		if j >= left && j != right && CheckDownRight(i, j, lines) {
			checkAndAddNumber(i+1, j+1, lines[i+1])
		}
		if j <= right && j != left && CheckDownLeft(i, j, lines) {
			checkAndAddNumber(i+1, j-1, lines[i+1])
		}
	}
	if i <= bottom && i != top {
		if CheckAbove(i, j, lines) {
			checkAndAddNumber(i-1, j, lines[i-1])
		}
		if j >= left && j != right && CheckUpRight(i, j, lines) {
			checkAndAddNumber(i-1, j+1, lines[i-1])
		}
		if j <= right && j != left && CheckUpLeft(i, j, lines) {
			checkAndAddNumber(i-1, j-1, lines[i-1])
		}
	}
	if j >= left && j != right && CheckRight(i, j, lines) {
		checkAndAddNumber(i, j+1, line)
	}
	if j <= right && j != left && CheckLeft(i, j, lines) {
		checkAndAddNumber(i, j-1, line)
	}
}

func GetNumber(y int, line string) (int, int, error) {
	if y < 0 || y >= len(line) {
		return 0, 0, fmt.Errorf("index out of range")
	}

	start, end := y, y
	for ; start > 0 && unicode.IsDigit(rune(line[start-1])); start-- {
	}
	for ; end < len(line)-1 && unicode.IsDigit(rune(line[end+1])); end++ {
	}

	numberStr := line[start : end+1]
	n, err := strconv.Atoi(numberStr)
	return start, n, err
}

func CheckDownRight(x, y int, lines []string) bool {
	return unicode.IsDigit(rune(lines[x+1][y+1]))
}

func CheckDownLeft(x, y int, lines []string) bool {
	return unicode.IsDigit(rune(lines[x+1][y-1]))
}

func CheckUpLeft(x, y int, lines []string) bool {
	return unicode.IsDigit(rune(lines[x-1][y-1]))
}

func CheckUpRight(x, y int, lines []string) bool {
	return unicode.IsDigit(rune(lines[x-1][y+1]))
}

func CheckLeft(x, y int, lines []string) bool {
	return unicode.IsDigit(rune(lines[x][y-1]))
}

func CheckRight(x, y int, lines []string) bool {
	return unicode.IsDigit(rune(lines[x][y+1]))
}

func CheckAbove(x, y int, lines []string) bool {
	return unicode.IsDigit(rune(lines[x-1][y]))
}

func CheckBelow(x, y int, lines []string) bool {
	return unicode.IsDigit(rune(lines[x+1][y]))
}
