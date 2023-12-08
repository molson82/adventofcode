package day8

import (
	"fmt"
	"strings"
	"sync"
)

func Part1(lines []string) int {
	var ans int
	seq := strings.Split(lines[0], "")

	network := make(map[string][]string)
	for i := 2; i < len(lines); i++ {
		l := lines[i]
		spl := strings.Split(l, "= ")
		key := strings.Fields(spl[0])[0]
		valSpl := strings.Split(spl[1], ", ")
		leftVal := valSpl[0][1:]
		rightVal := valSpl[1][:len(valSpl[1])-1]

		network[key] = []string{leftVal, rightVal}
	}

	// go over seq
	stepKey := "AAA"
	for i := 0; i < len(seq); i++ {
		ans++
		if stepKey == "ZZZ" {
			break
		}
		s := seq[i]
		currNode := network[stepKey]
		if s == "R" {
			stepKey = currNode[1]
		} else {
			stepKey = currNode[0]
		}
		// reset loop
		if i == len(seq)-1 {
			i = -1
		}
	}
	fmt.Println("2023/8 part1 ans: ", ans-1)
	return ans - 1
}

func Part2(lines []string) int {
	var startingNodes []string
	seq := strings.Split(lines[0], "")

	network := make(map[string][]string)
	for i := 2; i < len(lines); i++ {
		l := lines[i]
		spl := strings.Split(l, "= ")
		key := strings.Fields(spl[0])[0]
		if key[len(key)-1] == 'A' {
			startingNodes = append(startingNodes, key)
		}
		valSpl := strings.Split(spl[1], ", ")
		leftVal := valSpl[0][1:]
		rightVal := valSpl[1][:len(valSpl[1])-1]

		network[key] = []string{leftVal, rightVal}
	}

	// concurrently go over the starting nodes
	var wg sync.WaitGroup
	type stepCount struct {
		mu sync.Mutex
		c  int
	}
	steps := stepCount{}

	// NOTES:
	// The goal is NOT to run them concurrently, but
	// at the same time, step by step, to see if each starting
	// point ends in a Z. If not, keep going. So even if one node
	// does reach a Z, it has to restart in a loop until all
	// starting nodes have a Z at the end.
	// Don't think go routines will work since they will go out
	// of sync and not follow steps at the same rate.
	for _, sn := range startingNodes {
		wg.Add(1)
		go func(sn string, seq []string) {
			fmt.Println()
			stepKey := sn
			fmt.Printf("stepKey: %v | ", stepKey)
			for i := 0; i < len(seq); i++ {
				if stepKey[len(stepKey)-1] == 'Z' {
					fmt.Printf("stepKey: %v | ", stepKey)
					break
				}
				s := seq[i]
				fmt.Printf("step: %v | ", s)
				currNode := network[stepKey]
				if s == "R" {
					stepKey = currNode[1]
				} else {
					stepKey = currNode[0]
				}
				steps.mu.Lock()
				fmt.Printf("c++ | ")
				steps.c++
				steps.mu.Unlock()
				// reset loop
				if i == len(seq)-1 {
					i = -1
				}
			}
			wg.Done()
		}(sn, seq)
	}

	wg.Wait()

	fmt.Println("2023/8 part2 ans: ", steps.c)
	return steps.c
}
