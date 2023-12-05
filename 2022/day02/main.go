package main

import (
	"fmt"
	"strings"

	"github.com/conortm/aoc/2022/utils"
)

func partOne() {
	utils.Intro("PART 1")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	totalScore := 0
	shapeScoreMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	shapeSameMap := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	shapeWinMap := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	for _, line := range lines {
		shapes := strings.Split(line, " ")
		opponentShape := shapes[0]
		myShape := shapes[1]
		score := shapeScoreMap[myShape]
		if shapeSameMap[opponentShape] == myShape {
			score += 3
		} else if shapeWinMap[opponentShape] == myShape {
			score += 6
		}
		totalScore += score
	}
	fmt.Printf("totalScore: %d\n", totalScore)
}

func partTwo() {
	utils.Intro("PART 2")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	totalScore := 0
	shapeScoreMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	shapeLoseMap := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}
	shapeWinMap := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}
	for _, line := range lines {
		shapes := strings.Split(line, " ")
		opponentShape := shapes[0]
		myResponse := shapes[1]
		var myShape string
		score := 0
		switch myResponse {
		case "X": // lose
			myShape = shapeLoseMap[opponentShape]
		case "Y": // draw
			myShape = opponentShape
			score += 3
		case "Z": // win
			myShape = shapeWinMap[opponentShape]
			score += 6
		}
		score += shapeScoreMap[myShape]
		totalScore += score
	}
	fmt.Printf("totalScore: %d\n", totalScore)
}

func main() {
	partOne()
	partTwo()
}
