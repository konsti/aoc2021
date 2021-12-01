package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	Green  = color("32")
	Yellow = color("33")
	Purple = color("35")
	Teal   = color("36")
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic("")
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readInput(filename string) []int {
	lines, err := readLines(filename)
	failOnError(err, "Failed to read input file")

	var input []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		failOnError(err, "Failed to convert input to integer")
		input = append(input, num)
	}
	return input
}

func color(colorString string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf("\x1b["+colorString+"m%s\x1b[0m", fmt.Sprint(args...))
	}
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
	fmt.Println(Purple("Advent of Code - Day 1"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | How many measurements are larger than the previous measurement?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(Yellow("[Example Input]: %s \n"), Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(Green("[Real Input]: %s \n\n"), Teal(resultPart1))

	// Part 2

	fmt.Println("* Part 2 | How many sums are larger than the previous sum?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	fmt.Printf(Yellow("[Example Input]: %s \n"), Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input))
	fmt.Printf(Green("[Real Input]:	%s \n"), Teal(resultPart2))
}
