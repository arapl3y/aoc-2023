package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/arapl3y/aoc-23/utils"
)

func main() {
	file, err := os.Open("days/01/input/input.txt")

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

	part1(lines) // also does part 2
}

func part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += extractNums(line)
	}

	fmt.Println(sum)
	return sum
}

func extractNums(line string) int {
	var nums string = ""
	wordMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// Find first
outFirst:
	for i := 0; i < len(line); i++ {
		strChar := string(line[i])

		if utils.IsInt(strChar) {
			nums += strChar
			break
		} else {
			for word := range wordMap {
				// End index is exclusive
				if strings.Contains(line[:i+1], word) {
					nums += wordMap[word]
					break outFirst
				}
			}
		}
	}

	// Find last
outLast:
	for i := len(line) - 1; i >= 0; i-- {
		strChar := string(line[i])

		if utils.IsInt(strChar) {
			nums += strChar
			break
		} else {
			for word := range wordMap {
				if strings.Contains(line[i:], word) {
					nums += wordMap[word]
					break outLast
				}
			}
		}
	}

	return utils.ConvertToInt(nums)
}
