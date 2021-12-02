package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

type Vector struct {
	X, Y int
}

func (a Vector) add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func sum(array []Vector) Vector {
	result := Vector{}
	for _, v := range array {
		result = result.add(v)
	}
	return result
}

func lineToVector(line string) Vector {
	slice := strings.Fields(line)
	if len(slice) != 2 {
		log.Fatalf("Invalid input %s", slice)
		panic("")
	}

	value, err := strconv.Atoi(slice[1])
	logging.FailOnError(err, "Could not convert value to int")

	switch slice[0] {
	case "forward":
		return Vector{value, 0}
	case "down":
		return Vector{0, value}
	case "up":
		return Vector{0, -value}
	}

	log.Fatalf("Invalid command", slice[0])
	panic("")
}

func readInput(filename string) []Vector {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Could not read input")

	var input []Vector
	for _, line := range lines {
		input = append(input, lineToVector(line))
	}

	return input
}

func Part1(input []Vector) int {
	finalPosition := sum(input)
	fmt.Printf("Final Position: %v\n", color.Teal(finalPosition))
	return finalPosition.X * finalPosition.Y
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 2"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | What do you get if you multiply your final horizontal position by your final depth?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))
}
