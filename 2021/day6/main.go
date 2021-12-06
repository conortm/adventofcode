package main

import (
	"fmt"

	"github.com/conortm/aoc/2021/utils"
)

func getLanternfishInitialState(filename string) []int {
	// Lanternfish have 8 possible states
	lanternfish := make([]int, 9)
	lines, _ := utils.GetLinesFromTextFile(filename)
	initialState, _ := utils.StringToInts(lines[0], ",")
	for i := 0; i < len(lanternfish); i++ {
		lanternfish[i] = 0
	}
	for _, state := range initialState {
		lanternfish[state]++
	}
	return lanternfish
}

func simulateDays(lanternfish []int, numDays int) {
	fmt.Printf("SIMULATING %v DAYS\n", numDays)
	fmt.Printf("Initial State: %+v\n", lanternfish)
	for i := 1; i <= numDays; i++ {
		newLanternfishCount := 0
		for j := 0; j < 9; j++ {
			switch j {
			case 0:
				newLanternfishCount = lanternfish[0]
				lanternfish[0] = 0
			default:
				lanternfish[j-1] = lanternfish[j]
			}
		}
		lanternfish[8] = newLanternfishCount
		lanternfish[6] += newLanternfishCount
		fmt.Printf("After %v day(s): %+v\n", i, lanternfish)
	}
	count := 0
	for i := 0; i < len(lanternfish); i++ {
		count += lanternfish[i]
	}
	fmt.Printf("Count: %v\n", count)
}

func partOne() {
	utils.Intro("PART 1")
	lanternfish := getLanternfishInitialState("input.txt")
	simulateDays(lanternfish, 80)
}

func partTwo() {
	utils.Intro("PART 2")
	lanternfish := getLanternfishInitialState("input.txt")
	simulateDays(lanternfish, 256)
}

func main() {
	partOne()
	partTwo()
}
