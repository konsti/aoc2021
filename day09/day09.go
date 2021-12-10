package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

func readInput(filename string) [][]int {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	var input [][]int

	for _, line := range lines {
		var row []int
		numbers := strings.Split(line, "")
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			logging.FailOnError(err, "Error converting string to int")
			row = append(row, num)
		}
		input = append(input, row)
	}

	return input
}

func Part1(input [][]int) int {
	var lowPoints []int
	maxX := len(input[0])
	maxY := len(input)
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			point := input[y][x]

			if x == 0 {
				if y == 0 {
					if point < input[y+1][x] && point < input[y][x+1] {
						lowPoints = append(lowPoints, point)
					}
				} else if y == maxY-1 {
					if point < input[y-1][x] && point < input[y][x+1] {
						lowPoints = append(lowPoints, point)
					}
				} else {
					if point < input[y-1][x] && point < input[y][x+1] && point < input[y+1][x] {
						lowPoints = append(lowPoints, point)
					}
				}
			} else if x == maxX-1 {
				if y == 0 {
					if point < input[y+1][x] && point < input[y][x-1] {
						lowPoints = append(lowPoints, point)
					}
				} else if y == maxY-1 {
					if point < input[y-1][x] && point < input[y][x-1] {
						lowPoints = append(lowPoints, point)
					}
				} else {
					if point < input[y-1][x] && point < input[y][x-1] && point < input[y+1][x] {
						lowPoints = append(lowPoints, point)
					}
				}
			} else {
				if y == 0 {
					if point < input[y+1][x] && point < input[y][x-1] && point < input[y][x+1] {
						lowPoints = append(lowPoints, point)
					}
				} else if y == maxY-1 {
					if point < input[y-1][x] && point < input[y][x-1] && point < input[y][x+1] {
						lowPoints = append(lowPoints, point)
					}
				} else {
					if point < input[y-1][x] && point < input[y][x-1] && point < input[y][x+1] && point < input[y+1][x] {
						lowPoints = append(lowPoints, point)
					}
				}
			}
		}
	}

	sumRiskLevel := 0

	for _, point := range lowPoints {
		sumRiskLevel += point + 1
	}

	return sumRiskLevel
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day9"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | What is the sum of the risk levels of all low points on your heightmap?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	// fmt.Println("* Part 2 | What do you get if you add up all of the output values?")
	// exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	// fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	// resultPart2 := strconv.Itoa(Part2(input))
	// fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
