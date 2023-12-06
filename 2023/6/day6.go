package day6

import (
	"fmt"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

type Race struct {
	Time     int
	Distance int
}

func Part1(lines []string) int {
	ans := 1
	races := GetRaces(lines)
	winList := []int{}
	for _, r := range races {
		// calc # of ways to win
		// fmt.Println("race: ", r)
		var wins int
		for i := 0; i < r.Time; i++ {
			timeLeft := r.Time - i
			distance := timeLeft * i
			if distance > r.Distance {
				// fmt.Printf("Try: %v | Win: %v\n", i, distance)
				wins++
			}
		}
		winList = append(winList, wins)
	}
	for _, w := range winList {
		ans *= w
	}
	fmt.Println("2023/6 part1 ans: ", ans)
	return ans
}

func Part2(lines []string) int {
	var ans int
	// code here...
	fmt.Println("2023/6 part2 ans: ", ans)
	return ans
}

func GetRaces(lines []string) []Race {
	var races []Race
	times := strings.Fields(strings.Split(lines[0], "Time: ")[1])
	distance := strings.Fields(strings.Split(lines[1], "Distance: ")[1])

	for i := 0; i < len(times); i++ {
		newRace := Race{}
		newRace.Time = utils.Str_to_Int(times[i])
		newRace.Distance = utils.Str_to_Int(distance[i])
		races = append(races, newRace)
	}
	return races
}
