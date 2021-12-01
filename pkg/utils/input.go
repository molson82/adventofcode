package utils

import (
	"bufio"
	"os"
)

func ReadInput(file string) []string {
	f, err := os.Open(file)
	CheckErr(err)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
