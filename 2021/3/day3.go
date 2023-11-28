package day3

import (
	"fmt"
	"strconv"
)

type bitPos struct {
	zero int
	one  int
}

func countBits(input []string) []bitPos {
	byte := make([]bitPos, len(input[0]))
	for _, v := range input {
		for j, k := range v {
			if string(k) == "0" {
				byte[j].zero++
			} else {
				byte[j].one++
			}
		}
	}
	return byte
}

func countBitsAtIndex(input []string, index int) []int {
	var zero int
	var one int
	for _, v := range input {
		if string(v[index]) == "0" {
			zero++
		} else {
			one++
		}
	}

	return []int{zero, one}
}

func getOxygenList(input []string, index int, count []int) []string {
	var common string
	if count[0] > count[1] {
		common = "0"
	} else {
		common = "1"
	}

	var newList []string
	for _, v := range input {
		if string(v[index]) == common {
			newList = append(newList, v)
		}
	}
	return newList
}

func getCO2List(input []string, index int, count []int) []string {
	var common string
	if count[1] < count[0] {
		common = "1"
	} else {
		common = "0"
	}

	var newList []string
	for _, v := range input {
		if string(v[index]) == common {
			newList = append(newList, v)
		}
	}
	return newList
}

func Part1(input []string) int64 {
	var gammaRate string
	var epsilonRate string

	bitCounts := countBits(input)
	for _, v := range bitCounts {
		if v.one > v.zero {
			gammaRate = fmt.Sprintf("%s1", gammaRate)
			epsilonRate = fmt.Sprintf("%s0", epsilonRate)
		} else {
			gammaRate = fmt.Sprintf("%s0", gammaRate)
			epsilonRate = fmt.Sprintf("%s1", epsilonRate)
		}
	}

	// convert to decimal
	gamma, _ := strconv.ParseInt(gammaRate, 2, 0)
	epsilon, _ := strconv.ParseInt(epsilonRate, 2, 0)

	return gamma * epsilon
}

func Part2(input []string) int64 {

	oxygenList := input
	for i := 0; i < len(input[0]); i++ {
		if len(oxygenList) == 1 {
			break
		}
		bitCount := countBitsAtIndex(oxygenList, i)
		oxygenList = getOxygenList(oxygenList, i, bitCount)
	}
	oxygenNum, _ := strconv.ParseInt(oxygenList[0], 2, 0)

	co2List := input
	for i := 0; i < len(input[0]); i++ {
		if len(co2List) == 1 {
			break
		}
		bitCount := countBitsAtIndex(co2List, i)
		co2List = getCO2List(co2List, i, bitCount)
	}
	co2Num, _ := strconv.ParseInt(co2List[0], 2, 0)

	return oxygenNum * co2Num
}
