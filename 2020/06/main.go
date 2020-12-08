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
		count := 0
		if len(g.people) == 1 {
			count += len(g.people[0].questions)
		} else {
			per := g.people[0]
			str := ""
			for _, q := range per.questions {
				num := 0
				for i := 1; i < len(g.people); i++ {
					if !strings.Contains(str, string(q)) && strings.Contains(g.people[i].questions, string(q)) {
						num++
					} else {
						break
					}
				}
				if num == len(g.people)-1 {
					count++
					str += string(q)
				}
			}
		}
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
 - attempt 2 = 3116 | wrong
		 - I only need to compare the first two people in the group. NOT every person
 - attempt 3 = 3019 | wrong
		 - Wait, I was wrong. I do need to check every person in the group. This is confusing
		 - I needed to flip my loop. Iterate over questions first, then each person in the group
 - attempt 4 = 2947 | correct
*/
func main() {
	fmt.Println("Advent of Code day 6")
	groupList := readInput("input.txt")
	fmt.Printf("Day 6 Part 1 ---------------------------------\n")
	fmt.Printf("Sum: %v\n", partOne(groupList))

	gList2 := readInputPart2("input.txt")
	fmt.Printf("\nDay 6 Part 2 ---------------------------------\n")
	fmt.Printf("Sum: %v\n", partTwo(gList2))
}
