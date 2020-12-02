package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(file string) int {

	f, err := os.Open(file)
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	var count int
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")
		key := strings.Trim(lineSplit[0], " ")
		value := strings.Trim(lineSplit[1], " ")
		if validate(key, value) {
			count++
		}
	}
	return count
}

func validate(key string, pass string) bool {
	keySplit := strings.Split(key, " ")
	hl := strings.Split(keySplit[0], "-")
	low, err := strconv.Atoi(hl[0])
	check(err)
	high, err2 := strconv.Atoi(hl[1])
	check(err2)
	char := keySplit[1]

	count := strings.Count(pass, char)
	if low <= count && high >= count {
		// fmt.Printf("K: %v | l: %v | h: %v | c: %v | p: %v\n", key, low, high, char, pass)
		return true
	}
	return false
}

/*
Part 1
 - attempt 1 = 459 | wrong
		- Using a map made it so there were keys getting overridden, since there are duplicates in the input
 - attempt 2 = 586 | correct
Part 2
 - attempt 1 =
*/
func main() {
	fmt.Println("Day 2 Part 1 ---------------------------------")
	total := readInput("input.txt")
	fmt.Printf("Total correct passwords: %v\n", total)
}
