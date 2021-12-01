package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(file string) []string {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	var boardingPassList []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		boardingPassList = append(boardingPassList, scanner.Text())
	}
	return boardingPassList
}

func getRow(pass string) int {
	currMin := 0
	currMax := 127
	for i := 0; i < 7; i++ {
		if string(pass[i]) == "F" {
			currMax = (currMax + currMin) / 2
			// fmt.Printf("lower half - min: %v | max: %v\n", currMin, currMax)
		} else {
			currMin = currMin + ((currMax - currMin + 1) / 2)
			// fmt.Printf("upper half - min: %v | max: %v\n", currMin, currMax)
		}
	}
	return currMax
}

func getColumn(pass string) int {
	cols := pass[len(pass)-3:]
	currMin := 0
	currMax := 7
	for _, v := range cols {
		if string(v) == "L" {
			currMax = (currMax + currMin) / 2
			// fmt.Printf("lower half - min: %v | max: %v\n", currMin, currMax)
		} else {
			currMin = currMin + ((currMax - currMin + 1) / 2)
			// fmt.Printf("upper half - min: %v | max: %v\n", currMin, currMax)
		}
	}
	return currMax
}

func partone(list []string) int {
	highest := 0
	for _, v := range list {
		row := getRow(v)
		col := getColumn(v)
		sID := (row * 8) + col
		// fmt.Printf("List: %v | sID: %v\n", v, sID)
		if sID > highest {
			highest = sID
		}
	}
	return highest
}

// smallest sid = 7
// largest sid = 908
func partTwo(list []string) int {
	var sidList []int
	for _, v := range list {
		row := getRow(v)
		col := getColumn(v)
		sID := (row * 8) + col
		sidList = append(sidList, sID)
	}
	sort.Ints(sidList)
	// find missing number from 7 - 908
	check := 7
	for _, v := range sidList {
		if v == check {
			check++
		} else {
			break
		}
	}
	return check
}

/*
Part 1
 - attempt 1 = 908 | correct
Part 2
 - attempt 1 = 619 | correct
*/
func main() {
	fmt.Println("Advent of Code day 5")
	bplist := readInput("input.txt")
	fmt.Printf("Day 5 Part 1 ---------------------------------\n")
	fmt.Printf("Highest sID: %v\n", partone(bplist))
	fmt.Printf("\nDay 5 Part 2 ---------------------------------\n")
	fmt.Printf("My sID: %v\n", partTwo(bplist))
}
