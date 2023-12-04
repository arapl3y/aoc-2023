package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/arapl3y/aoc-23/utils"
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
	part2(lines)
}

func part1(lines []string) int {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0

	for _, line := range lines {
		segments := strings.Split(line, ":")
		gameId := parseGameID(segments[0])
		subsets := strings.Split(segments[1], ";")
		possible := true

		for _, subset := range subsets {
			colorValues := extractColorValues(subset)

			if !isSubsetPossible(colorValues, maxCubes) {
				possible = false
				break
			}
		}

		if possible {
			sum += gameId
		}
	}

	fmt.Println(sum)
	return sum
}

func isSubsetPossible(colorValues [][]string, maxCubes map[string]int) bool {
	currentCubes := make(map[string]int)

	for _, colorInfo := range colorValues {
		count := utils.ConvertToInt(colorInfo[1])
		color := colorInfo[2]

		currentCubes[color] += count

		if currentCubes[color] > maxCubes[color] {
			return false
		}
	}
	return true
}

func parseGameID(segment string) int {
	startIndex := strings.IndexFunc(segment, func(r rune) bool {
		return r >= '0' && r <= '9'
	})

	numStr := segment[startIndex:]
	return utils.ConvertToInt(numStr)
}

var colorValueRegex = regexp.MustCompile(`(\d+)\s*([a-z]+)`)

func extractColorValues(subset string) [][]string {
	return colorValueRegex.FindAllStringSubmatch(strings.TrimSpace(subset), -1)
}

func part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		subsets := strings.Split(line, ";")

		minCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, subset := range subsets {
			colorValues := extractColorValues(subset)

			for _, colorInfo := range colorValues {
				count := utils.ConvertToInt(colorInfo[1])
				color := colorInfo[2]

				minCubes[color] = max(minCubes[color], count)
			}
		}

		power := minCubes["red"] * minCubes["green"] * minCubes["blue"]
		sum += power
	}

	fmt.Println(sum)
	return sum
}
