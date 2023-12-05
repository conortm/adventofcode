package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func Intro(label string) {
	fmt.Println("========", label, "========")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
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

func GetTwoDArrayFromTextFile(filepath string) ([][]string, error) {
	lines, err := GetLinesFromTextFile(filepath)
	if err != nil {
		return nil, err
	}

	twoDArray := make([][]string, len(lines))
	for i, line := range lines {
		twoDArray[i] = make([]string, len(line))
		for j, c := range line {
			twoDArray[i][j] = string(c)
		}
	}
	return twoDArray, nil
}
