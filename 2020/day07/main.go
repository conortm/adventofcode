package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

type bag struct {
	color    string
	contains map[string]int
}

func newBag(line string) *bag {
	lineParts := strings.Split(strings.TrimRight(line, "."), " bags contain ")
	color := lineParts[0]
	containBags := strings.Split(lineParts[1], ", ")
	contains := make(map[string]int)
	if containBags[0] != "no other bags" {
		for _, containBag := range containBags {
			containWords := strings.Split(containBag, " ")
			containCount, _ := strconv.Atoi(containWords[0])
			containColor := strings.Join(containWords[1:len(containWords)-1], " ")
			contains[containColor] = containCount
		}
	}
	b := bag{color: color, contains: contains}
	return &b
}

func bagContainsColor(b bag, color string, bags map[string]*bag, tracker map[string]bool) bool {
	contains, present := tracker[b.color]
	if present {
		return contains
	}
	for containColor := range b.contains {
		if color == containColor || bagContainsColor(*bags[containColor], color, bags, tracker) {
			tracker[b.color] = true
			return true
		}
	}
	tracker[b.color] = false
	return false
}

func countBagsWithinBag(b bag, bags map[string]*bag) int {
	count := 1
	for containColor, containCount := range b.contains {
		count += containCount * countBagsWithinBag(*bags[containColor], bags)
	}
	return count
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	bags := make(map[string]*bag)
	for _, line := range lines {
		// fmt.Printf("line: %s\n", line)
		b := newBag(line)
		// fmt.Printf("bag: %+v\n", b)
		bags[b.color] = b
	}
	count := 0
	tracker := make(map[string]bool)
	for _, b := range bags {
		if bagContainsColor(*b, "shiny gold", bags, tracker) {
			count++
		}
	}
	fmt.Printf("Count Shiny Gold: %d\n", count)
	countBags := countBagsWithinBag(*bags["shiny gold"], bags)
	fmt.Printf("Count Bags within Shiny Gold: %d\n", countBags-1)

}
