package main

import (
	"fmt"
	"strings"

	"github.com/conortm/aoc/2023/utils"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type conversionMap struct {
	name   string
	ranges [][]int
}

func newConversionMap(name string, ranges [][]int) *conversionMap {
	return &conversionMap{name, ranges}
}

func (cm *conversionMap) String() string {
	return fmt.Sprintf("%s map: %v", cm.name, cm.ranges)
}

func getSeedsAndConversionMaps(filename string) ([]int, []*conversionMap) {
	lines, _ := utils.GetLinesFromTextFile(filename)
	seeds, _ := utils.StringToInts(strings.TrimPrefix(lines[0], "seeds: "), " ")
	fmt.Printf("seeds: %v\n", seeds)
	conversionMaps := make([]*conversionMap, 0)
	name := ""
	ranges := make([][]int, 0)
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if line != "" {
			if name == "" {
				name = strings.Split(line, " ")[0]
			} else {
				values, _ := utils.StringToInts(line, " ")
				ranges = append(ranges, values)
			}
		}
		if line == "" || i == len(lines)-1 {
			cm := newConversionMap(name, ranges)
			// fmt.Printf("cm: %v\n", cm)
			conversionMaps = append(conversionMaps, cm)
			name = ""
			ranges = make([][]int, 0)
		}
	}
	return seeds, conversionMaps
}

func getLocation(seed int, conversionMaps []*conversionMap) int {
	value := seed
	// log := fmt.Sprintf("Seed %d", seed)
	for _, conversionMap := range conversionMaps {
		for _, r := range conversionMap.ranges {
			destinationRangeStart := r[0]
			sourceRangeStart := r[1]
			rangeLength := r[2]
			sourceRangeEnd := sourceRangeStart + rangeLength - 1
			if value >= sourceRangeStart && value <= sourceRangeEnd {
				value = destinationRangeStart + (value - sourceRangeStart)
				break
			}
		}
		// log += fmt.Sprintf(", %s %d", strings.Split(conversionMap.name, "-")[2], value)
	}
	// fmt.Printf("%s\n", log)
	return value
}

func getClosestLocation(filename string) int {
	seeds, conversionMaps := getSeedsAndConversionMaps(filename)
	closestLocation := MaxInt
	for _, seed := range seeds {
		location := getLocation(seed, conversionMaps)
		if location < closestLocation {
			closestLocation = location
		}
	}
	return closestLocation
}

func partOne() {
	utils.Intro("PART 1")
	closestLocation := getClosestLocation("test.txt")
	fmt.Printf("closestLocation: %d\n", closestLocation)
}

func partTwo() {
	utils.Intro("PART 2")
}

func main() {
	partOne()
	partTwo()
}
