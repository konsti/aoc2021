package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
	log.Println("Advent of Code - Day 1")
	log.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	log.Println("* Part 1 [Example Input]")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	log.Printf("How many measurements are larger than the previous measurement?	%s \n\n", exampleResultPart1)

	log.Println("* Part 1 [Real Input]")
	resultPart1 := strconv.Itoa(Part1(input))
	log.Printf("How many measurements are larger than the previous measurement?	%s \n\n", resultPart1)

	log.Println("* Part 2 [Example Input]")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	log.Printf("How many sums are larger than the previous sum?	%s \n\n", exampleResultPart2)

	log.Println("* Part 2 [Real Input]")
	resultPart2 := strconv.Itoa(Part2(input))
	log.Printf("How many measurements are larger than the previous measurement?	%s \n\n", resultPart2)
}
