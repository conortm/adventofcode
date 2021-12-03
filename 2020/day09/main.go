package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func getNumbers(lines []string) []int {
	numbers := make([]int, 0)
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	return numbers
}

func getPreamble(numbers []int, left int, right int) []int {
	return numbers[left:right]
}

func updatePreamble(number int, preamble []int) []int {
	return append(preamble, number)[1:]
}

func isValidNumber(number int, preamble []int) bool {
	// Number is valid if it is the sum of any 2 numbers from preamble
	halfOfNumber := number / 2
	lenPreamble := len(preamble)
	sortedPreamble := make([]int, len(preamble))
	copy(sortedPreamble, preamble)
	sort.Ints(sortedPreamble)
	for i := 0; i < lenPreamble; i++ {
		if sortedPreamble[i] > halfOfNumber {
			break
		}
		for j := i + 1; j < lenPreamble; j++ {
			sum := sortedPreamble[i] + sortedPreamble[j]
			// fmt.Printf("number: %d, halfOfNumber: %d, i: %d, sortedPreamble[i]: %d, j: %d, sortedPreamble[j]: %d, sum: %d\n", number, halfOfNumber, i, sortedPreamble[i], j, sortedPreamble[j], sum)
			if sum == number {
				return true
			} else if sum > number {
				break
			}
		}
	}
	return false
}

func getSum(numbers []int, left int, right int) int {
	sum := 0
	for _, number := range numbers[left:right] {
		sum += number
	}
	return sum
}

func getMin(numbers []int, left int, right int) int {
	min := math.MaxInt32
	for _, number := range numbers[left:right] {
		if number < min {
			min = number
		}
	}
	return min
}

func getMax(numbers []int, left int, right int) int {
	max := 0
	for _, number := range numbers[left:right] {
		if number > max {
			max = number
		}
	}
	return max
}

func getWindowMinMaxSum(target int, numbers []int, left int, right int) int {
	sum := getSum(numbers, left, right)
	if sum == target {
		min := getMin(numbers, left, right)
		max := getMax(numbers, left, right)
		minMaxSum := min + max
		// log.Panicf("sum: %d, target: %d, left: %d, right: %d, min: %d, max: %d, slice: %v\n", sum, target, left, right, max, numbers[left:right])
		return minMaxSum
	} else if sum < target {
		return getWindowMinMaxSum(target, numbers, left, right+1)
	} else {
		return getWindowMinMaxSum(target, numbers, left+1, right)
	}
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	numbers := getNumbers(lines)
	// Part 1: find the first number in the list (after the preamble)
	// which is not the sum of two of the 25 numbers before it.
	// load preamble
	// range over rest of list, for each:
	// 1. check if is valid
	// 2. update preamble
	var invalidNumber int
	lenNumbers := len(numbers)
	lenPreamble := 25
	preamble := getPreamble(numbers, 0, lenPreamble)
	for i := lenPreamble; i < lenNumbers; i++ {
		// fmt.Printf("preamble: %v\n", preamble)
		number := numbers[i]
		if !isValidNumber(number, preamble) {
			invalidNumber = number
			break
		}
		preamble = updatePreamble(number, preamble)
	}
	fmt.Printf("invalidNumber: %d\n", invalidNumber)
	// Part 2: find a contiguous set of at least two numbers in your list
	// which sum to the invalid number from step 1.
	minMaxSum := getWindowMinMaxSum(invalidNumber, numbers, 0, 1)
	fmt.Printf("minMaxSum: %d\n", minMaxSum)
}
