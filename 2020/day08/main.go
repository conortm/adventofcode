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

type instruction struct {
	op  string
	arg int
	ex  bool
}

func newInstruction(line string) *instruction {
	parts := strings.Split(line, " ")
	op := parts[0]
	arg, _ := strconv.Atoi(parts[1])
	ex := false
	i := &instruction{op: op, arg: arg, ex: ex}
	return i
}

func execute(instructions []*instruction) (int, bool) {
	pos := 0
	accumulator := 0
	didTerminate := true
	for pos < len(instructions) {
		ins := instructions[pos]
		if ins.ex {
			didTerminate = false
			break
		}
		switch ins.op {
		case "nop":
			pos++
		case "acc":
			accumulator += ins.arg
			pos++
		case "jmp":
			pos += ins.arg
		}
		ins.ex = true
	}
	return accumulator, didTerminate
}

func getInstructions(lines []string) []*instruction {
	instructions := make([]*instruction, 0)
	for _, line := range lines {
		// fmt.Printf("line: %s\n", line)
		ins := newInstruction(line)
		// fmt.Printf("instruction: %+v\n", ins)
		instructions = append(instructions, ins)
	}
	return instructions
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	instructions := getInstructions(lines)
	accumulator, _ := execute(instructions)
	fmt.Printf("Accumulator before loop: %d\n", accumulator)
	// Part 2.
	for i := 0; i < len(instructions); i++ {
		editedInstructions := getInstructions(lines)
		switch instructions[i].op {
		case "jmp":
			editedInstructions[i].op = "nop"
		case "nop":
			editedInstructions[i].op = "jmp"
		}
		// fmt.Printf("instructions[%d]: %+v\n", i, instructions[i])
		// fmt.Printf("editedInstructions[%d]: %+v\n", i, editedInstructions[i])
		accumulator, didTerminate := execute(editedInstructions)
		if didTerminate {
			fmt.Printf("Accumulator when didTerminate: %d\n", accumulator)
		}
	}
}
