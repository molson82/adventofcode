package day9

import (
	"fmt"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

func Part1(lines []string) int {
	var ans int
	var seqList [][]int
	for _, l := range lines {
		var seq []int
		strList := strings.Split(l, " ")
		for _, s := range strList {
			seq = append(seq, utils.Str_to_Int(string(s)))
		}
		seqList = append(seqList, seq)
	}
	for _, sl := range seqList {
		ans += CalcSequence(sl, [][]int{sl})
	}
	fmt.Println("2023/9 part1 ans: ", ans)
	return ans
}

func Part2(lines []string) int {
	var ans int
	var seqList [][]int
	for _, l := range lines {
		var seq []int
		strList := strings.Split(l, " ")
		for _, s := range strList {
			seq = append(seq, utils.Str_to_Int(string(s)))
		}
		seqList = append(seqList, seq)
	}
	for _, sl := range seqList {
		ans -= CalcSequencePart2(sl, [][]int{sl})
	}
	fmt.Println("2023/9 part2 ans: ", ans)
	return ans
}

func CalcSequence(seq []int, seqList [][]int) int {
	var diff []int
	if seq[0] == 0 && seq[len(seq)-1] == 0 {
		return GetNextNum(seqList)
	}
	for i := 0; i < len(seq)-1; i++ {
		diff = append(diff, seq[i+1]-seq[i])
	}
	seqList = append(seqList, diff)
	return CalcSequence(diff, seqList)
}

func CalcSequencePart2(seq []int, seqList [][]int) int {
	var diff []int
	if seq[0] == 0 && seq[len(seq)-1] == 0 {
		return GetNextNumLeft(seqList)
	}
	for i := 0; i < len(seq)-1; i++ {
		diff = append(diff, seq[i+1]-seq[i])
	}
	seqList = append(seqList, diff)
	return CalcSequencePart2(diff, seqList)
}

func GetNextNumLeft(seqList [][]int) int {
	var ans int
	for i := len(seqList) - 1; i >= 0; i-- {
		l := seqList[i]
		firstValue := l[0]
		fmt.Println("ans: ", ans)
		ans -= firstValue
	}
	return ans
}

func GetNextNum(seqList [][]int) int {
	var ans int
	for i := len(seqList) - 1; i >= 0; i-- {
		l := seqList[i]
		lastValue := l[len(l)-1]
		ans += lastValue
	}
	return ans
}
