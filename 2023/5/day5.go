package day5

import (
	"fmt"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

const (
	seedKey         string = "seeds:"
	seedToSoilKey   string = "seed-to-soil map:"
	soilToFertKey   string = "soil-to-fertilizer map:"
	fertToWaterKey  string = "fertilizer-to-water map:"
	waterToLightKey string = "water-to-light map:"
	lightToTempKey  string = "light-to-temperature map:"
	tempToHumKey    string = "temperature-to-humidity map:"
	humToLocKey     string = "humidity-to-location map:"
)

// Each row of a "Map" contains three numbers:
//   - Destination Range Start
//   - Source Range Start
//   - Range length
//
// Example: In the "seed-to-soil" map with row:
//
//	[50 98 2]
//
// This info means:
//   - source "seed" #98 --> dest "soil" #50
//   - source "seed" #99 --> dest "soil" #51
//
// Because of the range length 2
type Almanac struct {
	Seeds    []int
	SoilMap  [][]int
	FertMap  [][]int
	WaterMap [][]int
	LightMap [][]int
	TempMap  [][]int
	HumMap   [][]int
	LocMap   [][]int
}

func Part1(lines []string) int {
	almanac := BuildAlmanac(lines)

	seedToSoilPlot := BuildPlot(almanac.SoilMap)
	soilToFertPlot := BuildPlot(almanac.FertMap)
	fertToWaterPlot := BuildPlot(almanac.WaterMap)
	waterToLightPlot := BuildPlot(almanac.LightMap)
	lightToTempPlot := BuildPlot(almanac.TempMap)
	tempToHumPlot := BuildPlot(almanac.HumMap)
	humToLocPlot := BuildPlot(almanac.LocMap)

	// get location of each seed
	minLoc := 999999999999
	for _, s := range almanac.Seeds {
		var soil int
		if t, ok := seedToSoilPlot[s]; !ok {
			soil = s
		} else {
			soil = t
		}
		// fmt.Printf("seed: %v | soil: %v\n", s, soil)

		var fert int
		if t, ok := soilToFertPlot[soil]; !ok {
			fert = soil
		} else {
			fert = t
		}
		// fmt.Printf("soil: %v | fert: %v\n", soil, fert)

		var water int
		if t, ok := fertToWaterPlot[fert]; !ok {
			water = fert
		} else {
			water = t
		}
		// fmt.Printf("fert: %v | water: %v\n", fert, water)

		var light int
		if t, ok := waterToLightPlot[water]; !ok {
			light = water
		} else {
			light = t
		}
		// fmt.Printf("water: %v | light: %v\n", water, light)

		var temp int
		if t, ok := lightToTempPlot[light]; !ok {
			temp = light
		} else {
			temp = t
		}
		// fmt.Printf("light: %v | temp: %v\n", light, temp)

		var hum int
		if t, ok := tempToHumPlot[temp]; !ok {
			hum = temp
		} else {
			hum = t
		}
		// fmt.Printf("temp: %v | hum: %v\n", temp, hum)

		var loc int
		if t, ok := humToLocPlot[hum]; !ok {
			loc = hum
		} else {
			loc = t
		}
		// fmt.Printf("hum: %v | loc: %v\n", hum, loc)

		if loc < minLoc {
			minLoc = loc
		}
	}

	fmt.Println("2023/5 part1 ans: ", minLoc)
	return minLoc
}

func Part2() {
	var ans int
	// code here...
	fmt.Println("2023/5 part2 ans: ", ans)
}

func BuildPlot(m [][]int) map[int]int {
	plot := make(map[int]int)
	for _, v := range m {
		destStart := v[0]
		sourceStart := v[1]
		length := v[2]

		plot[sourceStart] = destStart
		for i := 1; i < length; i++ {
			plot[sourceStart+i] = destStart + i
		}
	}

	return plot
}

func BuildAlmanac(lines []string) Almanac {
	var almanac Almanac
	tMap := []string{}
	for i, l := range lines {
		tMap = append(tMap, l)

		// get seeds
		if strings.Contains(l, seedKey) {
			seedStr := strings.Fields(strings.Split(l, seedKey)[1])
			seedList := []int{}
			for _, s := range seedStr {
				seedList = append(seedList, utils.Str_to_Int(s))
			}
			almanac.Seeds = seedList
		}
		// seed to soil map
		if strings.Contains(l, seedToSoilKey) {
			tMap = []string{}
			continue
		}
		// soil to fert map
		if strings.Contains(l, soilToFertKey) {
			// build seed to soil map
			almanac.SoilMap = BuildMap(tMap[:len(tMap)-2])
			tMap = []string{}
			continue
		}
		// fert to water map
		if strings.Contains(l, fertToWaterKey) {
			// build soil to fert map
			almanac.FertMap = BuildMap(tMap[:len(tMap)-2])
			tMap = []string{}
			continue
		}
		// water to light map
		if strings.Contains(l, waterToLightKey) {
			// build fert to water map
			almanac.WaterMap = BuildMap(tMap[:len(tMap)-2])
			tMap = []string{}
			continue
		}
		// light to temp map
		if strings.Contains(l, lightToTempKey) {
			// build water to light
			almanac.LightMap = BuildMap(tMap[:len(tMap)-2])
			tMap = []string{}
			continue
		}
		// temp to hum map
		if strings.Contains(l, tempToHumKey) {
			// build light to temp map
			almanac.TempMap = BuildMap(tMap[:len(tMap)-2])
			tMap = []string{}
			continue
		}
		// hum to loc map
		if strings.Contains(l, humToLocKey) {
			// build temp to hum map
			almanac.HumMap = BuildMap(tMap[:len(tMap)-2])
			tMap = []string{}
			continue
		}
		if i == len(lines)-1 {
			// build hum to loc map
			almanac.LocMap = BuildMap(tMap)
		}
	}
	return almanac
}

func BuildMap(lines []string) [][]int {
	var m [][]int
	for _, l := range lines {
		numStrs := strings.Fields(l)
		tempList := []int{}
		for _, s := range numStrs {
			n := utils.Str_to_Int(s)
			tempList = append(tempList, n)
		}
		m = append(m, tempList)
	}
	return m
}
