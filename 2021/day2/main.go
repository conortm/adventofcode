package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/conortm/aoc/2021/utils"
)

type command struct {
	raw       string
	direction string
	distance  int
}

func newCommand(text string) *command {
	parts := strings.Split(text, " ")
	direction := parts[0]
	distance, _ := strconv.Atoi(parts[1])
	return &command{raw: text, direction: direction, distance: distance}
}

func getCommandsFromInputFile(filename string) []*command {
	commands := make([]*command, 0)
	lines, _ := utils.GetLinesFromTextFile(filename)
	for _, line := range lines {
		commands = append(commands, newCommand(line))
	}
	// for _, command := range commands {
	// 	fmt.Printf("%+v\n", command)
	// }
	return commands
}

func getFinalCoordinates(commands []*command) (int, int) {
	horizontalPosition := 0
	depth := 0

	for _, command := range commands {
		switch command.direction {
		case "forward":
			horizontalPosition += command.distance
		case "down":
			depth += command.distance
		case "up":
			depth -= command.distance
		}
	}
	fmt.Println("horizontalPosition:", horizontalPosition)
	fmt.Println("depth:", depth)

	return horizontalPosition, depth
}

func getFinalCoordinatesWithAim(commands []*command) (int, int) {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, command := range commands {
		switch command.direction {
		case "forward":
			horizontalPosition += command.distance
			depth += aim * command.distance
		case "down":
			aim += command.distance
		case "up":
			aim -= command.distance
		}
	}
	fmt.Println("horizontalPosition:", horizontalPosition)
	fmt.Println("depth:", depth)
	fmt.Println("aim:", aim)

	return horizontalPosition, depth
}

func main() {
	utils.Intro("PART 1")
	commands := getCommandsFromInputFile("input.txt")
	horizontalPosition, depth := getFinalCoordinates(commands)
	product := horizontalPosition * depth
	fmt.Println("product:", product)

	utils.Intro("PART 2")
	commands = getCommandsFromInputFile("input.txt")
	horizontalPosition, depth = getFinalCoordinatesWithAim(commands)
	product = horizontalPosition * depth
	fmt.Println("product:", product)
}
