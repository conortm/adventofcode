package main

import (
	"fmt"
	"strings"

	"github.com/conortm/aoc/2021/utils"
)

func getBingoNumbers(lines []string) []string {
	return strings.Split(lines[0], ",")
}

type bingoBoardNumber struct {
	value    string
	isMarked bool
}

func (bbn *bingoBoardNumber) String() string {
	return fmt.Sprintf("{value:%s, isMarked:%t}", bbn.value, bbn.isMarked)
}

func (bbn *bingoBoardNumber) Mark(value string) bool {
	if !bbn.isMarked && bbn.value == value {
		bbn.isMarked = true
		return true
	}
	return false
}

func newBingoBoardNumber(value string) *bingoBoardNumber {
	return &bingoBoardNumber{value: value, isMarked: false}
}

type bingoBoard struct {
	grid     [5][5]*bingoBoardNumber
	isWinner bool
}

func (bb *bingoBoard) String() string {
	rowStrings := make([]string, 0)
	for _, row := range bb.grid {
		rowStrings = append(rowStrings, fmt.Sprintf("row: %v\n", row))
	}
	return fmt.Sprintf("{grid:%v}\n", rowStrings)
}

func (bb *bingoBoard) Mark(value string) bool {
	isChanged := false
	if !bb.isWinner {
		for i := 0; i < len(bb.grid); i++ {
			for j := 0; j < len(bb.grid[i]); j++ {
				if bb.grid[i][j].Mark(value) {
					isChanged = true
				}
			}
		}
	}
	return isChanged
}

func (bb *bingoBoard) IsWinner() bool {
	if bb.isWinner {
		return true
	}
	for _, row := range bb.grid {
		// TODO: make this dynamic
		if row[0].isMarked && row[1].isMarked && row[2].isMarked && row[3].isMarked && row[4].isMarked {
			bb.isWinner = true
			return true
		}
	}
	for i := 0; i < len(bb.grid); i++ {
		// TODO: make this dynamic
		if bb.grid[0][i].isMarked && bb.grid[1][i].isMarked && bb.grid[2][i].isMarked && bb.grid[3][i].isMarked && bb.grid[4][i].isMarked {
			bb.isWinner = true
			return true
		}
	}
	return false
}

func (bb *bingoBoard) GetSumOfUnmarkedNumbers() int {
	sumOfUnmarkedNumbers := 0
	for i := 0; i < len(bb.grid); i++ {
		for j := 0; j < len(bb.grid[i]); j++ {
			if !bb.grid[i][j].isMarked {
				intVal, _ := utils.StringToInt(bb.grid[i][j].value)
				sumOfUnmarkedNumbers += intVal
			}
		}
	}
	return sumOfUnmarkedNumbers
}

func newBingoBoard(lines []string) *bingoBoard {
	var grid [5][5]*bingoBoardNumber
	for i, line := range lines {
		for j, value := range strings.Fields(line) {
			grid[i][j] = newBingoBoardNumber(value)
		}
	}
	return &bingoBoard{grid: grid, isWinner: false}
}

func getBingoBoards(lines []string) []*bingoBoard {
	bingoBoards := make([]*bingoBoard, 0)
	rowsPerBingoBoard := 5
	for i := 2; i < len(lines); i += rowsPerBingoBoard + 1 {
		bingoBoard := newBingoBoard(lines[i:(i + rowsPerBingoBoard)])
		bingoBoards = append(bingoBoards, bingoBoard)
	}
	return bingoBoards
}

// Returns index and last number called of first & last winning boards (else -1, nil)
func playBingo(numbers []string, boards []*bingoBoard) (int, string, int, string) {
	// TODO: Keep track of turns/marks/winning boards instead of just first/last
	indexOfFirstWinningBoard := -1
	winningNumberOfFirstWinningBoard := ""
	indexOfLastWinningBoard := -1
	winningNumberOfLastWinningBoard := ""
	for _, number := range numbers {
		for i := 0; i < len(boards); i++ {
			if !boards[i].IsWinner() && boards[i].Mark(number) && boards[i].IsWinner() {
				if indexOfFirstWinningBoard == -1 {
					indexOfFirstWinningBoard = i
					winningNumberOfFirstWinningBoard = number
				}
				indexOfLastWinningBoard = i
				winningNumberOfLastWinningBoard = number

			}
		}
	}
	return indexOfFirstWinningBoard, winningNumberOfFirstWinningBoard, indexOfLastWinningBoard, winningNumberOfLastWinningBoard
}

func partOne() {
	utils.Intro("PART 1")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	bingoNumbers := getBingoNumbers(lines)
	fmt.Println("bingoNumbers:", bingoNumbers)
	bingoBoards := getBingoBoards(lines)
	// fmt.Printf("bingoBoards: %v\n", bingoBoards)
	winningBoardIndex, winningNumber, _, _ := playBingo(bingoNumbers, bingoBoards)
	if winningBoardIndex != -1 {
		fmt.Println("winningBoardIndex:", winningBoardIndex)
		fmt.Printf("board: %v\n", bingoBoards[winningBoardIndex])
		sumOfUnmarkedNumbers := bingoBoards[winningBoardIndex].GetSumOfUnmarkedNumbers()
		fmt.Println("sumOfUnmarkedNumbers:", sumOfUnmarkedNumbers)
		fmt.Println("winningNumber:", winningNumber)
		winningNumberInt, _ := utils.StringToInt(winningNumber)
		product := sumOfUnmarkedNumbers * winningNumberInt
		fmt.Println("product:", product)
	} else {
		fmt.Println("No winning board")
	}
}

func partTwo() {
	utils.Intro("PART 2")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	bingoNumbers := getBingoNumbers(lines)
	fmt.Println("bingoNumbers:", bingoNumbers)
	bingoBoards := getBingoBoards(lines)
	// fmt.Printf("bingoBoards: %v\n", bingoBoards)
	_, _, winningBoardIndex, winningNumber := playBingo(bingoNumbers, bingoBoards)
	if winningBoardIndex != -1 {
		fmt.Println("winningBoardIndex:", winningBoardIndex)
		fmt.Printf("board: %v\n", bingoBoards[winningBoardIndex])
		sumOfUnmarkedNumbers := bingoBoards[winningBoardIndex].GetSumOfUnmarkedNumbers()
		fmt.Println("sumOfUnmarkedNumbers:", sumOfUnmarkedNumbers)
		fmt.Println("winningNumber:", winningNumber)
		winningNumberInt, _ := utils.StringToInt(winningNumber)
		product := sumOfUnmarkedNumbers * winningNumberInt
		fmt.Println("product:", product)
	} else {
		fmt.Println("No winning board")
	}
}

func main() {
	partOne()
	partTwo()
}
