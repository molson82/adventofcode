package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(file string) []map[string]string {
	f, err := os.Open(file)
	check(err)
	defer f.Close()
	var passportList []map[string]string
	passport := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "" {
			lineSplit := strings.Split(scanner.Text(), " ")
			for _, v := range lineSplit {
				keyval := strings.Split(v, ":")
				passport[keyval[0]] = keyval[1]
			}
		} else {
			passportList = append(passportList, passport)
			passport = make(map[string]string)
		}
	}
	passportList = append(passportList, passport)
	return passportList
}

func countValidPassports(passports []map[string]string) int {
	var count int
	for _, v := range passports {
		if _, ok := v["byr"]; ok {
			if _, ok := v["iyr"]; ok {
				if _, ok := v["eyr"]; ok {
					if _, ok := v["hgt"]; ok {
						if _, ok := v["hcl"]; ok {
							if _, ok := v["ecl"]; ok {
								if _, ok := v["pid"]; ok {
									count++
								}
							}
						}
					}
				}
			}
		}
	}
	return count
}

// All of this code is terrible. I'm going for simple not best. Refactor later hopefully
func countValidPart2(passports []map[string]string) int {
	var count int
	for _, v := range passports {
		if val, ok := v["byr"]; ok {
			n, _ := strconv.Atoi(val)
			if len(val) == 4 && n >= 1920 && n <= 2002 {
				if val, ok := v["iyr"]; ok {
					n, _ := strconv.Atoi(val)
					if len(val) == 4 && n >= 2010 && n <= 2020 {
						if val, ok := v["eyr"]; ok {
							n, _ := strconv.Atoi(val)
							if len(val) == 4 && n >= 2020 && n <= 2030 {
								if val, ok := v["hgt"]; ok {
									if strings.Contains(val, "cm") {
										n, _ := strconv.Atoi(strings.Split(val, "cm")[1])
										if n < 150 || n > 193 {
											continue
										}
									} else if strings.Contains(val, "in") {
										n, _ := strconv.Atoi(strings.Split(val, "in")[1])
										if n < 59 || n > 76 {
											continue
										}
									}
									if val, ok := v["hcl"]; ok {
										r, _ := regexp.Compile(`(?m)#([0-9]|[a-f]){6}`)
										if r.MatchString(val) {
											if val, ok := v["ecl"]; ok {
												if val == "amb" || val == "blu" || val == "brn" || val == "gry" ||
													val == "grn" || val == "hzl" || val == "oth" {
													if val, ok := v["pid"]; ok {
														r, _ := regexp.Compile(`(?m)[0-9]{9}`)
														if r.MatchString(val) {
															count++
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return count
}

/*
Part 1
 - attempt 1 = 124 | wrong
	 - I was solving for the wrong thing. Remembered the problem wrong
 - attempt 2 = 182 | correct
Part 2
 - attempt 1 =
*/
func main() {
	passports := readInput("input.txt")
	fmt.Println("Advent of Code day 4")
	fmt.Printf("Day 3 Part 1 ---------------------------------\n")
	fmt.Printf("Total Valid: %v\n", countValidPassports(passports))
	fmt.Printf("\nDay 3 Part 1 ---------------------------------\n")
	fmt.Printf("Total Valid: %v\n", countValidPart2(passports))
}
