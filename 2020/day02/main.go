package main

import (
	"bufio"
	//"errors"
	"fmt"
	"log"
	"os"

	//"sort"
	"strconv"
	"strings"
	"unicode"
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

func parseLine(line string) ([]string, error) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	return strings.FieldsFunc(line, f), nil
}

func isValidPassword(line string) (bool, error) {
	fields, _ := parseLine(line)
	min, _ := strconv.Atoi(fields[0])
	max, _ := strconv.Atoi(fields[1])
	char := fields[2]
	pass := fields[3]
	count := strings.Count(pass, char)
	// Part 1: isValid based on count of char being between min & max (inclusive)
	// isValid := (min <= count && count <= max)
	// Part 2: isValid based on exactly 1 occurence of char at either positions (min or max)
	minIsChar := (string(pass[min-1]) == char)
	maxIsChar := (string(pass[max-1]) == char)
	// Go doesn't have Exclusive Or, so:
	isValid := ((minIsChar || maxIsChar) && !(minIsChar && maxIsChar))
	fmt.Printf("Line: %s, Min: %d, Max: %d, Char: %s, Pass: %s, Count: %d, isValid: %t\n", line, min, max, char, pass, count, isValid)
	return isValid, nil
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	validPasswords := 0
	for _, line := range lines {
		isValidPassword, err := isValidPassword(line)
		if err != nil {
		}
		if isValidPassword {
			validPasswords++
		}
	}
	fmt.Printf("valid passwords: %d\n", validPasswords)
}
