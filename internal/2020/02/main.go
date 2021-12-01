package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(file string) (int, int) {
	f, err := os.Open(file)
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	var count int
	var count2 int
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")
		policy := strings.Trim(lineSplit[0], " ")
		pass := strings.Trim(lineSplit[1], " ")
		if validate(policy, pass) {
			count++
		}
		if tobogganValidate(policy, pass) {
			count2++
		}
	}
	return count, count2
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

func tobogganValidate(key string, pass string) bool {
	keySplit := strings.Split(key, " ")
	hl := strings.Split(keySplit[0], "-")
	low, err := strconv.Atoi(hl[0])
	check(err)
	high, err2 := strconv.Atoi(hl[1])
	check(err2)
	char, _ := utf8.DecodeLastRuneInString(keySplit[1])
	passRune := []rune(pass)

	// xor
	return (passRune[low-1] == char || passRune[high-1] == char) && passRune[low-1] != passRune[high-1]
}

/*
Part 1
 - attempt 1 = 459 | wrong
		- Using a map made it so there were keys getting overridden, since there are duplicates in the input
 - attempt 2 = 586 | correct
Part 2
 - attempt 1 = 143 | wrong
		- had a bad xor
 - attempt 2 = 352 | correct
*/
func main() {
	total1, total2 := readInput("input.txt")
	fmt.Println("Day 2 Part 1 ---------------------------------")
	fmt.Printf("Total correct passwords: %v\n", total1)
	fmt.Printf("\nDay 2 Part 2 ---------------------------------\n")
	fmt.Printf("Total correct passwords: %v\n", total2)
}
