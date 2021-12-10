package main

import (
	"fmt"
	"sort"

	"github.com/conortm/aoc/2021/utils"
)

func getPositions(filename string) []int {
	lines, _ := utils.GetLinesFromTextFile(filename)
	positions, _ := utils.StringToInts(lines[0], ",")
	sort.Ints(positions)
	return positions
}

func getTotalFuel(positions []int, destination int, constantRate bool, distanceToFuel map[int]int) int {
	totalFuel := 0
	for _, position := range positions {
		distance := utils.Abs(position - destination)
		if constantRate {
			totalFuel += distance
		} else {
			fuel, exists := distanceToFuel[distance]
			if !exists {
				fuel = 0
				for i := 1; i <= distance; i++ {
					fuel += i
				}
				distanceToFuel[distance] = fuel
			}
			totalFuel += fuel
		}
	}
	return totalFuel
}

func getDestinationWithMinFuel(positions []int, constantRate bool) (int, int) {
	minDestination := -1
	minFuel := -1
	distanceToFuel := make(map[int]int)
	for destination := positions[0]; destination <= positions[len(positions)-1]; destination++ {
		fuel := getTotalFuel(positions, destination, constantRate, distanceToFuel)
		if minFuel == -1 || fuel < minFuel {
			minFuel = fuel
			minDestination = destination
		}
	}
	return minDestination, minFuel
}

func partOne() {
	utils.Intro("PART 1")
	positions := getPositions("input.txt")
	fmt.Printf("positions: %v\n", positions)
	destination, fuel := getDestinationWithMinFuel(positions, true)
	fmt.Printf("destination: %v, fuel: %v\n", destination, fuel)
}

func partTwo() {
	utils.Intro("PART 2")
	positions := getPositions("input.txt")
	fmt.Printf("positions: %v\n", positions)
	destination, fuel := getDestinationWithMinFuel(positions, false)
	fmt.Printf("destination: %v, fuel: %v\n", destination, fuel)
}

func main() {
	partOne()
	partTwo()
}
