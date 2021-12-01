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

func main() {
	log.Println("Advent of Code - Day 1")
	log.Print("======================\n\n")

	log.Println("* Part 1 [Example Input]")

	exampleInput := readInput("example.txt")
	exampleResult := strconv.Itoa(Part1(exampleInput))
	log.Printf("How many measurements are larger than the previous measurement?	%s \n\n", exampleResult)

	input := readInput("input.txt")
	result := strconv.Itoa(Part1(input))
	log.Printf("How many measurements are larger than the previous measurement?	%s \n\n", result)
}
