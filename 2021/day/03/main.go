package main

import (
	"fmt"

	"github.com/conortm/aoc/2021/utils"
)

func getOxygenGeneratorRating(filename string) (string, error) {
	oxygenGeneratorRating := "0"
	values, err := utils.GetLinesFromTextFile(filename)
	if err != nil {
		return oxygenGeneratorRating, err
	}
	for i := 0; i < 12; i++ {
		zeroBits := 0
		oneBits := 0
		for _, value := range values {
			if value[i] == '0' {
				zeroBits++
			} else {
				oneBits++
			}
		}
		var mostCommonBitAtPosition byte = '0'
		if oneBits >= zeroBits {
			mostCommonBitAtPosition = '1'
		}
		newValues := make([]string, 0)
		for _, value := range values {
			if value[i] == mostCommonBitAtPosition {
				newValues = append(newValues, value)
			}
		}
		if len(newValues) == 1 {
			return newValues[0], nil
		}
		values = newValues
	}
	return oxygenGeneratorRating, nil
}

func getCO2ScrubberRating(filename string) (string, error) {
	cO2ScrubberRating := "0"
	values, err := utils.GetLinesFromTextFile(filename)
	if err != nil {
		return cO2ScrubberRating, err
	}
	for i := 0; i < 12; i++ {
		zeroBits := 0
		oneBits := 0
		for _, value := range values {
			if value[i] == '0' {
				zeroBits++
			} else {
				oneBits++
			}
		}
		var leastCommonBitAtPosition byte = '0'
		if oneBits < zeroBits {
			leastCommonBitAtPosition = '1'
		}
		newValues := make([]string, 0)
		for _, value := range values {
			if value[i] == leastCommonBitAtPosition {
				newValues = append(newValues, value)
			}
		}
		if len(newValues) == 1 {
			return newValues[0], nil
		}
		values = newValues
	}
	return cO2ScrubberRating, nil
}

func getRates(filename string) (string, string) {
	var gamma string
	var epsilon string
	counts := make([][]int, 12)
	for i := 0; i < 12; i++ {
		counts[i] = make([]int, 2)
		counts[i][0] = 0
		counts[i][1] = 0
	}
	lines, _ := utils.GetLinesFromTextFile(filename)
	for _, line := range lines {
		for i, c := range line {
			switch c {
			case '0':
				counts[i][0]++
			case '1':
				counts[i][1]++
			}
		}
	}
	// fmt.Printf("%+v\n", counts)
	for _, c := range counts {
		if c[0] > c[1] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	return gamma, epsilon
}

func main() {
	utils.Intro("PART 1")
	gamma, epsilon := getRates("input.txt")
	fmt.Println("gamma:", gamma)
	fmt.Println("epsilon:", epsilon)
	gammaInt := utils.BinaryStringToInt(gamma)
	epsilonInt := utils.BinaryStringToInt(epsilon)
	fmt.Println("gammaInt:", gammaInt)
	fmt.Println("epsilonInt:", epsilonInt)
	product := gammaInt * epsilonInt
	fmt.Println("product:", product)

	utils.Intro("PART 2")
	oxygenGeneratorRating, _ := getOxygenGeneratorRating("input.txt")
	fmt.Println("oxygenGeneratorRating:", oxygenGeneratorRating)
	cO2ScrubberRating, _ := getCO2ScrubberRating("input.txt")
	fmt.Println("cO2ScrubberRating:", cO2ScrubberRating)
	oxygenGeneratorRatingInt := utils.BinaryStringToInt(oxygenGeneratorRating)
	cO2ScrubberRatingInt := utils.BinaryStringToInt(cO2ScrubberRating)
	fmt.Println("oxygenGeneratorRatingInt:", oxygenGeneratorRatingInt)
	fmt.Println("cO2ScrubberRatingInt:", cO2ScrubberRatingInt)
	product = oxygenGeneratorRatingInt * cO2ScrubberRatingInt
	fmt.Println("product:", product)
}
