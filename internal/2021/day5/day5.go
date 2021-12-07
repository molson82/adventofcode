package day5

import (
	"fmt"
	"math"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

type point struct {
	x int
	y int
}

type line []*point

func (l line) print() {
	fmt.Printf("\n")
	for _, v := range l {
		fmt.Printf("%v,%v  ", v.x, v.y)
	}
	fmt.Printf("\n")
}

func (l line) drawOnGraph(graph *grid) {
	// check if horizontal line
	if l[0].y == l[1].y {
		for i := range graph[l[0].y] {
			if l[0].x > l[1].x {
				if l[1].x <= i && i <= l[0].x {
					graph[l[0].y][i]++
				}
			} else {
				if l[0].x <= i && i <= l[1].x {
					graph[l[0].y][i]++
				}
			}
		}
	} else if l[0].x == l[1].x {
		// vertical line
		for i := 0; i <= int(math.Abs(float64(l[0].y-l[1].y))); i++ {
			if l[0].y > l[1].y {
				graph[l[0].y-i][l[0].x]++
			} else {
				graph[l[0].y+i][l[0].x]++
			}
		}
	}
}

type grid [1000][1000]int

func (g *grid) print() {
	for _, v := range g {
		fmt.Println()
		for _, j := range v {
			fmt.Printf("%v", j)
		}
	}
}

func (g *grid) calculate() int {
	var res int
	for _, v := range g {
		for _, j := range v {
			if j > 1 {
				res++
			}
		}
	}

	return res
}

func Part1(input []string) int {
	var graph grid
	var lines []line
	for _, v := range input {
		values := strings.Split(v, " -> ")
		var l line
		for _, j := range values {
			spl := utils.StrList_to_IntList(strings.Split(j, ","))
			p := point{spl[0], spl[1]}
			l = append(l, &p)
		}
		lines = append(lines, l)
	}

	// lines parsed. Now need to draw them all on the graph
	for _, v := range lines {
		v.drawOnGraph(&graph)
	}
	// graph.print()

	return graph.calculate()
}

func Part2(input []string) int {
	var graph grid
	var lines []line
	for _, v := range input {
		values := strings.Split(v, " -> ")
		var l line
		for _, j := range values {
			spl := utils.StrList_to_IntList(strings.Split(j, ","))
			p := point{spl[0], spl[1]}
			l = append(l, &p)
		}
		lines = append(lines, l)
	}

	// lines parsed. Now need to draw them all on the graph
	for _, v := range lines {
		v.drawOnGraph(&graph)
	}
	// graph.print()

	return graph.calculate()
}
