package main

import (
	"fmt"
	"strconv"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

func readInput(filename string) []int {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Failed to read input file")

	var input []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		logging.FailOnError(err, "Failed to convert input to integer")
		input = append(input, num)
	}
	return input
}

func Part1(input []int) int {
	accumulator := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			accumulator++
		}
	}
	return accumulator
}

func Part2(input []int) int {
	accumulator := 0
	for i := 0; i < len(input)-3; i++ {
		window1 := input[i] + input[i+1] + input[i+2]
		window2 := input[i+1] + input[i+2] + input[i+3]
		if window2 > window1 {
			accumulator++
		}
	}

	return accumulator
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 1"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | How many measurements are larger than the previous measurement?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	fmt.Println("* Part 2 | How many sums are larger than the previous sum?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input))
	fmt.Printf(color.Green("[Real Input]:	%s \n"), color.Teal(resultPart2))
}
