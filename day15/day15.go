package main

import (
	"fmt"
	"strconv"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
	"github.com/yourbasic/graph"
)

type Position struct {
	id        int
	X, Y      int
	riskLevel int
	Neighbors []*Position
	priority  int
	index     int
}

func (position *Position) getNeighbors(grid map[string]*Position) []*Position {
	neighbors := []*Position{}

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}

			neighbor := grid[fmt.Sprintf("%d,%d", position.X+x, position.Y+y)]
			if neighbor != nil {
				neighbors = append(neighbors, neighbor)
			}
		}
	}

	return neighbors
}

func getRiskLevel(lines []string, x, y int) int {
	squareSize := len(lines)

	riskLevel, err := strconv.Atoi(string(lines[y%squareSize][x%squareSize]))
	logging.FailOnError(err, "Error converting char to int")

	riskLevel = riskLevel + (x / squareSize) + (y / squareSize)
	if riskLevel > 9 {
		riskLevel = riskLevel - 9
	}

	return riskLevel
}

func readInput(filename string, multiplier int) []Position {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	var positions []Position
	grid := make(map[string]*Position)

	idCount := 0
	for y := 0; y < len(lines)*multiplier; y++ {
		for x := 0; x < len(lines)*multiplier; x++ {
			riskLevel := getRiskLevel(lines, x, y)

			position := Position{
				id: idCount,
				X:  x, Y: y,
				riskLevel: riskLevel,
				Neighbors: []*Position{},
				priority:  -1,
				index:     -1,
			}
			positions = append(positions, position)
			grid[fmt.Sprintf("%d,%d", x, y)] = &position
			idCount++
		}
	}

	for index, position := range positions {
		positions[index].Neighbors = position.getNeighbors(grid)
	}

	return positions
}

func buildGraph(positions []Position, withDiagonales bool) *graph.Mutable {
	graph := graph.New(len(positions))

	for _, position := range positions {
		for _, neighbor := range position.Neighbors {
			if withDiagonales {
				graph.AddBothCost(position.id, neighbor.id, int64(neighbor.riskLevel))
			} else if position.X == neighbor.X || position.Y == neighbor.Y {
				graph.AddBothCost(position.id, neighbor.id, int64(neighbor.riskLevel))
			}
		}
	}

	return graph
}

func pathRisk(positions []Position, path []int) int {
	totalRisk := 0

	for _, id := range path {
		if id != 0 {
			totalRisk += int(positions[id].riskLevel)
		}
	}

	return totalRisk
}

func Dijkstra(g *graph.Mutable, v, end int, positions []Position) []int {
	n := g.Order()
	dist := make([]int64, n)
	parent := make([]int, n)
	for i := range dist {
		dist[i], parent[i] = -1, -1
	}
	dist[v] = 0

	// Dijkstra's algorithm
	Q := emptyPrioQueue(dist)
	Q.Push(v)
	for Q.Len() > 0 {
		v := Q.Pop()
		g.Visit(v, func(w int, d int64) (skip bool) {
			if d < 0 {
				return
			}
			alt := dist[v] + int64(positions[v].riskLevel)
			switch {
			case dist[w] == -1:
				dist[w], parent[w] = alt, v
				Q.Push(w)
			case alt < dist[w]:
				dist[w], parent[w] = alt, v
				Q.Fix(w)
			}
			return
		})
	}

	var path []int
	for current := end; current != -1; current = parent[current] {
		path = append(path, current)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func printPath(positions []Position, path []int) {
	currentY := 0
	for _, position := range positions {
		if position.Y > currentY {
			fmt.Println()
			currentY = position.Y
		}
		if contains(path, position.id) {
			fmt.Print(color.Purple(int(position.riskLevel)))
		} else {
			fmt.Print(int(position.riskLevel))
		}
	}
}

func Part1(positions []Position, print bool) int {
	caveGraph := buildGraph(positions, false)

	path := Dijkstra(caveGraph, 0, len(positions)-1, positions)
	if print {
		printPath(positions, path)
	}

	return pathRisk(positions, path)
}

func Part2(positions []Position, print bool) int {
	caveGraph := buildGraph(positions, false)

	path := Dijkstra(caveGraph, 0, len(positions)-1, positions)
	if print {
		printPath(positions, path)
	}

	return pathRisk(positions, path)
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day15"))
	fmt.Print("======================\n")

	// Part 1

	example := readInput("example.txt", 1)
	input := readInput("input.txt", 1)

	fmt.Println("* Part 1 | What is the lowest total risk of any path from the top left to the bottom right?")
	exampleResult1Part1 := strconv.Itoa(Part1(example, true))
	fmt.Printf(color.Yellow("[Example Input 1]: %s \n"), color.Teal(exampleResult1Part1))

	resultPart1 := strconv.Itoa(Part1(input, false))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	example = readInput("example.txt", 5)
	input = readInput("input.txt", 5)

	fmt.Println("* Part 2 | HUsing the full map, what is the lowest total risk of any path from the top left to the bottom right?")
	exampleResultPart2 := strconv.Itoa(Part2(example, true))
	fmt.Printf(color.Yellow("[Example Input 1]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input, false))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
