package day2

import (
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

type direction struct {
	dir   string
	count int
}

func Part1(input []string) int {
	var dirs []direction
	// convert each line to a list of structs
	for _, v := range input {
		s := strings.Split(v, " ")
		c := utils.Str_to_Int(s[1])
		dir := direction{s[0], c}
		dirs = append(dirs, dir)
	}

	// do calculation
	var depth int
	var horiz int
	for _, v := range dirs {
		if v.dir == "up" {
			depth -= v.count
		} else if v.dir == "down" {
			depth += v.count
		} else {
			horiz += v.count
		}
	}

	return depth * horiz
}

func Part2(input []string) int {
	var dirs []direction
	// convert each line to a list of structs
	for _, v := range input {
		s := strings.Split(v, " ")
		c := utils.Str_to_Int(s[1])
		dir := direction{s[0], c}
		dirs = append(dirs, dir)
	}

	// do calculation
	var depth int
	var horiz int
	var aim int
	for _, v := range dirs {
		if v.dir == "up" {
			aim -= v.count
		} else if v.dir == "down" {
			aim += v.count
		} else {
			horiz += v.count
			depth += aim * v.count
		}
	}

	return depth * horiz
}
