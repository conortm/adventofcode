package main

import (
	"fmt"
	"strings"

	"github.com/conortm/aoc/2021/utils"
)

func getPaperAndInstructions(filename string) ([][]string, []string) {
	lines, _ := utils.GetLinesFromTextFile(filename)
	isCood := true
	xMax := 0
	yMax := 0
	xys := make([]string, 0)
	instructions := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			isCood = false
		} else if isCood {
			xys = append(xys, line)
			xy := strings.Split(line, ",")
			x, _ := utils.StringToInt(xy[0])
			y, _ := utils.StringToInt(xy[1])
			if x > xMax {
				xMax = x
			}
			if y > yMax {
				yMax = y
			}
		} else {
			words := strings.Fields(line)
			instruction := words[len(words)-1]
			instructions = append(instructions, instruction)
		}
	}
	paper := make([][]string, yMax+1)
	for y := 0; y <= yMax; y++ {
		paper[y] = make([]string, xMax+1)
		for x := 0; x <= xMax; x++ {
			paper[y][x] = "."
		}
	}
	for _, xy := range xys {
		xya := strings.Split(xy, ",")
		x, _ := utils.StringToInt(xya[0])
		y, _ := utils.StringToInt(xya[1])
		paper[y][x] = "#"
	}
	// fmt.Printf("xMax: %v, yMax: %v\n", xMax, yMax)
	// fmt.Printf("%v\n", paper)
	return paper, instructions
}

func fold(paper [][]string, axis string, value int) [][]string {
	switch axis {
	case "x":
		for y := 0; y <= len(paper)-1; y++ {
			for x := value + 1; x <= len(paper[y])-1; x++ {
				if paper[y][x] == "#" && value-(x-value) > -1 {
					paper[y][value-(x-value)] = "#"
				}
			}
			paper[y] = paper[y][:value]
		}
	case "y":
		for y := value + 1; y <= len(paper)-1; y++ {
			for x := 0; x <= len(paper[y])-1; x++ {
				if paper[y][x] == "#" && value-(y-value) > -1 {
					paper[value-(y-value)][x] = "#"
				}
			}
		}
		paper = paper[:value]
	}
	return paper
}

func getDotsCount(paper [][]string) int {
	dotsCount := 0
	for y := 0; y < len(paper); y++ {
		for x := 0; x < len(paper[y]); x++ {
			if paper[y][x] == "#" {
				dotsCount++
			}
		}
	}
	return dotsCount
}

func partOne() {
	utils.Intro("PART 1")
	paper, instructions := getPaperAndInstructions("input.txt")
	instruction := instructions[0]
	av := strings.Split(instruction, "=")
	axis := av[0]
	value, _ := utils.StringToInt(av[1])
	paper = fold(paper, axis, value)
	fmt.Println("After 1 Step:")
	// fmt.Printf("paper: %v\n", paper)
	dotsCount := getDotsCount(paper)
	fmt.Printf("dotsCount:%v\n", dotsCount)

}

func partTwo() {
	utils.Intro("PART 2")
	paper, instructions := getPaperAndInstructions("input.txt")
	for i, instruction := range instructions {
		av := strings.Split(instruction, "=")
		axis := av[0]
		value, _ := utils.StringToInt(av[1])
		paper = fold(paper, axis, value)
		fmt.Printf("After %d Step(s):\n", i+1)
		// fmt.Printf("paper: %v\n", paper)
		dotsCount := getDotsCount(paper)
		fmt.Printf("dotsCount:%v\n", dotsCount)
	}
	fmt.Printf("paper:\n%v\n", paper)
}

func main() {
	partOne()
	partTwo()
}
