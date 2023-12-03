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
	ans = 0
	top = 0
	right = len(lines[0]) - 1
	left = 0
	bottom = len(lines) - 1

	fmt.Println("\n2023/3 part2 ans: ", ans)
	return ans
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
