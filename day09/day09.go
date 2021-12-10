package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

type Point struct {
	number     int
	neighbours []*Point
}

func setNeighbours(input [][]Point) [][]Point {
	maxX := len(input[0])
	maxY := len(input)

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			point := &input[y][x]

			if x == 0 {
				if y == 0 {
					point.neighbours = append(point.neighbours, &input[y+1][x], &input[y][x+1])
				} else if y == maxY-1 {
					point.neighbours = append(point.neighbours, &input[y-1][x], &input[y][x+1])
				} else {
					point.neighbours = append(point.neighbours, &input[y-1][x], &input[y][x+1], &input[y+1][x])
				}
			} else if x == maxX-1 {
				if y == 0 {
					point.neighbours = append(point.neighbours, &input[y][x-1], &input[y+1][x])
				} else if y == maxY-1 {
					point.neighbours = append(point.neighbours, &input[y-1][x], &input[y][x-1])
				} else {
					point.neighbours = append(point.neighbours, &input[y-1][x], &input[y][x-1], &input[y+1][x])
				}
			} else {
				if y == 0 {
					point.neighbours = append(point.neighbours, &input[y][x-1], &input[y][x+1], &input[y+1][x])
				} else if y == maxY-1 {
					point.neighbours = append(point.neighbours, &input[y][x-1], &input[y][x+1], &input[y-1][x])
				} else {
					point.neighbours = append(point.neighbours, &input[y][x-1], &input[y][x+1], &input[y-1][x], &input[y+1][x])
				}
			}
		}
	}

	return input
}

func (point *Point) allNeighboursBigger() bool {
	for _, neighbour := range point.neighbours {
		if neighbour.number <= point.number {
			return false
		}
	}
	return true
}

func contains(slice []*Point, item *Point) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func readInput(filename string) [][]Point {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	var input [][]Point

	for _, line := range lines {
		var row []Point
		numbers := strings.Split(line, "")
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			logging.FailOnError(err, "Error converting string to int")
			row = append(row, Point{number: num})
		}
		input = append(input, row)
	}

	setNeighbours(input)
	return input
}

func Part1(input [][]Point) int {
	var lowPoints []int
	sumRiskLevel := 0

	for _, row := range input {
		for _, point := range row {
			if point.allNeighboursBigger() {
				lowPoints = append(lowPoints, point.number)
			}
		}
	}

	for _, point := range lowPoints {
		sumRiskLevel += point + 1
	}

	return sumRiskLevel
}

func fillBasin(basin []*Point, point *Point) []*Point {
	basin = append(basin, point)

	for index, neighbour := range point.neighbours {
		if !contains(basin, point.neighbours[index]) && neighbour.number != 9 {
			basin = fillBasin(basin, point.neighbours[index])
		}
	}

	return basin
}

func Part2(input [][]Point) int {
	var basins [][]*Point
	var lowPoints []Point

	for _, row := range input {
		for _, point := range row {
			if point.allNeighboursBigger() {
				lowPoints = append(lowPoints, point)
			}
		}
	}

	for _, point := range lowPoints {
		basin := fillBasin([]*Point{}, &point)
		basins = append(basins, basin)
	}

	var basinSizes []int

	for _, basin := range basins {
		basinSizes = append(basinSizes, len(basin)-1)
	}
	sort.Ints(basinSizes)

	result := 1

	for _, size := range basinSizes[len(basinSizes)-3:] {
		result *= size
	}

	return result
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

	fmt.Println("* Part 2 | What do you get if you multiply together the sizes of the three largest basins?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
