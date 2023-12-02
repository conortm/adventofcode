package main

import (
	"fmt"
	"slices"

	"github.com/conortm/aoc/2022/utils"
)

func getElfCalories(lines []string) []int {
	var elfCalories []int
	index := 0
	total := 0
	elfCalories = append(elfCalories, total)
	for _, line := range lines {
		if len(line) > 0 {
			calories, _ := utils.StringToInt(line)
			elfCalories[index] += calories
		} else {
			index++
			total = 0
			elfCalories = append(elfCalories, total)
		}
	}
	return elfCalories
}

func partOne() {
	utils.Intro("PART 1")

	lines, _ := utils.GetLinesFromTextFile("input.txt")
	elfCalories := getElfCalories(lines)

	mostCalroies := 0
	for _, calories := range elfCalories {
		// fmt.Printf("Elf %d total: %d\n", i+1, calories)
		if calories > mostCalroies {
			mostCalroies = calories
		}
	}
	fmt.Printf("mostCalroies: %d\n", mostCalroies)
}

func partTwo() {
	utils.Intro("PART 2")

	lines, _ := utils.GetLinesFromTextFile("input.txt")
	elfCalories := getElfCalories(lines)

	topThreeSum := 0
	slices.Sort(elfCalories)
	for i := len(elfCalories) - 1; i >= len(elfCalories)-3; i-- {
		topThreeSum += elfCalories[i]
	}
	fmt.Printf("topThreeSum: %d\n", topThreeSum)
}

func main() {
	partOne()
	partTwo()
}
