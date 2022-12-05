package day5

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/molson82/adventofcode/pkg/utils"
)

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	Items []string
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(strings []string, str string) {
	s.Items = append([]string{str}, strings...)
}

func (s *Stack) Move(str string) {
	s.Items = append(s.Items, str)
}

// Pop removes the top item from the stack and returns it.
func (s *Stack) Pop() string {
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}

// Print prints the contents of the stack.
func (s *Stack) Print() {
	fmt.Printf("Stack: %v\n", s.Items)
}

func printStacks(stacks []Stack) {
	for _, v := range stacks {
		v.Print()
	}
}

func Part1(input []string) string {
	stacks := buildDiagram(input)
	steps := buildSteps(input)

	for _, i := range steps {
		for j := 0; j < i[0]; j++ {
			// get value from stack
			val := stacks[i[1]-1].Pop()
			// push to new stack
			s := stacks[i[2]-1]
			s.Move(val)
			stacks[i[2]-1] = s
		}
	}

	// get all top values
	var resp string
	for _, v := range stacks {
		if len(v.Items) > 0 {
			resp = fmt.Sprintf("%s%s", resp, v.Pop())
		}
	}

	return resp
}

func Part2(input []string) string {
	stacks := buildDiagram(input)
	steps := buildSteps(input)

	for _, i := range steps {
		var vals []string
		for j := 0; j < i[0]; j++ {
			// get value from stack
			val := stacks[i[1]-1].Pop()
			vals = append(vals, val)
		}
		// push to new stack
		for k := len(vals) - 1; k >= 0; k-- {
			s := stacks[i[2]-1]
			s.Move(vals[k])
			stacks[i[2]-1] = s
		}
	}

	// get all top values
	var resp string
	for _, v := range stacks {
		if len(v.Items) > 0 {
			resp = fmt.Sprintf("%s%s", resp, v.Pop())
		}
	}

	return resp
}

func buildDiagram(input []string) []Stack {
	resp := make([]Stack, 9)
	for _, v := range input {
		if strings.Contains(v, "move") {
			return resp
		}
		for i, j := range v {
			if j != 32 && j != 91 && j != 93 && !unicode.IsNumber(j) { // != "[" && != "]"
				col := i / 4
				//fmt.Println("col: ", (i/4)+1, " | ", string(j))
				s := resp[col]
				s.Push(s.Items, string(j))
				resp[col] = s
			}
		}
	}

	return resp
}

func buildSteps(input []string) [][]int {
	var resp [][]int

	for _, v := range input {
		if strings.Contains(v, "move") {
			move, err := strconv.ParseInt(strings.Split(strings.Split(v, " from ")[0], "move ")[1], 10, 32)
			utils.CheckErr(err)
			from, err := strconv.ParseInt(strings.Split(strings.Split(v, " from ")[1], " to ")[0], 10, 32)
			utils.CheckErr(err)
			to, err := strconv.ParseInt(strings.Split(strings.Split(v, " from ")[1], " to ")[1], 10, 32)
			utils.CheckErr(err)

			temp := []int{int(move), int(from), int(to)}
			resp = append(resp, temp)
		}
	}

	return resp
}
