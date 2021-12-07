package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

func readInput(filename string) []float64 {
	var numbers []float64

	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	numbersStr := strings.Split(lines[0], ",")

	for _, number := range numbersStr {
		i, err := strconv.ParseFloat(number, 64)
		logging.FailOnError(err, "Error converting string to float64")
		numbers = append(numbers, i)
	}

	return numbers
}

func median(input []float64) float64 {
	sort.Float64s(input)
	median := float64(0)
	if len(input)%2 == 0 {
		median = (input[len(input)/2-1] + input[len(input)/2]) / 2
	} else {
		median = input[len(input)/2]
	}
	return median
}

func Part1(crabPositions []float64) int {
	medianPosition := median(crabPositions)
	fuel := 0

	fmt.Println("Median position:", medianPosition)
	for _, position := range crabPositions {
		fuel += int(math.Abs(position - medianPosition))
	}

	return fuel
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 7"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | How much fuel must they spend to align to the median position?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))
}
