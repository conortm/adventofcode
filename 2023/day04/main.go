package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/conortm/aoc/2023/utils"
)

type scratchcard struct {
	id              int
	winningNumbers  []int
	numbersYouHave  []int
	matchingNumbers []int
	instances       int
}

func newScratchcard(line string) *scratchcard {
	lineParts := strings.Split(line, ": ")
	idString := strings.TrimLeft(lineParts[0], "Card ")
	id, _ := utils.StringToInt(idString)
	numbersParts := strings.Split(lineParts[1], "|")
	getIntsFromNumberString := func(numberString string) ([]int, error) {
		fNumberString := strings.ReplaceAll(numberString, "  ", " ")
		fNumberString = strings.Trim(fNumberString, " ")
		return utils.StringToInts(fNumberString, " ")
	}
	winningNumbers, _ := getIntsFromNumberString(numbersParts[0])
	numbersYouHave, _ := getIntsFromNumberString(numbersParts[1])
	matchingNumbers := make([]int, 0)
	for _, winningNumber := range winningNumbers {
		if slices.Contains(numbersYouHave, winningNumber) {
			matchingNumbers = append(matchingNumbers, winningNumber)
		}
	}
	return &scratchcard{id: id, winningNumbers: winningNumbers, numbersYouHave: numbersYouHave, matchingNumbers: matchingNumbers, instances: 1}
}

func (sc *scratchcard) String() string {
	return fmt.Sprintf("Card %d: winningNumbers: %v; numbersYouHave: %v, matchingNumbers: %v, instances: %d", sc.id, sc.winningNumbers, sc.numbersYouHave, sc.matchingNumbers, sc.instances)
}

func getScratchcardsFromLines(lines []string) []*scratchcard {
	scratchcards := make([]*scratchcard, 0)
	for _, line := range lines {
		scratchcard := newScratchcard(line)
		// fmt.Printf("%s\n", scratchcard)
		scratchcards = append(scratchcards, scratchcard)
	}
	return scratchcards
}

func partOne() {
	utils.Intro("PART 1")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	scratchcards := getScratchcardsFromLines(lines)
	totalPoints := 0
	for _, scratchcard := range scratchcards {
		points := 0
		for i := 0; i < len(scratchcard.matchingNumbers); i++ {
			if points == 0 {
				points = 1
			} else {
				points = points * 2
			}
		}
		totalPoints += points
	}
	fmt.Printf("totalPoints: %d\n", totalPoints)
}

func partTwo() {
	utils.Intro("PART 2")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	scratchcards := getScratchcardsFromLines(lines)
	for i, _ := range scratchcards {
		scratchcard := scratchcards[i]
		matchingNumbersCount := len(scratchcard.matchingNumbers)
		if matchingNumbersCount > 0 {
			for j := 0; j < scratchcard.instances; j++ {
				for k := i + 1; k <= i+matchingNumbersCount; k++ {
					scratchcards[k].instances++
				}
			}
		}
	}
	totalScratchCards := 0
	for _, scratchcard := range scratchcards {
		totalScratchCards += scratchcard.instances
	}
	fmt.Printf("totalScratchCards: %d\n", totalScratchCards)
}

func main() {
	partOne()
	partTwo()
}
