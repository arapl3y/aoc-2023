package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("days/03/input/input.txt")

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

	// part1(lines)
	part2(lines)
}

func part1(lines []string) int {
	sum := 0
	partNums := getPartNums(lines)

	for _, value := range partNums {
		sum += value
	}

	fmt.Println(sum)
	return sum
}

func part2(lines []string) int {
	sum := 0
	nums := getGearPartNums(lines)
	gearRatios := []int{}

	for _, slice := range nums {
		// Only want to acknowledge as gear ratio if length of 2
		if len(slice) == 2 {
			gearRatio := slice[0] * slice[1]
			gearRatios = append(gearRatios, gearRatio)
		}
	}

	for _, value := range gearRatios {
		sum += value
	}

	fmt.Println(sum)
	return sum
}

func getGearPartNums(lines []string) map[string][]int {
	numsMap := make(map[string][]int)

	for lineIndex, line := range lines {
		for runeIndex, r := range line {
			if isGearSymbol(r) {
				checkGearAdjacent(r, lines, lineIndex, runeIndex, numsMap)
			}
		}
	}

	return numsMap
}

func checkGearAdjacent(r rune, lines []string, lineIndex int, runeIndex int, numsMap map[string][]int) {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	for _, d := range directions {
		newLineIndex := lineIndex + d[0]
		newRuneIndex := runeIndex + d[1]

		if newRuneIndex >= 0 && newRuneIndex < len(lines[lineIndex]) && newLineIndex >= 0 && newLineIndex < len(lines) {
			newRune := rune(lines[newLineIndex][newRuneIndex])

			if unicode.IsDigit(newRune) {
				num, _ := findWholeNum(newLineIndex, newRuneIndex, lines[newLineIndex])
				key := fmt.Sprintf("%d,%d", lineIndex, runeIndex)

				if !slices.Contains(numsMap[key], num) {
					numsMap[key] = append(numsMap[key], num)
				}
			}
		}
	}
}

func isGearSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r == '*'
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func getPartNums(lines []string) map[string]int {
	numsMap := make(map[string]int)

	for lineIndex, line := range lines {
		for runeIndex, r := range line {
			if isSymbol(r) {
				checkAdjacent(r, lines, lineIndex, runeIndex, numsMap)
			}
		}
	}

	return numsMap
}

func checkAdjacent(r rune, lines []string, lineIndex int, runeIndex int, numsMap map[string]int) {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	for _, d := range directions {
		newLineIndex := lineIndex + d[0]
		newRuneIndex := runeIndex + d[1]

		if newRuneIndex >= 0 && newRuneIndex < len(lines[lineIndex]) && newLineIndex >= 0 && newLineIndex < len(lines) {
			newRune := rune(lines[newLineIndex][newRuneIndex])

			if unicode.IsDigit(newRune) {
				num, startIndex := findWholeNum(newLineIndex, newRuneIndex, lines[newLineIndex])
				key := fmt.Sprintf("%d,%d", newLineIndex, startIndex)
				numsMap[key] = num
			}
		}
	}
}

func findWholeNum(lineIndex int, runeIndex int, line string) (int, int) {
	startIndex := runeIndex
	strNum := ""

	for i := runeIndex; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			startIndex = i
		} else {
			break
		}
	}

	for i := startIndex; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			strNum += string(line[i])
		} else {
			break
		}
	}

	fullNum, _ := strconv.Atoi(strNum)
	return fullNum, startIndex
}
