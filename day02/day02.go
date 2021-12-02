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

type Command struct {
	Direction string
	Value     int
}

func lineToCommand(line string) Command {
	slice := strings.Fields(line)
	if len(slice) != 2 {
		log.Fatalf("Invalid input %s", slice)
		panic("")
	}

	value, err := strconv.Atoi(slice[1])
	logging.FailOnError(err, "Could not convert value to int")

	if slice[0] == "up" || slice[0] == "down" || slice[0] == "forward" {
		return Command{
			Direction: slice[0],
			Value:     value,
		}
	}

	log.Fatalf("Invalid command", slice[0])
	panic("")
}

func lineToVector(line string) Vector {
	command := lineToCommand(line)

	switch command.Direction {
	case "down":
		return Vector{0, command.Value}
	case "up":
		return Vector{0, -command.Value}
	}

	return Vector{command.Value, 0}
}

func readInputPart1(filename string) []Vector {
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

func Part2(filename string) int {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Could not read input")

	aim := 0
	var movementVectors []Vector
	for _, line := range lines {
		command := lineToCommand(line)
		switch command.Direction {
		case "down":
			aim = aim + command.Value
		case "up":
			aim = aim - command.Value
		case "forward":
			vector := Vector{command.Value, aim * command.Value}
			movementVectors = append(movementVectors, vector)
		}
	}

	finalPosition := sum(movementVectors)
	fmt.Printf("Final Position: %v\n", color.Teal(finalPosition))
	return finalPosition.X * finalPosition.Y
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 2"))
	fmt.Print("======================\n\n")

	// Part 1

	fmt.Println("* Part 1 | What do you get if you multiply your final horizontal position by your final depth?")
	exampleInput := readInputPart1("example.txt")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	input := readInputPart1("input.txt")
	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	fmt.Println("* Part 2 | What do you get if you multiply your final horizontal position by your final depth?")
	exampleResultPart2 := strconv.Itoa(Part2("example.txt"))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2("input.txt"))
	fmt.Printf(color.Green("[Real Input]:	%s \n"), color.Teal(resultPart2))
}
