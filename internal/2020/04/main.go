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

func validateByrValue(v string) bool {
	i, _ := strconv.Atoi(v)

	return i >= 1920 && i <= 2002
}

func validateIyrValue(v string) bool {
	i, _ := strconv.Atoi(v)

	return i >= 2010 && i <= 2020
}

func validateEyrValue(v string) bool {
	i, _ := strconv.Atoi(v)

	return i >= 2020 && i <= 2030
}

func validateHgtValue(v string) bool {
	if strings.HasSuffix(v, "in") {
		splitOnIn := strings.Split(v, "in")
		i, _ := strconv.Atoi(splitOnIn[0])
		return i >= 59 && i <= 76
	}
	if strings.HasSuffix(v, "cm") {
		splitOnIn := strings.Split(v, "cm")
		i, _ := strconv.Atoi(splitOnIn[0])
		return i >= 150 && i <= 193
	}
	return false
}

func validateHclValue(v string) bool {
	valid, _ := regexp.MatchString("^#[0-9a-f]{6}", v)
	return valid
}

func validateEclValue(v string) bool {
	return v == "amb" || v == "blu" || v == "brn" || v == "gry" || v == "grn" || v == "hzl" || v == "oth"
}

func validatePidValue(v string) bool {
	valid, _ := regexp.MatchString("^[0-9]{9}", v)
	return valid && len(v) == 9
}

func areValuesValid(fieldToValue map[string]string) bool {
	for field, value := range fieldToValue {
		if field == "byr" && !validateByrValue(value) {
			return false
		}
		if field == "iyr" && !validateIyrValue(value) {
			return false
		}
		if field == "eyr" && !validateEyrValue(value) {
			return false
		}
		if field == "hgt" && !validateHgtValue(value) {
			return false
		}
		if field == "hcl" && !validateHclValue(value) {
			return false
		}
		if field == "ecl" && !validateEclValue(value) {
			return false
		}
		if field == "pid" && !validatePidValue(value) {
			return false
		}
	}
	return true
}

func areFieldsValid(v map[string]string) bool {
	if _, ok := v["byr"]; ok {
		if _, ok := v["iyr"]; ok {
			if _, ok := v["eyr"]; ok {
				if _, ok := v["hgt"]; ok {
					if _, ok := v["hcl"]; ok {
						if _, ok := v["ecl"]; ok {
							if _, ok := v["pid"]; ok {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

// All of this code is terrible. I'm going for simple not best. Refactor later hopefully
func countValidPart2(passports []map[string]string) int {
	var count int
	for _, v := range passports {
		if areValuesValid(v) && areFieldsValid(v) {
			count++
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
 - attempt 1 = 1 | wrong
		 - no surprise there. Made the code way to complicated. Need to simplify it
 - attempt 2 = 109 | correct
*/
func main() {
	passports := readInput("input.txt")
	fmt.Println("Advent of Code day 4")
	fmt.Printf("Day 4 Part 1 ---------------------------------\n")
	fmt.Printf("Total Valid: %v\n", countValidPassports(passports))
	fmt.Printf("\nDay 4 Part 2 ---------------------------------\n")
	fmt.Printf("Total Valid: %v\n", countValidPart2(passports))
}
