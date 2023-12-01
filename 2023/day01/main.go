package main

import (
	"fmt"
	"strings"

	"github.com/conortm/aoc/2023/utils"
)

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func partOne() {
	utils.Intro("PART 1")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	total := 0
	for _, line := range lines {
		var b strings.Builder
		b.WriteByte(line[strings.IndexFunc(line, isDigit)])
		b.WriteByte(line[strings.LastIndexFunc(line, isDigit)])
		number, _ := utils.StringToInt(b.String())
		// fmt.Printf("line: %s; number: %d\n", line, number)
		total += number
	}
	fmt.Printf("total: %d\n", total)
}

func getDigit(line string, i int) int {
	r := rune(line[i])
	if isDigit(r) {
		digit, _ := utils.StringToInt(string(r))
		return digit
	}
	lineLength := len(line)
	digits := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for word, digit := range digits {
		if lineLength-i >= len(word) && line[i:i+len(word)] == word {
			return digit
		}
	}
	return -1
}

func getFirstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		digit := getDigit(line, i)
		if digit > -1 {
			return digit
		}
	}
	return -1
}

func getLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		digit := getDigit(line, i)
		if digit > -1 {
			return digit
		}
	}
	return -1
}

func partTwo() {
	utils.Intro("PART 2")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	total := 0
	for _, line := range lines {
		firstDigit := getFirstDigit(line)
		lastDigit := getLastDigit(line)
		number := 10*firstDigit + lastDigit
		// fmt.Printf("line: %s; number: %d\n", line, number)
		total += number
	}
	fmt.Printf("total: %d\n", total)
}

func main() {
	partOne()
	partTwo()
}
