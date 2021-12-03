package main

import (
	"bufio"
	//"errors"
	"fmt"
	"log"
	"os"
	//"sort"
	//"strconv"
	//"strings"
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

func isCharAtIndex(line string, char string, index int) (bool, error) {
	// line repeats, so we mod the index by the length of line
	length := len(line)
	modIndex := index % length
	isCharAtIndex := (string(line[modIndex]) == char)
	fmt.Printf("Line: %s, Length: %d, Char: %s, Index: %d, modIndex: %d, isCharAtIndex: %t\n", line, length, char, index, modIndex, isCharAtIndex)
	return isCharAtIndex, nil
}

func getCountCharsAtSlope(lines []string, char string, right int, down int) (int, error) {
	count := 0
	x := 0
	y := 0
	distance := len(lines)
	for y <= distance-1 {
		line := lines[y]
		isChar, _ := isCharAtIndex(line, char, x)
		if isChar {
			count++
		}
		x += right
		y += down
	}
	return count, nil
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	trees11, _ := getCountCharsAtSlope(lines, "#", 1, 1)
	trees31, _ := getCountCharsAtSlope(lines, "#", 3, 1)
	trees51, _ := getCountCharsAtSlope(lines, "#", 5, 1)
	trees71, _ := getCountCharsAtSlope(lines, "#", 7, 1)
	trees12, _ := getCountCharsAtSlope(lines, "#", 1, 2)
	product := trees11 * trees31 * trees51 * trees71 * trees12
	fmt.Printf("product: %d\n", product)
}
