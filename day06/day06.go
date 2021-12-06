package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

func readInput(filename string) []int {
	var numbersInt []int

	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Failed to read input file")

	numbers := strings.Split(lines[0], ",")

	for _, number := range numbers {
		i, err := strconv.Atoi(number)
		logging.FailOnError(err, "Failed to convert string to int")
		numbersInt = append(numbersInt, i)
	}

	return numbersInt
}

func NextGeneration(input []int) []int {
	var nextGeneration []int
	var numberNewLanternfishes int

	for _, value := range input {
		if value == 0 {
			nextGeneration = append(nextGeneration, 6)
			numberNewLanternfishes++
		} else {
			nextGeneration = append(nextGeneration, value-1)
		}
	}

	for i := 0; i < numberNewLanternfishes; i++ {
		nextGeneration = append(nextGeneration, 8)
	}

	return nextGeneration
}

func Part1(input []int) int {
	// fmt.Println("Initial state:", input)
	for i := 0; i < 80; i++ {
		input = NextGeneration(input)
		// fmt.Println("After ", i+1, "day: ", input)
	}

	return len(input)
}

func Part2(input []int) int {
	return 0
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 6"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | How many lanternfish would there be after 80 days?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	// fmt.Println("* Part 2 | How many lanternfish would there be after 256 days?")
	// exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	// fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	// resultPart2 := strconv.Itoa(Part2(input))
	// fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
