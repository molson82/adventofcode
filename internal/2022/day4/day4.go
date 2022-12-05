package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

func Part1(input []string) int {
	var sum int
	for _, v := range input {
		if checkPairs(v) {
			sum += 1
		}
	}

	return sum
}

func Part2(input []string) int {
	var sum int
	for _, v := range input {
		if checkPairsPart2(v) {
			sum += 1
		}
	}
	return sum
}

func checkPairsPart2(line string) bool {
	ranges := strings.Split(line, ",")
	temp := strings.Split(ranges[0], "-")
	temp2 := strings.Split(ranges[1], "-")

	start, err := strconv.ParseInt(temp[0], 10, 32)
	utils.CheckErr(err)
	end, err := strconv.ParseInt(temp[1], 10, 32)
	utils.CheckErr(err)

	start2, err := strconv.ParseInt(temp2[0], 10, 32)
	utils.CheckErr(err)
	end2, err := strconv.ParseInt(temp2[1], 10, 32)
	utils.CheckErr(err)

	// Check if range1 contains any numbers from range2.
	for _, n := range []int64{start2, end2} {
		if n >= start && n <= end {
			return true
		}
	}

	// Check if range2 contains any numbers from range1.
	for _, n := range []int64{start, end} {
		if n >= start2 && n <= end2 {
			return true
		}
	}

	return false
}

func checkPairs(line string) bool {
	ranges := strings.Split(line, ",")
	temp := strings.Split(ranges[0], "-")
	temp2 := strings.Split(ranges[1], "-")

	start, err := strconv.ParseInt(temp[0], 10, 32)
	utils.CheckErr(err)
	end, err := strconv.ParseInt(temp[1], 10, 32)
	utils.CheckErr(err)

	start2, err := strconv.ParseInt(temp2[0], 10, 32)
	utils.CheckErr(err)
	end2, err := strconv.ParseInt(temp2[1], 10, 32)
	utils.CheckErr(err)

	if (start >= start2 && end <= end2) || (start2 >= start && end2 <= end) {
		return true
	}

	return false
}

func buildPairs(line string) []string {
	var resp []string
	ranges := strings.Split(line, ",")
	for _, v := range ranges {
		temp := strings.Split(v, "-")
		start, err := strconv.ParseInt(temp[0], 10, 32)
		utils.CheckErr(err)
		end, err := strconv.ParseInt(temp[1], 10, 32)
		utils.CheckErr(err)

		var rangeStr string
		for i := start; i <= end; i++ {
			if rangeStr == "" {
				rangeStr = fmt.Sprintf("%s", fmt.Sprint(i))
			} else {
				rangeStr = fmt.Sprintf("%s,%s", rangeStr, fmt.Sprint(i))
			}
		}
		resp = append(resp, rangeStr)
	}

	return resp
}
