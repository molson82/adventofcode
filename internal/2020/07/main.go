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

func contains(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func findGold(bagMap map[string]bag) []string {
	var goldBagParents []string
	for _, bag := range bagMap {
		for _, v := range bag.bags {
			if v.color == "shiny gold" {
				goldBagParents = append(goldBagParents, bag.color)
			}
		}
	}
	return goldBagParents
}

func verifyGold(bag bag) bool {
	if len(bag.bags) == 0 || bag.color == "shiny gold" {
		return false
	}
	// fmt.Printf("Checking bag: %v\n", bag)
	for _, b := range bag.bags {
		if contains(goldBagParents, b.color) {
			// fmt.Printf("found gold parent: %v\n", b.color)
			goldBagParents = append(goldBagParents, bag.color)
			return true
		}
		// fmt.Printf("nothing found\n")
		return verifyGold(b)
	}
	return false
}

var goldBagParents = findGold(readInput("input.txt"))

func partOne(bagMap map[string]bag) int {
	total := len(goldBagParents)
	// fmt.Printf("%v\n%v\n", goldBagParents, len(goldBagParents))
	for _, bag := range bagMap {
		// fmt.Printf("On bag: %v\n", bag)
		for _, v := range bag.bags {
			if contains(goldBagParents, v.color) {
				// fmt.Printf("Found Gold: %v\n", bag.color)
				total++
				break
			} else {
				if verifyGold(bagMap[v.color]) {
					// fmt.Printf("Found Gold: %v\n", bag.color)
					total++
					break
				}
			}
		}
	}
	return total
}

/*
Part 1
 - attempt 1 = 21 | wrong
		 - it's too low. I need to iterate over the other bags that may lead to one in the list of gold parent bags.
 - attempt 2 = 26 | wrong
		 - same
 - attempt 3 = 37 | wrong
     - Logic probably isn't write
Part 2
 - attempt 1 =
*/
func main() {
	fmt.Println("Advent of Code day 7")
	input := readInput("input.txt")
	fmt.Printf("Day 7 Part 1 ---------------------------------\n")
	fmt.Printf("Total Gold bags: %v\n", partOne(input))
	fmt.Printf("\nDay 7 Part 2 ---------------------------------\n")
}
