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

type bag struct {
	color string
	bags  []bag
	count int
}

func readInput(file string) map[string]bag {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	bagMap := make(map[string]bag)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		splitOne := strings.Split(scanner.Text(), "contain ")
		bagKey := strings.Split(splitOne[0], " bags")[0]
		bagValue := bag{
			bagKey,
			[]bag{},
			1,
		}
		// add list of bags to bag value
		if strings.Contains(splitOne[1], "no other bags") {
			bagMap[bagKey] = bagValue
		} else if strings.Contains(splitOne[1], ",") {
			bagList := strings.Split(splitOne[1], ", ")
			for _, b := range bagList {
				newBagCount, err := strconv.Atoi(string(b[0]))
				check(err)
				newBagColor := strings.Split(b, " bag")[0][2:]
				newBag := bag{
					newBagColor,
					[]bag{},
					newBagCount,
				}
				bagValue.bags = append(bagValue.bags, newBag)
			}
			bagMap[bagKey] = bagValue
		} else {
			newBagCount, err := strconv.Atoi(string(splitOne[1][0]))
			check(err)
			newBagColor := strings.Split(splitOne[1], " bag")[0][2:]
			newBag := bag{
				newBagColor,
				[]bag{},
				newBagCount,
			}
			bagValue.bags = append(bagValue.bags, newBag)
			bagMap[bagKey] = bagValue
		}
	}

	return bagMap
}

/*
Part 1
 - attempt 1 =
Part 2
 - attempt 1 =
*/
func main() {
	fmt.Println("Advent of Code day 7")
	readInput("input.txt")
	fmt.Printf("Day 7 Part 1 ---------------------------------\n")

	fmt.Printf("\nDay 7 Part 2 ---------------------------------\n")
}
