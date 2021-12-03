package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Intro(label string) {
	fmt.Println("========", label, "========")
}

func BinaryStringToInt(binaryString string) int64 {
	i, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func GetLinesFromTextFile(filepath string) ([]string, error) {
	lines := make([]string, 0)

	file, err := os.Open(filepath)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func GetIntValsFromTextFile(filepath string) ([]int, error) {
	intVals := make([]int, 0)

	lines, err := GetLinesFromTextFile(filepath)
	if err != nil {
		return intVals, err
	}

	for _, line := range lines {
		if intVal, err := strconv.Atoi(line); err == nil {
			// fmt.Println("intVal:", intVal)
			intVals = append(intVals, intVal)
		} else {
			return intVals, err
		}
	}

	return intVals, nil
}
