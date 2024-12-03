#!/bin/bash

# Prompt the user for the year and day
read -p "Enter the year: " year
read -p "Enter the day number: " day

# Define file paths
main_file="$year/$day/day$day.go"
test_file="$year/$day/day${day}_test.go"
sample_file="$year/$day/sample.txt"
input_file="$year/$day/input.txt"

# Check if the main file already exists
if [[ -f "$main_file" ]]; then
    echo "Error: File '$main_file' already exists."
    exit 1
fi

# Check if the test file already exists
if [[ -f "$test_file" ]]; then
    echo "Error: File '$test_file' already exists."
    exit 1
fi

# Create the directory structure
mkdir -p "$year/$day"

# Write the main file content
cat <<EOF > "$main_file"
package day$day

import "fmt"

func Part1(input []string) int {
	var ans int
	// code here...
	fmt.Println("$year/$day part1 ans: ", ans)

	return ans
}

func Part2(input []string) int {
	var ans int
	// code here...
	fmt.Println("$year/$day part2 ans: ", ans)

	return ans
}
EOF

# Write the test file content
cat <<EOF > "$test_file"
package day${day}_test

import (
	"testing"

	day${day} "github.com/molson82/adventofcode/${year}/${day}"
	"github.com/molson82/adventofcode/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	sample := utils.ReadInput("./sample.txt")
	ans := day${day}.Part1(sample)
	assert.NotEqual(t, 0, ans)
}

// func TestPart1(t *testing.T) {
// 	input := utils.ReadInput("./input.txt")
// 	ans := day${day}.Part1(input)
// 	assert.Equal(t, 0, ans)
// }
//
// func TestPart2Sample(t *testing.T) {
// 	sample := utils.ReadInput("./sample.txt")
// 	ans := day${day}.Part2(sample)
// 	assert.Equal(t, 0, ans)
// }
//
// func TestPart2(t *testing.T) {
// 	input := utils.ReadInput("./input.txt")
// 	ans := day${day}.Part2(input)
// 	assert.Equal(t, 0, ans)
// }
EOF

# Create empty sample and input files
touch "$sample_file"
touch "$input_file"

# Confirm to the user
echo "Files created in $year/$day/"
