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
		var wins int
		for i := 0; i < r.Time; i++ {
			timeLeft := r.Time - i
			distance := timeLeft * i
			if distance > r.Distance {
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
	ans := 1
	timeS := strings.Join(strings.Fields(strings.Split(lines[0], "Time: ")[1]), "")
	distanceS := strings.Join(strings.Fields(strings.Split(lines[1], "Distance: ")[1]), "")

	time := utils.Str_to_Int(timeS)
	distance := utils.Str_to_Int(distanceS)

	fmt.Printf("time: %v | distance: %v\n", time, distance)
	for i := 0; i < time; i++ {
		timeLeft := time - i
		d := timeLeft * i
		if d > distance {
			ans++
		}
	}

	fmt.Println("2023/6 part2 ans: ", ans)
	return ans - 1
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
