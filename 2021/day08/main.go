package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/conortm/aoc/2021/utils"
)

/*
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg

2: 1
3: 7
4: 4
5: 2,3,5 => if contains segments in 1, then 3. if contains (diff 4 and 1), then 5. Else 2.
6: 0,6,9   => if contiains segments in 4, then 9. if contains segmenst in 1, then 0. Else 6.
7: 8
*/

type digit struct {
	segments map[string]bool
	pattern  string
	value    int
}

func (d *digit) String() string {
	return fmt.Sprintf("{pattern:%s, value:%v}", d.pattern, d.value)
}

func (d *digit) ContainsPattern(pattern string) bool {
	for _, c := range pattern {
		if _, contains := d.segments[string(c)]; !contains {
			return false
		}
	}
	return true
}

func (d1 *digit) DiffPattern(d2 *digit) string {
	diffPattern := ""
	for c := range d1.segments {
		if _, contains := d2.segments[string(c)]; !contains {
			diffPattern += c
		}
	}
	for c := range d2.segments {
		if _, contains := d1.segments[string(c)]; !contains {
			diffPattern += c
		}
	}
	return diffPattern
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func newDigit(digitSegments string) *digit {
	d := digit{}
	d.pattern = sortString(digitSegments)
	d.segments = make(map[string]bool)
	for _, c := range d.pattern {
		d.segments[string(c)] = true
	}
	d.value = -1
	switch len(d.pattern) {
	case 2:
		d.value = 1
	case 3:
		d.value = 7
	case 4:
		d.value = 4
	case 7:
		d.value = 8
	}
	return &d
}

type digitFactory struct {
	digitsByPattern map[string]*digit
	digitsByValue   map[int]*digit
}

func newDigitFactory(uniqueSignalPatterns []string) *digitFactory {
	digitsByPattern := make(map[string]*digit)
	digitsByValue := make(map[int]*digit)
	for _, pattern := range uniqueSignalPatterns {
		d := newDigit(pattern)
		digitsByPattern[d.pattern] = d
		if d.value >= 0 {
			digitsByValue[d.value] = d
		}
	}
	for _, d := range digitsByPattern {
		if d.value < 0 {
			switch len(d.pattern) {
			case 5:
				// if contains segments in 1, then 3. if contains (diff 4 and 1), then 5. Else 2.
				if d.ContainsPattern(digitsByValue[1].pattern) {
					d.value = 3
				} else if diff := digitsByValue[1].DiffPattern(digitsByValue[4]); d.ContainsPattern(diff) {
					d.value = 5
				} else {
					d.value = 2
				}
			case 6:
				// if contiains segments in 4, then 9. if contains segments in 1, then 0. Else 6.
				if d.ContainsPattern(digitsByValue[4].pattern) {
					d.value = 9
				} else if d.ContainsPattern(digitsByValue[1].pattern) {
					d.value = 0
				} else {
					d.value = 6
				}

			}
		}
	}
	return &digitFactory{digitsByPattern: digitsByPattern, digitsByValue: digitsByValue}
}

func (df *digitFactory) Decode(patterns []string) int {
	value := 0
	n := 0
	for i := len(patterns) - 1; i >= 0; i-- {
		pattern := sortString(patterns[i])
		digit := df.digitsByPattern[pattern].value
		value += digit * int(math.Pow10(n))
		n++
	}
	return value
}

func partOne() {
	utils.Intro("PART 1")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	digitsWithUniqueNumberOfSegmentsCount := 0
	for _, line := range lines {
		// fmt.Println("line:", line)
		parts := strings.Split(line, " | ")
		for _, outputDigit := range strings.Fields(parts[1]) {
			switch len(outputDigit) {
			case 2, 3, 4, 7:
				digitsWithUniqueNumberOfSegmentsCount++
			}
		}
	}
	fmt.Println("digitsWithUniqueNumberOfSegmentsCount:", digitsWithUniqueNumberOfSegmentsCount)
}

func partTwo() {
	utils.Intro("PART 2")
	lines, _ := utils.GetLinesFromTextFile("input.txt")
	sum := 0
	for _, line := range lines {
		// fmt.Println("line:", line)
		parts := strings.Split(line, " | ")
		uniqueSignalPatterns := strings.Fields(parts[0])
		df := newDigitFactory(uniqueSignalPatterns)
		// fmt.Println("uniqueSignalPatterns:", strings.Join(uniqueSignalPatterns, " "))
		outputPatterns := strings.Fields(parts[1])
		// fmt.Println("outputPatterns:", strings.Join(outputPatterns, " "))
		value := df.Decode(outputPatterns)
		// fmt.Println("value:", value)
		sum += value
	}
	fmt.Println("sum:", sum)
}

func main() {
	partOne()
	partTwo()
}
