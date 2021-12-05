package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

type Point struct {
	X, Y int
}

type Vector struct {
	Y, X int
	A, B Point
}

func pointsToVector(p1, p2 Point) Vector {
	return Vector{p2.X - p1.X, p2.Y - p1.Y, p1, p2}
}

func crossProduct(v1, v2 Vector) int {
	return v1.X*v2.Y - v1.Y*v2.X
}

func pointsEqual(p1, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func allEqual(slice []bool) bool {
	if len(slice) == 0 {
		return true
	}
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[0] {
			return false
		}
	}
	return true
}

// Implemented after math in https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect
func checkForOverlap(v1, v2 Vector) bool {
	vectorAB := pointsToVector(v1.A, v2.A)

	uNumerator := crossProduct(vectorAB, v1)
	denominator := crossProduct(v1, v2)

	println(uNumerator, denominator)

	if uNumerator == 0 && denominator == 0 {
		// The lines are collinear

		// Do the lines have the same endpoints?
		if pointsEqual(v1.A, v2.A) ||
			pointsEqual(v1.A, v2.B) ||
			pointsEqual(v1.B, v2.A) ||
			pointsEqual(v1.B, v2.B) {
			return true
		}

		return !allEqual([]bool{v2.A.X-v1.A.X < 0, v2.A.X-v1.B.X < 0, v2.B.X-v1.A.X < 0, v2.B.X-v1.B.X < 0}) ||
			!allEqual([]bool{v2.A.Y-v1.A.Y < 0, v2.A.Y-v1.B.Y < 0, v2.B.Y-v1.A.Y < 0, v2.B.Y-v1.B.Y < 0})
	}

	if denominator == 0 {
		// The lines are parallel
		return false
	}

	u := uNumerator / denominator
	t := crossProduct(vectorAB, v2) / denominator

	println("----> u: ", u, ", --->t: ", t)

	return u >= 0 && u <= 1 && t >= 0 && t <= 1
}

func readInput(filename string) []Vector {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Failed to read input file")

	var vectors []Vector

	for _, line := range lines {
		r := regexp.MustCompile(`(\d),(\d)\s->\s(\d),(\d)`)
		matches := r.FindAllStringSubmatch(line, -1)

		x1, _ := strconv.Atoi(matches[0][1])
		y1, _ := strconv.Atoi(matches[0][2])
		x2, _ := strconv.Atoi(matches[0][3])
		y2, _ := strconv.Atoi(matches[0][4])

		vectors = append(vectors, pointsToVector(Point{x1, y1}, Point{x2, y2}))
	}

	return vectors
}

func Part1(input []Vector) int {
	var straightVectors []Vector
	countOverlaps := 0

	// Filter vectors that are horizontal or vertical
	for _, vector := range input {
		if vector.X == 0 || vector.Y == 0 {
			straightVectors = append(straightVectors, vector)
		}
	}

	fmt.Println(straightVectors)

	for index, vector := range straightVectors {
		for i := index + 1; i < len(straightVectors); i++ {
			if checkForOverlap(vector, straightVectors[i]) {
				countOverlaps++
			}
		}
	}

	return countOverlaps
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 5"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	// input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | At how many points do at least two lines overlap?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	// resultPart1 := strconv.Itoa(Part1(input))
	// fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))
}
