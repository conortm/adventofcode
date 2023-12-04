package main

import (
	"fmt"
	"strings"

	"github.com/conortm/aoc/2023/utils"
)

type cubeGame struct {
	id          int
	cubeSets    []map[string]int
	colorCounts map[string]int
	power       int
}

func newCubeGame(line string) *cubeGame {
	cubeSets := make([]map[string]int, 0)
	colorCounts := make(map[string]int)
	lineParts := strings.Split(line, ": ")
	idParts := strings.Split(lineParts[0], " ")
	id, _ := utils.StringToInt(idParts[1])
	cubeSetsParts := strings.Split(lineParts[1], "; ")
	for _, cubeSetsPart := range cubeSetsParts {
		cubeSetStrings := strings.Split(cubeSetsPart, ", ")
		cubeSet := make(map[string]int)
		for _, cubeSetString := range cubeSetStrings {
			cubeSetStringParts := strings.Split(cubeSetString, " ")
			count, _ := utils.StringToInt(cubeSetStringParts[0])
			color := cubeSetStringParts[1]
			cubeSet[color] = count
			if count > colorCounts[color] {
				colorCounts[color] = count
			}
		}
		cubeSets = append(cubeSets, cubeSet)
	}
	power := colorCounts["red"] * colorCounts["green"] * colorCounts["blue"]
	return &cubeGame{id: id, cubeSets: cubeSets, colorCounts: colorCounts, power: power}
}

func (cg *cubeGame) String() string {
	return fmt.Sprintf("Game %d: %v; colorCounts: %v, power: %d", cg.id, cg.cubeSets, cg.colorCounts, cg.power)
}

func getCubeGamesFromLines(lines []string) []*cubeGame {
	cubeGames := make([]*cubeGame, 0)
	for _, line := range lines {
		cubeGame := newCubeGame(line)
		// fmt.Printf("%s\n", cubeGame)
		cubeGames = append(cubeGames, cubeGame)
	}
	return cubeGames
}

func partOne() {
	utils.Intro("PART 1")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	cubeGames := getCubeGamesFromLines(lines)
	possibleCubeGames := make([]*cubeGame, 0)
	colorCounts := map[string]int{"red": 12, "green": 13, "blue": 14}
	for _, cubeGame := range cubeGames {
		isPossible := true
		for color, count := range colorCounts {
			isPossible = (isPossible && cubeGame.colorCounts[color] <= count)
		}
		if isPossible {
			possibleCubeGames = append(possibleCubeGames, cubeGame)
		}
	}
	idSum := 0
	for _, possibleCubeGame := range possibleCubeGames {
		idSum += possibleCubeGame.id
	}
	fmt.Printf("idSum: %d\n", idSum)
}

func partTwo() {
	utils.Intro("PART 2")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	cubeGames := getCubeGamesFromLines(lines)
	powerSum := 0
	for _, cubeGame := range cubeGames {
		powerSum += cubeGame.power
	}
	fmt.Printf("powerSum: %d\n", powerSum)
}

func main() {
	partOne()
	partTwo()
}
