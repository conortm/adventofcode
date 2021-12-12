package main

import (
	"fmt"

	"github.com/conortm/aoc/2021/utils"
)

type octopus struct {
	energyLevel int
	isFlashed   bool
	i           int
	j           int
}

// Increments energyLevel of octopus, and returns if it flashes.
func (o *octopus) incrementEnergyLevel() bool {
	flashHappened := false
	o.energyLevel++
	if o.energyLevel > 9 {
		if !o.isFlashed {
			flashHappened = true
			o.isFlashed = true
		}
		o.energyLevel = 0
	}
	return flashHappened
}

func newOctopus(initialEnergyLevel int, i int, j int) *octopus {
	return &octopus{energyLevel: initialEnergyLevel, isFlashed: false, i: i, j: j}
}

type octopusGrid struct {
	grid [][]*octopus
}

func newOctopusGrid(filename string) *octopusGrid {
	twoD, _ := utils.GetTwoDArrayFromTextFile(filename)
	grid := make([][]*octopus, len(twoD))
	for i := 0; i < len(twoD); i++ {
		grid[i] = make([]*octopus, len(twoD[i]))
		for j := 0; j < len(twoD[i]); j++ {
			oct := newOctopus(twoD[i][j], i, j)
			grid[i][j] = oct
		}
	}
	return &octopusGrid{grid: grid}
}

func (og *octopusGrid) String() string {
	output := ""
	for i := 0; i < len(og.grid); i++ {
		line := ""
		for j := 0; j < len(og.grid[i]); j++ {
			line += fmt.Sprintf("%d", og.grid[i][j].energyLevel)
		}
		output += line + "\n"
	}
	return output
}

func (og *octopusGrid) Step() (int, bool) {
	flashCount := 0
	flashedOctopi := make([]*octopus, 0)
	iMin := 0
	iMax := len(og.grid) - 1
	jMin := 0
	jMax := len(og.grid[0]) - 1
	for i := iMin; i <= iMax; i++ {
		for j := jMin; j <= jMax; j++ {
			if og.grid[i][j].incrementEnergyLevel() {
				flashedOctopi = append(flashedOctopi, og.grid[i][j])
			}
		}
	}
	for {
		if len(flashedOctopi) == 0 {
			break
		}
		flashCount++
		o := flashedOctopi[0]
		flashedOctopi = flashedOctopi[1:]
		iUp := o.i - 1
		if iUp < iMin {
			iUp = iMin
		}
		iDown := o.i + 1
		if iDown > iMax {
			iDown = iMax
		}
		jLeft := o.j - 1
		if jLeft < jMin {
			jLeft = jMin
		}
		jRight := o.j + 1
		if jRight > jMax {
			jRight = jMax
		}
		for i := iUp; i <= iDown; i++ {
			for j := jLeft; j <= jRight; j++ {
				if !(i == o.i && j == o.j) {
					if og.grid[i][j].incrementEnergyLevel() {
						flashedOctopi = append(flashedOctopi, og.grid[i][j])
					}
				}
			}
		}
	}
	// Reset flashes
	allFlashed := true
	for i := 0; i < len(og.grid); i++ {
		for j := 0; j < len(og.grid[i]); j++ {
			if og.grid[i][j].isFlashed {
				og.grid[i][j].energyLevel = 0
				og.grid[i][j].isFlashed = false
			} else {
				allFlashed = false
			}
		}
	}
	return flashCount, allFlashed
}

func partOne() {
	utils.Intro("PART 1")
	totalFlashCount := 0
	octopusGrid := newOctopusGrid("input.txt")
	// fmt.Printf("Before any steps:\n%v\n", octopusGrid)
	for i := 1; i <= 100; i++ {
		flashCount, _ := octopusGrid.Step()
		totalFlashCount += flashCount
		// fmt.Printf("After step %d:\n%v\n", i, octopusGrid)
	}
	fmt.Println("totalFlashCount:", totalFlashCount)
}

func partTwo() {
	utils.Intro("PART 2")
	octopusGrid := newOctopusGrid("input.txt")
	// fmt.Printf("Before any steps:\n%v\n", octopusGrid)
	i := 1
	for {
		_, allFlashed := octopusGrid.Step()
		// fmt.Printf("After step %d:\n%v\n", i, octopusGrid)
		if allFlashed {
			break
		}
		i++
	}
	fmt.Println("First step during which all octopuses flash:", i)
}

func main() {
	partOne()
	partTwo()
}
