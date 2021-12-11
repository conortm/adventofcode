package main

import (
	"fmt"
	"sort"

	"github.com/conortm/aoc/2021/utils"
)

func getExpectedCloser(opener string) string {
	switch opener {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	}
	return ""
}

func processLine(line string) (bool, utils.Stack, string) {
	var expectedClosers utils.Stack
	for _, c := range line {
		char := string(c)
		expectedCloser := getExpectedCloser(char)
		if expectedCloser != "" {
			expectedClosers.Push(expectedCloser)
		} else {
			expectedCloser, _ = expectedClosers.Pop()
			if char != expectedCloser {
				return false, expectedClosers, char
			}
		}
	}
	return true, expectedClosers, ""
}

func partOne() {
	utils.Intro("PART 1")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	illegalCharCount := make(map[string]int, 4)
	illegalCharCount[")"] = 0
	illegalCharCount["]"] = 0
	illegalCharCount["}"] = 0
	illegalCharCount[">"] = 0
	for _, line := range lines {
		isValid, _, char := processLine(line)
		if !isValid {
			illegalCharCount[char]++
		}
	}
	// fmt.Printf("illegalCharCount: %v\n", illegalCharCount)
	totalScore := 0
	for char, count := range illegalCharCount {
		switch char {
		case ")":
			totalScore += count * 3
		case "]":
			totalScore += count * 57
		case "}":
			totalScore += count * 1197
		case ">":
			totalScore += count * 25137
		}
	}
	fmt.Println("totalScore:", totalScore)
}

func partTwo() {
	utils.Intro("PART 2")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	scores := make([]int, 0)
	for _, line := range lines {
		isValid, expectedClosers, _ := processLine(line)
		if isValid {
			score := 0
			for {
				expectedCloser, isMore := expectedClosers.Pop()
				if !isMore {
					break
				}
				score = 5 * score
				switch expectedCloser {
				case ")":
					score += 1
				case "]":
					score += 2
				case "}":
					score += 3
				case ">":
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	middleScore := scores[int(len(scores)/2)]
	fmt.Println("middleScore:", middleScore)
}

func main() {
	partOne()
	partTwo()
}
