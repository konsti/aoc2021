package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

type Point struct {
	X, Y, count int
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

func dotProduct(v1, v2 Vector) int {
	return v1.X*v2.X + v1.Y*v2.Y
}

func magnitude(vector Vector) float64 {
	return math.Sqrt(math.Pow(float64(vector.B.Y)-float64(vector.A.Y), 2) + math.Pow(float64(vector.B.X)-float64(vector.A.X), 2))
}

func normalize(vector Vector) Vector {
	magnitude := magnitude(vector)
	return Vector{vector.X / int(magnitude), vector.Y / int(magnitude), vector.A, vector.B}
}

// Implemented after math in https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect
func checkForOverlap(r, s Vector) string {
	p := r.A
	q := s.A
	vectorPQ := pointsToVector(p, q)

	uNumerator := crossProduct(vectorPQ, r)
	denominator := crossProduct(r, s)

	if uNumerator == 0 && denominator == 0 {
		// The lines are collinear

		t0 := float64(dotProduct(vectorPQ, r)) / float64(dotProduct(r, r))
		t1 := t0 + float64(dotProduct(s, r))/float64(dotProduct(r, r))

		if t0 >= 0 && t0 <= 1 || t1 >= 0 && t1 <= 1 {
			fmt.Println(color.Blue("collinear-overlapping", r, s, t0, t1))
			return "collinear-overlapping"
		}

		return "collinear-disjoint"
	}

	if denominator == 0 {
		// The lines are parallel
		return "parallel"
	}

	u := uNumerator / denominator
	t := crossProduct(vectorPQ, s) / denominator

	if u >= 0 && u <= 1 && t >= 0 && t <= 1 {
		fmt.Println(color.Blue("intersection", r, s, u, t))
		return "intersection"
	}

	return "no-intersection"
}

func readInput(filename string) []Vector {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Failed to read input file")

	var vectors []Vector

	for _, line := range lines {
		r := regexp.MustCompile(`(\d*),(\d*)\s->\s(\d*),(\d*)`)
		matches := r.FindAllStringSubmatch(line, -1)

		x1, _ := strconv.Atoi(matches[0][1])
		y1, _ := strconv.Atoi(matches[0][2])
		x2, _ := strconv.Atoi(matches[0][3])
		y2, _ := strconv.Atoi(matches[0][4])

		vectors = append(vectors, pointsToVector(Point{x1, y1, 0}, Point{x2, y2, 0}))
	}

	return vectors
}

func markPoints(grid [][]int, vector Vector) {
	currentPoint := vector.A
	normalizedVector := normalize(vector)

	// Mark the first point
	grid[currentPoint.X][currentPoint.Y]++

	for {
		currentPoint.X += normalizedVector.X
		currentPoint.Y += normalizedVector.Y
		grid[currentPoint.X][currentPoint.Y]++
		if currentPoint.X == vector.B.X && currentPoint.Y == vector.B.Y {
			break
		}
	}
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[j][i])
		}
		fmt.Println()
	}
}

func Part1(input []Vector, showGrid bool) int {
	var straightVectors []Vector
	maxX := 0
	maxY := 0

	// Filter vectors that are horizontal or vertical
	for _, vector := range input {
		if vector.X == 0 || vector.Y == 0 {
			straightVectors = append(straightVectors, vector)
		}
	}

	// Find the max X and Y values
	for _, vector := range straightVectors {
		if vector.A.X > maxX {
			maxX = vector.A.X
		}
		if vector.A.Y > maxY {
			maxY = vector.A.Y
		}
		if vector.B.X > maxX {
			maxX = vector.B.X
		}
		if vector.B.Y > maxY {
			maxY = vector.B.Y
		}
	}

	grid := make([][]int, maxX+1)
	for i := range grid {
		grid[i] = make([]int, maxY+1)
	}

	for _, vector := range straightVectors {
		markPoints(grid, vector)
	}

	if showGrid {
		printGrid(grid)
	}

	// Count the number of points that have been marked above 1
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 5"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | At how many points do at least two lines overlap?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput, true))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input, false))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))
}
