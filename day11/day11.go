package main

import (
	"fmt"
	"strconv"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

type Coordinates struct {
	X, Y int
}

type DumboOctopus struct {
	EnergyLevel int
	X, Y        int
	Neighbors   []*DumboOctopus
}

func (octopus DumboOctopus) getNeighbors(grid map[Coordinates]*DumboOctopus) []*DumboOctopus {
	neighbors := []*DumboOctopus{}

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}

			neighbor := grid[Coordinates{octopus.X + x, octopus.Y + y}]
			if neighbor != nil {
				neighbors = append(neighbors, neighbor)
			}
		}
	}

	return neighbors
}

func (octopus *DumboOctopus) IncreaseEnergyLevel() *DumboOctopus {
	octopus.EnergyLevel++

	// Only flash once, so at 10
	if octopus.EnergyLevel == 10 {
		// Flash
		for index, neighbor := range octopus.Neighbors {
			octopus.Neighbors[index] = neighbor.IncreaseEnergyLevel()
		}
	}

	return octopus
}

func readInput(filename string) []*DumboOctopus {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	var dumboOctopuses []*DumboOctopus
	grid := make(map[Coordinates]*DumboOctopus)

	for y, line := range lines {
		for x, char := range line {
			energyLevel, err := strconv.Atoi(string(char))
			logging.FailOnError(err, "Error converting char to int")

			octopus := DumboOctopus{
				EnergyLevel: energyLevel,
				X:           x,
				Y:           y,
				Neighbors:   []*DumboOctopus{},
			}

			grid[Coordinates{x, y}] = &octopus
			dumboOctopuses = append(dumboOctopuses, &octopus)
		}
	}

	for index, octopus := range dumboOctopuses {
		dumboOctopuses[index].Neighbors = octopus.getNeighbors(grid)
	}

	return dumboOctopuses
}

// Returns the number of flashes
func step(octopuses []*DumboOctopus) int {
	flashes := 0
	for index, octopus := range octopuses {
		octopuses[index] = octopus.IncreaseEnergyLevel()
	}
	for _, octopus := range octopuses {
		if octopus.EnergyLevel > 9 {
			flashes++
			octopus.EnergyLevel = 0
		}
	}

	return flashes
}

func Part1(input []*DumboOctopus) int {
	totalFlashes := 0
	for i := 0; i < 100; i++ {
		flashes := step(input)
		totalFlashes += flashes
	}
	return totalFlashes
}

func Part2(input []DumboOctopus) int {
	return 0
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day11"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | How many total flashes are there after 100 steps?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	// fmt.Println("* Part 2 | What is the middle score?")
	// exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	// fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	// resultPart2 := strconv.Itoa(Part2(input))
	// fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
