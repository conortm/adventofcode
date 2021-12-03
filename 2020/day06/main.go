package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func groupLines(lines []string) [][]string {
	var groups [][]string
	group := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = make([]string, 0)
		} else {
			group = append(group, line)
		}
	}
	groups = append(groups, group)
	return groups
}

func getGroupCount(group []string) int {
	groupCount := 0
	numInGroup := len(group)

	ansMap := make(map[rune]int)
	for i, line := range group {
		isLastGroupMember := (i == numInGroup-1)
		for _, ans := range line {
			ansCount := ansMap[ans]
			ansCount++
			ansMap[ans] = ansCount
			if isLastGroupMember && numInGroup == ansCount {
				groupCount++
			}
		}
	}
	return groupCount
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	// for _, line := range lines {
	// 	fmt.Printf("line: %s\n", line)
	// }
	groups := groupLines(lines)
	totalCount := 0
	for _, group := range groups {
		fmt.Printf("group: %s\n", group)
		groupCount := getGroupCount(group)
		fmt.Printf("groupCount: %d\n", groupCount)
		totalCount += groupCount
	}
	fmt.Printf("totalCount: %d\n", totalCount)
}
