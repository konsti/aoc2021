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

// The part2 fuel consumption is a triangular number sequence formula.
func fuelConsumption(distance float64) float64 {
	return (distance * (distance + 1)) / 2
}

func fullFuelConsumption(crabPositions []float64, target float64) float64 {
	fuel := 0.0
	for _, position := range crabPositions {
		fuel += fuelConsumption(math.Abs(position - target))
	}
	return fuel
}

func Part2(crabPositions []float64) int {
	min := 100.0
	max := 0.0
	var fuels []float64

	for _, position := range crabPositions {
		if position > max {
			max = position
		}
		if position < min {
			min = position
		}
	}

	for i := min; i <= max; i++ {
		fuels = append(fuels, fullFuelConsumption(crabPositions, i))
	}

	sort.Float64s(fuels)

	return int(fuels[0])
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

	// Part 2

	fmt.Println("* Part 2 | How much fuel must they spend to align to the new position regarding expensive fuel?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
