package main

import (
	"fmt"
	"log"

	"github.com/conortm/aoc/2021/utils"
)

func getIncreaseCount(values []int) int {
	inreaseCount := 0

	var prevValue int

	for i, value := range values {
		if i > 0 && value > prevValue {
			// fmt.Println("INCREASE!", value, "is greater than", prevValue)
			inreaseCount++
		} else {
			// fmt.Println("DECREASE!", value, "is NOT greater than", prevValue)
		}
		prevValue = value
	}

	return inreaseCount
}

func getWindowDepths(depths []int) []int {
	windowDepths := make([]int, 0)

	for i, depth := range depths[:(len(depths) - 2)] {
		windowDepth := depth + depths[i+1] + depths[i+2]
		windowDepths = append(windowDepths, windowDepth)
	}

	return windowDepths
}

func main() {
	depths, err := utils.GetIntValsFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("depthsCount:", len(depths))
	// Part 1
	depthInreaseCount := getIncreaseCount(depths)
	fmt.Println("depthInreaseCount:", depthInreaseCount)
	// Part 2
	windowDepths := getWindowDepths(depths)
	fmt.Println("windowDepthsCount:", len(windowDepths))
	windowDepthIncreaseCount := getIncreaseCount(windowDepths)
	fmt.Println("windowDepthIncreaseCount:", windowDepthIncreaseCount)
}
