package day3

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/molson82/adventofcode/pkg/utils"
)

func Part1(lines []string) int {
	ans = 0
	top = 0
	right = len(lines[0]) - 1
	left = 0
	bottom = len(lines) - 1

	validNumbers := make(map[struct {
		x int
		y int
	}]int)
	for i, l := range lines {
		for j, c := range l {
			if !unicode.IsDigit(c) && c != '.' {
				// fmt.Println("cursor: ", string(c))
				if i >= top && i != bottom {
					if CheckBelow(i, j, lines) {
						start, num, err := GetNumber(j, lines[i+1])
						// fmt.Println("num: ", num)
						utils.CheckErr(err)
						validNumbers[struct {
							x int
							y int
						}{i + 1, start}] = num
					}
					if (j >= left && j != right) && CheckDownRight(i, j, lines) {
						start, num, err := GetNumber(j+1, lines[i+1])
						// fmt.Println("num: ", num)
						utils.CheckErr(err)
						validNumbers[struct {
							x int
							y int
						}{i + 1, start}] = num
					}
					if (j <= right && j != left) && CheckDownLeft(i, j, lines) {
						start, num, err := GetNumber(j-1, lines[i+1])
						// fmt.Println("num: ", num)
						utils.CheckErr(err)
						validNumbers[struct {
							x int
							y int
						}{i + 1, start}] = num
					}
				}
				if i <= bottom && i != top {
					if CheckAbove(i, j, lines) {
						start, num, err := GetNumber(j, lines[i-1])
						// fmt.Println("num: ", num)
						utils.CheckErr(err)
						validNumbers[struct {
							x int
							y int
						}{i - 1, start}] = num
					}
					if (j >= left && j != right) && CheckUpRight(i, j, lines) {
						start, num, err := GetNumber(j+1, lines[i-1])
						// fmt.Println("num: ", num)
						utils.CheckErr(err)
						validNumbers[struct {
							x int
							y int
						}{i - 1, start}] = num
					}
					if (j <= right && j != left) && CheckUpLeft(i, j, lines) {
						start, num, err := GetNumber(j-1, lines[i-1])
						// fmt.Println("num: ", num)
						utils.CheckErr(err)
						validNumbers[struct {
							x int
							y int
						}{i - 1, start}] = num
					}
				}
				if j >= left && j != right {
					if CheckRight(i, j, lines) {
						start, num, err := GetNumber(j+1, l)
						// fmt.Println("num: ", num)
						utils.CheckErr(err)
						validNumbers[struct {
							x int
							y int
						}{i, start}] = num
					}
				}
				if j <= right && j != left {
					if CheckLeft(i, j, lines) {
						start, num, err := GetNumber(j-1, l)
						// fmt.Println("num: ", num)
						utils.CheckErr(err)
						validNumbers[struct {
							x int
							y int
						}{i, start}] = num
					}
				}
			}
		}
	}
	// fmt.Println("Valid num: ", validNumbers)
	for _, v := range validNumbers {
		ans += v
	}
	fmt.Println("\n2023/3 part1 ans: ", ans)
	return ans
}

var (
	ans    int
	top    int
	right  int
	left   int
	bottom int
)

func Part2(lines []string) int {
	ans = 0
	top = 0
	right = len(lines[0]) - 1
	left = 0
	bottom = len(lines) - 1

	fmt.Println("\n2023/3 part2 ans: ", ans)
	return ans
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
	if x >= top && x != bottom {
		return unicode.IsDigit(rune(lines[x+1][y]))
	}
	return false
}

func GetNumber(y int, line string) (int, int, error) {
	// fmt.Printf("y: %v | line: %v\n", y, line)
	if y < 0 || y >= len(line) {
		return 0, 0, fmt.Errorf("index out of range")
	}

	// Find the start of the number
	start := y
	for start > 0 && unicode.IsDigit(rune(line[start-1])) {
		start--
	}

	// Find the end of the number
	end := y
	for end < len(line)-1 && unicode.IsDigit(rune(line[end+1])) {
		end++
	}
	numberStr := line[start : end+1]
	n, err := strconv.Atoi(numberStr)
	return start, n, err
}
