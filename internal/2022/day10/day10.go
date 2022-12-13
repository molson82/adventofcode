package day10

import (
	"fmt"
	"time"
)

var tick = 10
var x int
var cycle int

func delayFunc(line string) {
	time.Sleep(time.Duration(tick * int(time.Millisecond)))
	// code to do the thing...
	fmt.Println(line)
	//cycle++
	//l := strings.Split(line, " ")
	//if len(l) == 1 {
	//// noop use case
	//return
	//}
	//v := l[1]
	//value, err := strconv.Atoi(v)
	//utils.CheckErr(err)
	//x += value
	//fmt.Printf("cycle: %v | x: %v\n", cycle, x)
}

func Part1(input []string) int {
	for i := 0; i < len(input); i++ {
		delayFunc(input[i])
	}

	return 0
}

func Part2(input []string) int {
	return 0
}
