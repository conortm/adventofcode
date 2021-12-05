package main

import (
	"fmt"
	"strings"

	"github.com/conortm/aoc/2021/utils"
)

type hydrothermalVent struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func getCoordinatesFromPair(pair string) (int, int) {
	xy := strings.Split(pair, ",")
	x, _ := utils.StringToInt(xy[0])
	y, _ := utils.StringToInt(xy[1])
	return x, y
}

func getHydrothermalVents(lines []string) []hydrothermalVent {
	hydrothermalVents := make([]hydrothermalVent, 0)
	for _, line := range lines {
		pairs := strings.Split(line, " -> ")
		x1, y1 := getCoordinatesFromPair(pairs[0])
		x2, y2 := getCoordinatesFromPair(pairs[1])
		hydrothermalVent := hydrothermalVent{x1: x1, y1: y1, x2: x2, y2: y2}
		hydrothermalVents = append(hydrothermalVents, hydrothermalVent)
	}
	return hydrothermalVents
}

func getMaxXY(vents []hydrothermalVent) (int, int) {
	xMax := 0
	yMax := 0
	for _, vent := range vents {
		if vent.x1 > xMax {
			xMax = vent.x1
		}
		if vent.x2 > xMax {
			xMax = vent.x2
		}
		if vent.y1 > yMax {
			yMax = vent.y1
		}
		if vent.y2 > yMax {
			yMax = vent.y2
		}
	}
	return xMax, yMax
}

func ventIsHorizontal(vent hydrothermalVent) bool {
	return vent.y1 == vent.y2
}

func ventIsVertical(vent hydrothermalVent) bool {
	return vent.x1 == vent.x2
}

func getIntsInOrder(first int, second int) (int, int) {
	if first > second {
		return second, first
	}
	return first, second
}

func getOceanFloor(vents []hydrothermalVent, considerDiagonal bool) [][]int {
	xMax, yMax := getMaxXY(vents)
	oceanFloor := make([][]int, xMax+1)
	for x := 0; x <= xMax; x++ {
		oceanFloor[x] = make([]int, yMax+1)
		for y := 0; y <= yMax; y++ {
			oceanFloor[x][y] = 0
		}
	}
	for _, vent := range vents {
		if ventIsHorizontal(vent) {
			xLeft, xRight := getIntsInOrder(vent.x1, vent.x2)
			for x := xLeft; x <= xRight; x++ {
				oceanFloor[x][vent.y1]++
			}
		} else if ventIsVertical(vent) {
			yTop, yBot := getIntsInOrder(vent.y1, vent.y2)
			for y := yTop; y <= yBot; y++ {
				oceanFloor[vent.x1][y]++
			}
		} else if considerDiagonal {
			if vent.x1 < vent.x2 {
				len := vent.x2 - vent.x1
				if vent.y1 < vent.y2 {
					// left to right, up to down (SE)
					for i := 0; i <= len; i++ {
						oceanFloor[vent.x1+i][vent.y1+i]++
					}
				} else {
					// left to right, down to up (NE)
					for i := 0; i <= len; i++ {
						oceanFloor[vent.x1+i][vent.y1-i]++
					}
				}
			} else {
				len := vent.x1 - vent.x2
				if vent.y1 < vent.y2 {
					// right to left, up to down (SW)
					for i := 0; i <= len; i++ {
						oceanFloor[vent.x1-i][vent.y1+i]++
					}
				} else {
					// right to left, down to up (NW)
					for i := 0; i <= len; i++ {
						oceanFloor[vent.x1-i][vent.y1-i]++
					}
				}
			}
		}
	}
	return oceanFloor
}

func getDangerousCoordinateCount(oceanFloor [][]int) int {
	dangerousCoordinateCount := 0
	threshold := 2
	for x := 0; x < len(oceanFloor); x++ {
		for y := 0; y < len(oceanFloor[x]); y++ {
			if oceanFloor[x][y] >= threshold {
				dangerousCoordinateCount++
			}
		}
	}
	return dangerousCoordinateCount
}

func partOne() {
	utils.Intro("PART 1")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	hydrothermalVents := getHydrothermalVents(lines)
	// fmt.Println("len hydrothermalVents:", len(hydrothermalVents))
	// fmt.Printf("hydrothermalVents: %+v\n", hydrothermalVents)
	oceanFloor := getOceanFloor(hydrothermalVents, false)
	// fmt.Printf("oceanFloor: %+v\n", oceanFloor)
	dangerousCoordinateCount := getDangerousCoordinateCount(oceanFloor)
	fmt.Println("dangerousCoordinateCount:", dangerousCoordinateCount)
}

func partTwo() {
	utils.Intro("PART 2")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	hydrothermalVents := getHydrothermalVents(lines)
	// fmt.Println("len hydrothermalVents:", len(hydrothermalVents))
	// fmt.Printf("hydrothermalVents: %+v\n", hydrothermalVents)
	oceanFloor := getOceanFloor(hydrothermalVents, true)
	// fmt.Printf("oceanFloor: %+v\n", oceanFloor)
	dangerousCoordinateCount := getDangerousCoordinateCount(oceanFloor)
	fmt.Println("dangerousCoordinateCount:", dangerousCoordinateCount)
}

func main() {
	partOne()
	partTwo()
}
