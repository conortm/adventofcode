package main

import (
	"fmt"
	"sort"

	"github.com/conortm/aoc/2021/utils"
)

type location struct {
	i          int
	j          int
	height     int
	riskLevel  int
	isLowPoint bool
}

func (l *location) String() string {
	return fmt.Sprintf("{i:%d, j:%d, h:%d, lp:%t}", l.i, l.j, l.height, l.isLowPoint)
}

func newLocation(i int, j int, h int) *location {
	return &location{i: i, j: j, height: h, riskLevel: h + 1}
}

type heightmap struct {
	locations             [][]*location
	lowPoints             []*location
	lowPointsRiskLevelSum int
	basins                []map[*location]bool
}

func newHeightmap(filename string) *heightmap {
	twoD, _ := utils.GetTwoDArrayFromTextFile(filename)
	locations := make([][]*location, len(twoD))
	lowPoints := make([]*location, 0)
	lowPointsRiskLevelSum := 0
	iMax := len(twoD) - 1
	jMax := len(twoD[0]) - 1
	for i := 0; i <= iMax; i++ {
		locations[i] = make([]*location, len(twoD[i]))
		for j := 0; j <= jMax; j++ {
			h := twoD[i][j]
			loc := newLocation(i, j, h)
			isLowPoint := true
			// UP
			if i > 0 {
				isLowPoint = (h < twoD[i-1][j])
			}
			// RIGHT
			if isLowPoint && j < jMax {
				isLowPoint = (h < twoD[i][j+1])
			}
			// DOWN
			if isLowPoint && i < iMax {
				isLowPoint = (h < twoD[i+1][j])
			}
			// LEFT
			if isLowPoint && j > 0 {
				isLowPoint = (h < twoD[i][j-1])
			}
			loc.isLowPoint = isLowPoint
			// fmt.Printf("%+v\n", loc)
			if loc.isLowPoint {
				lowPoints = append(lowPoints, loc)
				lowPointsRiskLevelSum += loc.riskLevel
			}
			locations[i][j] = loc
		}
	}
	return &heightmap{locations: locations, lowPoints: lowPoints, lowPointsRiskLevelSum: lowPointsRiskLevelSum}
}

func (hm *heightmap) AddBasinLocations(l *location, basin map[*location]bool) {
	if l.height == 9 {
		return
	}
	if _, isInBasin := basin[l]; isInBasin {
		return
	}
	basin[l] = true
	iMax := len(hm.locations) - 1
	jMax := len(hm.locations[0]) - 1
	// UP
	if l.i > 0 {
		hm.AddBasinLocations(hm.locations[l.i-1][l.j], basin)
	}
	// RIGHT
	if l.j < jMax {
		hm.AddBasinLocations(hm.locations[l.i][l.j+1], basin)
	}
	// DOWN
	if l.i < iMax {
		hm.AddBasinLocations(hm.locations[l.i+1][l.j], basin)
	}
	// LEFT
	if l.j > 0 {
		hm.AddBasinLocations(hm.locations[l.i][l.j-1], basin)
	}
}

func (hm *heightmap) GetBasin(lp *location) map[*location]bool {
	basin := make(map[*location]bool)
	hm.AddBasinLocations(lp, basin)
	return basin
}

func (hm *heightmap) GetBasins() []map[*location]bool {
	if hm.basins != nil {
		return hm.basins
	}
	hm.basins = make([]map[*location]bool, 0)
	for _, lp := range hm.lowPoints {
		// fmt.Println("lp:", lp)
		hm.basins = append(hm.basins, hm.GetBasin(lp))
	}
	return hm.basins
}

func partOne() {
	utils.Intro("PART 1")
	hm := newHeightmap("input.txt")
	fmt.Println("lowPointsRiskLevelSum:", hm.lowPointsRiskLevelSum)
}

func partTwo() {
	utils.Intro("PART 2")
	hm := newHeightmap("input.txt")
	hm.GetBasins()
	// fmt.Println("hm.basins:", hm.basins)
	basinCount := len(hm.basins)
	basinSizes := make([]int, basinCount)
	for i, basin := range hm.basins {
		basinSizes[i] = len(basin)
	}
	sort.Ints(basinSizes)
	// fmt.Println("basinSizes:", basinSizes)
	productOfThreeLargestBasins := basinSizes[basinCount-1] * basinSizes[basinCount-2] * basinSizes[basinCount-3]
	fmt.Println("productOfThreeLargestBasins:", productOfThreeLargestBasins)
}

func main() {
	partOne()
	partTwo()
}
