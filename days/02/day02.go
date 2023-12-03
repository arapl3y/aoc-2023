package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("days/02/input/input.txt")

	if err != nil {
		fmt.Println("Error opening the file.", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	part1(lines)
}

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}
