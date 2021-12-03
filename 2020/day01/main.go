package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// readInts reads a whole file into memory
// and returns a slice of its lines.
func readInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ints []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ints = append(ints, val)
	}
	return ints, scanner.Err()
}

func getProductOfTwoEntriesThatSum2020(entries []int) (int, error) {
	sort.Ints(entries)
	for i, entry := range entries {
		target := 2020 - entry
		for _, otherEntry := range entries[i+1:] {
			if target == otherEntry {
				fmt.Printf("Found entries: %d, %d\n", entry, target)
				return entry * target, nil
			}
		}
	}
	return -1, errors.New("Not found")
}

func getTwoEntriesThatSum(entries []int, sum int) (int, int, error) {
	sort.Ints(entries)
	for i, entry := range entries {
		target := sum - entry
		for _, otherEntry := range entries[i+1:] {
			if target == otherEntry {
				fmt.Printf("Found entries: %d, %d\n", entry, target)
				return entry, target, nil
			}
		}
	}
	return -1, -1, errors.New("Not found")
}

func getThreeEntriesThatSum2020(entries []int) (int, int, int, error) {
	sort.Ints(entries)
	for i, entry1 := range entries {
		sum := 2020 - entry1
		entry2, entry3, err := getTwoEntriesThatSum(entries[i+1:], sum)
		if err == nil {
			fmt.Printf("Found entries: %d, %d, %d\n", entry1, entry2, entry3)
			return entry1, entry2, entry3, nil
		}
	}
	return -1, -1, -1, errors.New("Not found")
}

func getProductOfThreeEntriesThatSum2020(entries []int) (int, error) {
	entry1, entry2, entry3, err := getThreeEntriesThatSum2020(entries)
	if err != nil {
		return -1, err
	}
	return entry1 * entry2 * entry3, nil
}

func main() {
	lines, err := readInts("input.txt")
	if err != nil {
		log.Fatalf("readInts: %s", err)
	}
	// sort.Ints(lines)
	// for i, line := range lines {
	//    fmt.Println(i, line)
	//}
	product, err := getProductOfTwoEntriesThatSum2020(lines)
	if err != nil {
		log.Fatalf("getProductTwo: %s\n", err)
	}
	fmt.Printf("getProductOfTwoEntriesThatSum2020: %d\n", product)

	threeProduct, err := getProductOfThreeEntriesThatSum2020(lines)
	if err != nil {
		log.Fatalf("getProductThree: %s\n", err)
	}
	fmt.Printf("getProductOfThreeEntriesThatSum2020: %d\n", threeProduct)
}
