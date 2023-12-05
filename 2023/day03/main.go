package main

import (
	"fmt"

	"github.com/conortm/aoc/2023/utils"
)

type number struct {
	value  int
	xStart int
	xEnd   int
	y      int
}

func newNumber(digits string, xStart int, xEnd int, y int) *number {
	value, _ := utils.StringToInt(digits)
	return &number{value: value, xStart: xStart, xEnd: xEnd, y: y}
}

func (n *number) String() string {
	return fmt.Sprintf("Value %d, xStart: %d, xEnd: %d, y: %d", n.value, n.xStart, n.xEnd, n.y)
}

func getNumbers(filename string) ([][]string, []*number) {
	twoDArray, _ := utils.GetTwoDArrayFromTextFile(filename)
	numbers := make([]*number, 0)
	isDigit := func(s string) bool {
		runeSlice := []rune(s)
		r := runeSlice[0]
		return r >= '0' && r <= '9'
	}
	for y, a := range twoDArray {
		maxX := len(a) - 1
		inNumber := false
		digits := ""
		for x, v := range a {
			var n *number
			if isDigit(v) {
				inNumber = true
				digits += v
				if x == maxX {
					n = newNumber(digits, x-len(digits)+1, x, y)
				}
			} else if inNumber {
				n = newNumber(digits, x-len(digits), x-1, y)
			}
			if n != nil {
				numbers = append(numbers, n)
				inNumber = false
				digits = ""
			}
		}
	}
	return twoDArray, numbers
}

func isSymbol(s string) bool {
	runeSlice := []rune(s)
	r := runeSlice[0]
	return (!(r >= '0' && r <= '9') && r != '.')
}

func hasAdjacentSymbol(twoDArray [][]string, n *number) bool {
	iMin := 0
	iMax := len(twoDArray[0]) - 1
	iStart := n.xStart - 1
	iEnd := n.xEnd + 1
	jMin := 0
	jMax := len(twoDArray) - 1
	jStart := n.y - 1
	jEnd := n.y + 1
	for j := jStart; j <= jEnd; j++ {
		for i := iStart; i <= iEnd; i++ {
			if i >= iMin && i <= iMax && j >= jMin && j <= jMax {
				if isSymbol(twoDArray[j][i]) {
					return true
				}
			}
		}
	}
	return false
}

func getPartNumbers(filename string) ([][]string, []*number) {
	twoDArray, numbers := getNumbers(filename)
	partNumbers := make([]*number, 0)
	for _, number := range numbers {
		if hasAdjacentSymbol(twoDArray, number) {
			// fmt.Printf("partNumber: %v\n", number)
			partNumbers = append(partNumbers, number)
		}
	}
	return twoDArray, partNumbers
}

func partOne() {
	utils.Intro("PART 1")
	_, partNumbers := getPartNumbers("input.txt")
	sumOfPartNumbers := 0
	for _, partNumber := range partNumbers {
		sumOfPartNumbers += partNumber.value
	}
	fmt.Printf("sumOfPartNumbers: %d\n", sumOfPartNumbers)
}

func getAdjacentPartNumbers(x int, y int, partNumbers []*number) []*number {
	adjacentPartNumbers := make([]*number, 0)
	for _, pn := range partNumbers {
		if x >= pn.xStart-1 && x <= pn.xEnd+1 && y >= pn.y-1 && y <= pn.y+1 {
			// fmt.Printf("adjacentPartNumber: %v\n", pn)
			adjacentPartNumbers = append(adjacentPartNumbers, pn)
		}
	}
	return adjacentPartNumbers
}

func partTwo() {
	utils.Intro("PART 2")
	twoDArray, partNumbers := getPartNumbers("input.txt")
	sumOfGearRatios := 0
	for y := 0; y < len(twoDArray); y++ {
		for x := 0; x < len(twoDArray[y]); x++ {
			if twoDArray[y][x] == "*" {
				// fmt.Printf("Gear at (%d, %d)\n", x, y)
				adjacentPartNumbers := getAdjacentPartNumbers(x, y, partNumbers)
				if len(adjacentPartNumbers) == 2 {
					gearRatio := adjacentPartNumbers[0].value * adjacentPartNumbers[1].value
					// fmt.Printf("gearRatio: %d\n", gearRatio)
					sumOfGearRatios += gearRatio
				}
			}
		}
	}
	fmt.Printf("sumOfGearRatios: %d\n", sumOfGearRatios)
}

func main() {
	partOne()
	partTwo()
}
