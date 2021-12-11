package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Intro(label string) {
	fmt.Println("========", label, "========")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func BinaryStringToInt(binaryString string) int64 {
	i, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToInts(s string, sep string) ([]int, error) {
	intVals := make([]int, 0)
	for _, stringVal := range strings.Split(s, sep) {
		intVal, err := strconv.Atoi(stringVal)
		if err != nil {
			return intVals, err
		}
		intVals = append(intVals, intVal)
	}
	return intVals, nil
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
		if intVal, err := StringToInt(line); err == nil {
			// fmt.Println("intVal:", intVal)
			intVals = append(intVals, intVal)
		} else {
			return intVals, err
		}
	}

	return intVals, nil
}

func GetTwoDArrayFromTextFile(filepath string) ([][]int, error) {
	lines, err := GetLinesFromTextFile(filepath)
	if err != nil {
		return nil, err
	}

	twoD := make([][]int, len(lines))
	for i, line := range lines {
		twoD[i] = make([]int, len(line))
		for j, c := range line {
			intVal, err := StringToInt(string(c))
			if err != nil {
				return twoD, err
			}
			twoD[i][j] = intVal
		}
	}
	return twoD, nil
}
