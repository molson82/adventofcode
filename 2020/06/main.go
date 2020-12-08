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
	for j, g := range groups {
		count := 0
		if len(g.people) == 1 {
			count += len(g.people[0].questions)
		} else {
			per := g.people[0]
			str := ""
			for i := 1; i < len(g.people); i++ {
				for _, q := range per.questions {
					if !strings.Contains(str, string(q)) && strings.Contains(g.people[i].questions, string(q)) {
						count++
						str += g.people[i].questions
					}
				}
			}
		}
		fmt.Printf("%v | group: %v | Count: %v\n", g.people, j, count)
		total += count
	}

	return total
}

/*
Part 1
 - attempt 1 = 6170 | correct
Part 2
 - attempt 1 = 513 | wrong
		 - need to update / simplify part 2 question logic
 - attempt 2 =
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
