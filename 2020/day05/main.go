package main

import (
	"bufio"
	"fmt"

	//"errors"

	"log"
	"os"
	//"unicode"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			return nil, err
		}
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func halfIsLower(half string) bool {
	return half == "F" || half == "L"
}

func search(min int, max int, halves string) int {
	// fmt.Printf("min: %d, max: %d, halves: %s\n", min, max, halves)
	half := string(halves[0])
	if len(halves) == 1 {
		if halfIsLower(half) {
			return min
		}
		return max
	}
	mid := min + ((max - min) / 2)
	// log.Printf("mid: %d", mid)
	if halfIsLower(half) {
		return search(min, mid, halves[1:])
	}
	return search(mid+1, max, halves[1:])
}

func decodeSeat(line string) (int, int, int) {
	row := search(0, 127, line[:7])
	col := search(0, 7, line[7:])
	id := (row*8 + col)
	return row, col, id
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	lowestSeatID := 1000
	highestSeatID := 0
	seatMap := make(map[int]bool)
	for _, line := range lines {
		row, col, id := decodeSeat(line)
		seatMap[id] = true
		fmt.Printf("line: %s, row: %d, col: %d, id: %d\n", line, row, col, id)
		if id < lowestSeatID {
			lowestSeatID = id
		}
		if id > highestSeatID {
			highestSeatID = id
		}
	}
	fmt.Printf("lowestSeatID: %d, highestSeatID: %d\n", lowestSeatID, highestSeatID)
	for i := lowestSeatID + 1; i < highestSeatID; i++ {
		_, inSeatMap := seatMap[i]
		if !inSeatMap {
			fmt.Printf("Seat ID: %d\n", i)
		}
	}
}
