package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

var (
	redCount   int = 12
	greenCount int = 13
	blueCount  int = 14
)

func Part1(lines []string) int {
	var ans int

lines:
	for _, g := range lines {
		drawings := strings.Split(g, ": ")
		gameNum := strings.Split(drawings[0], "Game ")[1]
		n := utils.Str_to_Int(gameNum)

		sets := strings.Split(drawings[1], "; ")
		for _, s := range sets {
			if !checkPairs(s) {
				continue lines
			}
		}
		ans += n
	}
	return ans
}

func Part2(lines []string) int {
	var ans int

	for _, g := range lines {
		drawings := strings.Split(g, ": ")
		var maxBlue, maxRed, maxGreen int

		splitFunc := func(c rune) bool {
			return c == ',' || c == ';'
		}
		for _, pair := range strings.FieldsFunc(drawings[1], splitFunc) {
			p := strings.Trim(pair, " ")
			count := utils.Str_to_Int(strings.Split(p, " ")[0])
			color := strings.Trim(strings.Split(p, " ")[1], ";")
			if color == "blue" && count > maxBlue {
				maxBlue = count
			}
			if color == "red" && count > maxRed {
				maxRed = count
			}
			if color == "green" && count > maxGreen {
				maxGreen = count
			}
		}
		power := maxBlue * maxRed * maxGreen
		ans += power

	}

	return ans
}

func checkPairs(input string) bool {
	pattern := `(\d+)\s+(\w+)`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		count, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Printf("Error converting count: %v\n", err)
			continue
		}

		color := match[2]

		if color == "blue" && count > blueCount {
			return false
		}
		if color == "red" && count > redCount {
			return false
		}
		if color == "green" && count > greenCount {
			return false
		}
	}
	return true
}
