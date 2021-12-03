package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetIntValsFromFile(filepath string) ([]int, error) {
	intVals := make([]int, 0)

	file, err := os.Open(filepath)
	if err != nil {
		return intVals, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if intVal, err := strconv.Atoi(text); err == nil {
			// fmt.Println("intVal:", intVal)
			intVals = append(intVals, intVal)
		} else {
			return intVals, err
		}
	}

	if err := scanner.Err(); err != nil {
		return intVals, err
	}

	return intVals, nil
}

func GetLinesFromTextFile(filepath string, lines chan<- string) {
	file, _ := os.Open(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines <- scanner.Text()
	}
	close(lines)
}

func Intro(label string) {
	fmt.Println("========", label, "========")
}
