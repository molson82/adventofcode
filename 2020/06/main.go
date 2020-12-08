package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var groupList []string
	var group string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "" {
			group += scanner.Text()
		} else {
			groupList = append(groupList, group)
			group = ""
		}
	}
	groupList = append(groupList, group)
	return groupList
}

func partOne(groups []string) int {
	total := 0
	for _, v := range groups {
		gcount := 0
		str := ""
		for _, c := range v {
			if !strings.Contains(str, string(c)) {
				str += string(c)
				gcount++
			}
		}
		total += gcount
	}
	return total
}

// PART 2 Stuff

type person struct {
	questions string
}

type group struct {
	people []person
}

func readInputPart2(file string) []group {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	var groupList []group
	var group group
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "" {
			p := person{scanner.Text()}
			group.people = append(group.people, p)
		} else {
			groupList = append(groupList, group)
			group.people = []person{}
		}
	}
	groupList = append(groupList, group)
	return groupList
}

func partTwo(groups []group) int {
	total := 0
	for _, g := range groups {
		str := ""
		count := 0
		for _, p := range g.people {
			for _, q := range p.questions {
				str += string(q)
				if strings.Contains(str, string(q)) {
					count++
				}
			}
		}
		total += count
	}

	return total
}

/*
Part 1
 - attempt 1 = 6170
Part 2
 - attempt 1 =
*/
func main() {
	fmt.Println("Advent of Code day 6")
	groupList := readInput("input.txt")
	fmt.Printf("Day 6 Part 1 ---------------------------------\n")
	fmt.Printf("Sum: %v\n", partOne(groupList))

	gList2 := readInputPart2("sample.txt")
	fmt.Printf("\nDay 6 Part 2 ---------------------------------\n")
	fmt.Printf("Sum: %v\n", partTwo(gList2))
}
