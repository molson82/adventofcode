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
	var ans int
	almanac := BuildAlmanac(lines)
	fmt.Printf("almanac: %+v", almanac)
	fmt.Println("2023/5 part1 ans: ", ans)
	return ans
}

func Part2() {
	var ans int
	// code here...
	fmt.Println("2023/5 part2 ans: ", ans)
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
