package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	var ans int
	for _, v := range lines {
		lineNum := getLineNum(v)
		num, err := strconv.ParseInt(lineNum, 10, 0)
		if err != nil {
			log.Fatalln("err: ", err)
		}
		ans += int(num)
	}
	fmt.Println("2023/1 part1 ans: ", ans)
	return ans
}

func Part2(lines []string) int {
	var ans int
	for _, v := range lines {
		lineNum := GetLineNumV2(v)
		num, err := strconv.ParseInt(lineNum, 10, 0)
		if err != nil {
			log.Fatalln("err: ", err)
		}
		ans += int(num)
	}

	fmt.Println("2023/1 part2 ans: ", ans)
	return ans
}

func GetLineNumV2(str string) string {
	var lineNum string
	var word string
	var lastChar string
	for _, c := range str {
		lastChar = string(c)
		word = fmt.Sprintf("%s%s", word, string(c))
		// check if it's a spelled out number
		if sn := spelledNumToInt(word); sn != "" {
			lineNum = fmt.Sprintf("%s%s", lineNum, sn)
			word = lastChar
		}
		// check if it's a valid number
		_, err := strconv.ParseInt(string(c), 10, 0)
		if err != nil {
			continue
		} else {
			lineNum = fmt.Sprintf("%s%s", lineNum, string(c))
			word = lastChar
		}
	}
	lineNum = fmt.Sprintf("%s%s", string(lineNum[0]), string(lineNum[len(lineNum)-1]))
	return lineNum
}

func getLineNum(str string) string {
	var lineNum string
	for _, c := range str {
		_, err := strconv.ParseInt(string(c), 10, 0)
		if err != nil {
			continue
		} else {
			lineNum = fmt.Sprintf("%s%s", lineNum, string(c))
		}
	}
	lineNum = fmt.Sprintf("%s%s", string(lineNum[0]), string(lineNum[len(lineNum)-1]))
	return lineNum
}

func spelledNumToInt(str string) string {
	if strings.Contains(str, "one") {
		return "1"
	}
	if strings.Contains(str, "two") {
		return "2"
	}
	if strings.Contains(str, "three") {
		return "3"
	}
	if strings.Contains(str, "four") {
		return "4"
	}
	if strings.Contains(str, "five") {
		return "5"
	}
	if strings.Contains(str, "six") {
		return "6"
	}
	if strings.Contains(str, "seven") {
		return "7"
	}
	if strings.Contains(str, "eight") {
		return "8"
	}
	if strings.Contains(str, "nine") {
		return "9"
	}

	return ""
}
