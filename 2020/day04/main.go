package main

import (
	"bufio"
	//"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func combineLines(lines []string) ([]string, error) {
	var combinedLines []string
	var b strings.Builder
	for _, line := range lines {
		if len(line) == 0 {
			combinedLines = append(combinedLines, strings.TrimSpace(b.String()))
			b.Reset()
		} else {
			b.WriteString(line)
			b.WriteString(" ")
		}
	}
	combinedLines = append(combinedLines, strings.TrimSpace(b.String()))
	return combinedLines, nil
}

func isFieldValid(name string, value string) bool {
	isFieldValid := true
	intValue, intValueErr := strconv.Atoi(value)
	switch name {
	case "byr":
		isFieldValid = (1920 <= intValue && intValue <= 2002)
	case "iyr":
		isFieldValid = (2010 <= intValue && intValue <= 2020)
	case "eyr":
		isFieldValid = (2020 <= intValue && intValue <= 2030)
	case "hgt":
		isCentimeters, _ := regexp.MatchString("^[0-9]{3}cm$", value)
		isInches, _ := regexp.MatchString("^[0-9]{2}in$", value)
		if isCentimeters {
			centimeters, _ := strconv.Atoi(strings.TrimRight(value, "cm"))
			isFieldValid = (150 <= centimeters && centimeters <= 193)
		} else if isInches {
			inches, _ := strconv.Atoi(strings.TrimRight(value, "in"))
			isFieldValid = (59 <= inches && inches <= 76)
		} else {
			isFieldValid = false
		}
	case "hcl":
		isFieldValid, _ = regexp.MatchString("^#[0-9a-f]{6}$", value)
	case "ecl":
		m := make(map[string]bool)
		m["amb"] = true
		m["blu"] = true
		m["brn"] = true
		m["gry"] = true
		m["grn"] = true
		m["hzl"] = true
		m["oth"] = true
		_, isFieldValid = m[value]
	case "pid":
		isFieldValid = (intValueErr == nil && len(value) == 9)
	}
	fmt.Printf("Field name: %s, value: %s, isFieldValid: %t\n", name, value, isFieldValid)
	// if !isFieldValid {
	// 	log.Fatalf("Field name: %s, value: %s, isFieldValid: %t\n", name, value, isFieldValid)
	// }
	return isFieldValid
}

func isValid(line string) bool {
	prefixes := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	// Check that all required fields exist.
	for _, prefix := range prefixes {
		if !strings.Contains(line, prefix+":") {
			return false
		}
	}
	// Validate each field
	fields := strings.Fields(line)
	for _, field := range fields {
		parts := strings.Split(field, ":")
		if !isFieldValid(parts[0], parts[1]) {
			return false
		}
	}
	return true
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	combinedLines, _ := combineLines(lines)
	validPasswords := 0
	for _, line := range combinedLines {
		fmt.Printf("Line: %s\n", line)
		isValid := isValid(line)
		if isValid {
			validPasswords++
		}
		fmt.Printf("isValid: %t\n======\n", isValid)
	}
	fmt.Printf("Valid passwords: %d\n", validPasswords)
}
