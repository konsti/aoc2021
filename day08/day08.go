package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

type SignalPattern struct {
	uniqueSignalPatterns []string
	fourDigitOutputValue []string
}

func readInput(filename string) []SignalPattern {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	var signalPatterns []SignalPattern

	for _, line := range lines {
		parts := strings.Split(line, "|")
		// strings.TrimSpace(?)
		signalPatterns = append(signalPatterns, SignalPattern{
			uniqueSignalPatterns: strings.Fields(strings.TrimSpace(parts[0])),
			fourDigitOutputValue: strings.Fields(strings.TrimSpace(parts[1])),
		})
	}

	return signalPatterns
}

func Part1(input []SignalPattern) int {
	count := 0

	// 0 = 6

	// 1 = 2 unique!

	// 2 = 5
	// 3 = 5

	// 4 = 4 unique!

	// 5 = 5
	// 6 = 6

	// 7 = 3 unique!
	// 8 = 7 unique!

	// 9 = 6

	for _, signalPattern := range input {
		for _, digit := range signalPattern.fourDigitOutputValue {
			if len(digit) == 2 || len(digit) == 4 || len(digit) == 3 || len(digit) == 7 {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day8"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | In the output values, how many times do digits 1, 4, 7, or 8 appear?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	// fmt.Println("* Part 2 | How much fuel must they spend to align to the new position regarding expensive fuel?")
	// exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	// fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	// resultPart2 := strconv.Itoa(Part2(input))
	// fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
