package main

import (
	"fmt"
	"strings"

	"github.com/conortm/aoc/2021/utils"
)

func getNewPolymer(polymer string, pairInsertionRules map[string]string) string {
	newPolymer := ""
	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		insert := pairInsertionRules[pair]
		newPair := pair
		if insert != "" {
			newPair = ""
			if i == 0 {
				newPair = string(pair[0])
			}
			newPair += insert + string(pair[1])
		}
		newPolymer += newPair
	}
	return newPolymer
}

func getMaxMinDifference(filename string, steps int) int {
	lines, _ := utils.GetLinesFromTextFile(filename)
	polymerTemplate := lines[0]
	// fmt.Printf("Template: %s\n", polymerTemplate)
	pairInsertionRules := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " -> ")
		pairInsertionRules[parts[0]] = parts[1]
	}
	polymer := polymerTemplate
	for s := 1; s <= steps; s++ {
		polymer = getNewPolymer(polymer, pairInsertionRules)
		// fmt.Printf("After step %d: %s\n", s, polymer)
	}
	quantities := make(map[string]int, 0)
	for i := 0; i < len(polymer); i++ {
		quantities[string(polymer[i])]++
	}
	// fmt.Printf("quantities: %v\n", quantities)
	maxQuantity := 0
	minQuantity := 0
	for _, quantity := range quantities {
		if quantity > maxQuantity {
			maxQuantity = quantity
		}
		if minQuantity == 0 || quantity < minQuantity {
			minQuantity = quantity
		}
	}
	// fmt.Printf("maxQuantity: %d, minQuantity: %d\n", maxQuantity, minQuantity)
	maxMinDifference := maxQuantity - minQuantity
	return maxMinDifference
}

func partOne() {
	utils.Intro("PART 1")
	maxMinDifference := getMaxMinDifference("input.txt", 10)
	fmt.Printf("maxMinDifference: %d\n", maxMinDifference)
}

func partTwo() {
	utils.Intro("PART 2")
	// maxMinDifference := getMaxMinDifference("test.txt", 40)
	// fmt.Printf("maxMinDifference: %d\n", maxMinDifference)
}

func main() {
	partOne()
	partTwo()
}
