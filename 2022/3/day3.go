package day3

import (
	"unicode"
)

func Part1(input []string) int {
	var sum int
	for _, v := range input {
		part1 := v[0 : len(v)/2]
		part2 := v[len(v)/2:]

		dupe := getDupe(part1, part2)
		priority := getPriority(dupe)
		sum += priority
	}
	return sum
}

func Part2(input []string) int {
	var sum int
	for i := 0; i < len(input)-2; i += 3 {
		group := [3]string{input[i], input[i+1], input[i+2]}
		dupe := getDupeFromGroup(group[0], group[1], group[2])
		priority := getPriority(dupe)
		sum += priority
	}

	return sum
}

func getDupeFromGroup(one, two, three string) rune {
	var res rune
	for _, i := range one {
		for _, j := range two {
			for _, k := range three {
				if i == j && j == k {
					return i
				}
			}
		}
	}
	return res
}

func getDupe(part1, part2 string) rune {
	var res rune
	for _, v := range part1 {
		for _, j := range part2 {
			if v == j {
				return v
			}
		}
	}
	return res
}

func getPriority(dupe rune) int {
	if unicode.IsLower(dupe) {
		return AlphabetPosition(dupe)
	}
	dupe = dupe | ' '
	return AlphabetPosition(dupe) + 26
}

func AlphabetPosition(letter rune) int {
	return int(letter-'a') + 1
}
