package main

import (
	"bufio"
	"fmt"
	"os"
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

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func countTrees(lines []string, xdelta int, ydelta int) (numTrees int) {
	curpos := 0
	w := len(lines[0])

	for i := ydelta; i < len(lines); i += ydelta {
		curpos = (curpos + xdelta) % w
		if string(lines[i][curpos]) == "#" {
			numTrees++
		}
	}
	return numTrees
}

/*
Part 1
 - attempt 1 = 193 | correct
Part 2
 - attempt 1 = 1355323200 | correct
*/
func main() {
	file := readInput("input.txt")
	fmt.Println("Day 3 Part 1 ---------------------------------")
	fmt.Printf("Total trees: %v\n", countTrees(file, 3, 1))
	fmt.Printf("\nDay 3 Part 2 ---------------------------------\n")
	total := countTrees(file, 1, 1)
	total *= countTrees(file, 3, 1)
	total *= countTrees(file, 5, 1)
	total *= countTrees(file, 7, 1)
	total *= countTrees(file, 1, 2)
	fmt.Printf("Total trees: %v\n", total)
}
