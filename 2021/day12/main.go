package main

import (
	"fmt"
	"strings"

	"github.com/conortm/aoc/2021/utils"
)

type cave struct {
	name        string
	size        string
	connections map[string]bool
}

func newCave(name string) *cave {
	size := "small"
	if name == strings.ToUpper(name) {
		size = "big"
	}
	connections := make(map[string]bool, 0)
	return &cave{name: name, size: size, connections: connections}
}

func (c *cave) String() string {
	return fmt.Sprintf("{name: %s, size: %s, connections: %+v}", c.name, c.size, c.connections)
}

func (c *cave) AddConnection(name string) {
	c.connections[name] = true
}

func getCaves(filename string) map[string]*cave {
	lines, _ := utils.GetLinesFromTextFile(filename)
	caves := make(map[string]*cave, 0)
	for _, line := range lines {
		names := strings.Split(line, "-")
		name1 := names[0]
		name2 := names[1]
		cave1, exists := caves[name1]
		if !exists {
			cave1 = newCave(name1)
			caves[name1] = cave1
		}
		cave2, exists := caves[name2]
		if !exists {
			cave2 = newCave(name2)
			caves[name2] = cave2
		}
		cave1.AddConnection(cave2.name)
		cave2.AddConnection(cave1.name)
	}
	return caves
}

func partOne() {
	utils.Intro("PART 1")
	caves := getCaves("input.txt")
	count := 0
	var appendCaveToPath func(path []string, name string)
	appendCaveToPath = func(path []string, name string) {
		if path == nil {
			path = make([]string, 0)
		}
		cave := caves[name]
		if !(cave.size == "small" && utils.Contains(path, name)) {
			path = append(path, name)
			if name == "end" {
				count++
				// fmt.Println(path)
			} else {
				for connection := range cave.connections {
					appendCaveToPath(path, connection)
				}
			}
		}
	}
	appendCaveToPath(nil, "start")
	fmt.Println("count:", count)
}

func hasSmallCaveVisitedTwice(path []string, caves map[string]*cave) bool {
	smallCaveCount := make(map[string]int)
	for _, name := range path {
		cave := caves[name]
		if cave.size == "small" {
			_, exists := smallCaveCount[name]
			if exists {
				return true
			} else {
				smallCaveCount[name] = 1
			}
		}
	}
	return false
}

func partTwo() {
	utils.Intro("PART 2")
	caves := getCaves("input.txt")
	count := 0
	var appendCaveToPath func(path []string, name string)
	appendCaveToPath = func(path []string, name string) {
		if path == nil {
			path = make([]string, 0)
		}
		cave := caves[name]
		canVisit := true
		if cave.size == "small" && utils.Contains(path, name) {
			switch name {
			case "start", "end":
				canVisit = false
			default:
				canVisit = !(hasSmallCaveVisitedTwice(path, caves))
			}
		}
		if canVisit {
			path = append(path, name)
			if name == "end" {
				count++
				// fmt.Println(path)
			} else {
				for connection := range cave.connections {
					appendCaveToPath(path, connection)
				}
			}
		}
	}
	appendCaveToPath(nil, "start")
	fmt.Println("count:", count)
}

func main() {
	partOne()
	partTwo()
}
