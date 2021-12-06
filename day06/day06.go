package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

func readInput(filename string) [9]int {
	var numbersInt []int

	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Failed to read input file")

	numbers := strings.Split(lines[0], ",")

	for _, number := range numbers {
		i, err := strconv.Atoi(number)
		logging.FailOnError(err, "Failed to convert string to int")
		numbersInt = append(numbersInt, i)
	}

	var population [9]int
	for _, value := range numbersInt {
		population[value]++
	}

	return population
}

func rotateLeft(nums [9]int, n int) [9]int {
	if n < 0 || len(nums) == 0 {
		return nums
	}
	rotatedNums := append(nums[n:], nums[:n]...)

	var returnNums [9]int
	copy(returnNums[:], rotatedNums)

	return returnNums
}

func growPopulation(population [9]int, days int) [9]int {
	for i := 0; i < days; i++ {
		population = rotateLeft(population, 1)
		population[6] += population[8]
	}
	return population
}

func sum(nums [9]int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func Part1(population [9]int) int {
	population80 := growPopulation(population, 80)
	return sum(population80)
}

func Part2(population [9]int) int {
	population256 := growPopulation(population, 256)
	return sum(population256)
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

	fmt.Println("* Part 2 | How many lanternfish would there be after 256 days?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
