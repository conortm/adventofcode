package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func getAdapters(lines []string) (map[int]bool, int) {
	adapters := make(map[int]bool)
	max := 0
	for _, line := range lines {
		adapter, _ := strconv.Atoi(line)
		adapters[adapter] = true
		if adapter > max {
			max = adapter
		}
	}
	return adapters, max
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	// Part 1: What is the number of 1-jolt differences multiplied
	// by the number of 3-jolt differences?
	adapters, max := getAdapters(lines)
	fmt.Printf("adapters: %v\n", adapters)
	diffs := make(map[int]int)
	diffs[1] = 0
	diffs[2] = 0
	diffs[3] = 0
	jolts := 0
	for i := 0; i <= max; i++ {
		_, prs := adapters[i]
		if prs {
			diff := i - jolts
			diffs[diff]++
			jolts = i
		}
	}
	jolts += 3
	diffs[3]++
	fmt.Printf("1-jolt diffs * 3-jolt diffs: %d\n", diffs[1]*diffs[3])

}
