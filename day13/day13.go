package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

type Point struct {
	X, Y int
}

type FoldInstruction struct {
	axis  string
	value int
}

func readInput(filename string) ([]Point, []FoldInstruction) {
	var points []Point
	var instructions []FoldInstruction

	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	parseInstructions := false
	for _, line := range lines {
		if line == "" {
			parseInstructions = true
			continue
		}
		if parseInstructions {
			r := regexp.MustCompile(`fold\salong\s([xy]{1})=(\d+)`)
			matches := r.FindAllStringSubmatch(line, -1)
			value, _ := strconv.Atoi(matches[0][2])
			instructions = append(instructions, FoldInstruction{matches[0][1], value})
		} else {
			chars := strings.Split(line, ",")
			x, _ := strconv.Atoi(chars[0])
			y, _ := strconv.Atoi(chars[1])
			points = append(points, Point{x, y})
		}
	}

	return points, instructions
}

func countUnique(points []Point) int {
	unique := make(map[Point]bool)
	for _, point := range points {
		unique[point] = true
	}
	return len(unique)
}

func findMax(points []Point, axis string) int {
	max := 0
	for _, point := range points {
		if axis == "x" {
			if point.X > max {
				max = point.X
			}
		} else {
			if point.Y > max {
				max = point.Y
			}
		}
	}
	return max
}

func findNewPosition(position int, axisPosition int, size int) int {
	positionInFoldingGrid := position - (axisPosition + 1)
	newPosition := (axisPosition - 1) - positionInFoldingGrid
	return newPosition
}

func foldGrid(points []Point, instruction FoldInstruction) []Point {
	var newPoints []Point
	maxX := findMax(points, "x")
	maxY := findMax(points, "y")

	// Sort points into the folding points and the new points
	for _, point := range points {
		if instruction.axis == "x" {
			if point.X > instruction.value {
				foldedPoint := Point{findNewPosition(point.X, instruction.value, maxX+1), point.Y}
				newPoints = append(newPoints, foldedPoint)
			} else {
				newPoints = append(newPoints, point)
			}
		} else {
			if point.Y > instruction.value {
				foldedPoint := Point{point.X, findNewPosition(point.Y, instruction.value, maxY+1)}
				newPoints = append(newPoints, foldedPoint)
			} else {
				newPoints = append(newPoints, point)
			}
		}
	}

	return newPoints
}

func Part1(points []Point, instructions []FoldInstruction) int {
	points = foldGrid(points, instructions[0])

	return countUnique(points)
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day13"))
	fmt.Print("======================\n\n")

	examplePoints, exampleInstructions := readInput("example.txt")
	inputPoints, inputInstructions := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | How many dots are visible after completing just the first fold instruction on your transparent paper?")
	exampleResultPart1 := strconv.Itoa(Part1(examplePoints, exampleInstructions))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(inputPoints, inputInstructions))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	// fmt.Println("* Part 2 | What is the first step during which all octopuses flash?")
	// exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	// fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	// resultPart2 := strconv.Itoa(Part2(input))
	// fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
