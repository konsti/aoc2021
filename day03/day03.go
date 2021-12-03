package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

func readInput(filename string) []string {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Failed to read input file")

	return lines
}

func addItem(counter map[bool]int, bit bool) {
	if _, ok := counter[bit]; ok {
		counter[bit]++
	} else {
		counter[bit] = 1
	}
}

func mostCommon(counter map[bool]int) bool {
	// Treat tie as true
	if counter[true] == counter[false] {
		return true
	}
	if counter[true] > counter[false] {
		return true
	}
	return false
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func makeBitCounter(input []string) []map[bool]int {
	// Prepare counters
	noBits := len(input[0])
	counters := make([]map[bool]int, noBits)
	for i := 0; i < noBits; i++ {
		counters[i] = make(map[bool]int)
	}

	// Count bits of each line
	for _, line := range input {
		for index, bit := range line {
			addItem(counters[index], string(bit) == "1")
		}
	}
	return counters
}

func getPart1Rate(counters []map[bool]int, useMostCommon bool) int {
	var rate strings.Builder
	for _, counter := range counters {
		if useMostCommon {
			rate.WriteString(strconv.Itoa(btoi(mostCommon(counter))))
		} else {
			rate.WriteString(strconv.Itoa(btoi(!mostCommon(counter))))
		}
	}
	rateInt, err := strconv.ParseInt(rate.String(), 2, 64)
	logging.FailOnError(err, "Failed to parse rate")
	return int(rateInt)
}

func Part1(input []string) int {
	counters := makeBitCounter(input)

	gammaRate := getPart1Rate(counters, true)
	fmt.Printf("Gamma Rate: %v\n", color.Teal(gammaRate))

	epsilonRate := getPart1Rate(counters, false)
	fmt.Printf("Epsilon Rate: %v\n", color.Teal(epsilonRate))

	return gammaRate * epsilonRate
}

func getPart2Rate(listOfNumbers []string, useMostCommon bool, checkIndex int) string {
	if len(listOfNumbers) == 1 {
		return listOfNumbers[0]
	}

	var filteredListOfNumbers []string

	counters := makeBitCounter(listOfNumbers)
	var mostCommonValue bool
	if useMostCommon {
		mostCommonValue = mostCommon(counters[checkIndex])
	} else {
		mostCommonValue = !mostCommon(counters[checkIndex])
	}

	for _, line := range listOfNumbers {
		if string(line[checkIndex]) == strconv.Itoa(btoi(mostCommonValue)) {
			filteredListOfNumbers = append(filteredListOfNumbers, line)
		}
	}

	return getPart2Rate(filteredListOfNumbers, useMostCommon, checkIndex+1)
}

func Part2(input []string) int {
	oxygenRating, err := strconv.ParseInt(getPart2Rate(input, true, 0), 2, 64)
	logging.FailOnError(err, "Failed to parse rate")
	fmt.Printf("Oxygen Rating: %v\n", color.Teal(oxygenRating))

	co2scrubberRating, err := strconv.ParseInt(getPart2Rate(input, false, 0), 2, 64)
	logging.FailOnError(err, "Failed to parse rate")
	fmt.Printf("CO2 Scrubber Rating: %v\n", color.Teal(co2scrubberRating))

	return int(oxygenRating) * int(co2scrubberRating)
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 1"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | What is the power consumption of the submarine?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	fmt.Println("* Part 2 | What is the life support rating of the submarine?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input))
	fmt.Printf(color.Green("[Real Input]:	%s \n"), color.Teal(resultPart2))
}
