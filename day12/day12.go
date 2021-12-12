package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
	"github.com/yourbasic/graph"
)

type Cave struct {
	id     string
	number int
	small  bool
	visits int
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// In order to solve the puzzle, we need to build a graph of the cave system.
// Then we can use path finding to find pathes through the graph.
func readInput(filename string) (*graph.Mutable, map[string]*Cave, map[int]*Cave) {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	caves := make(map[string]*Cave)
	caveDictornary := make(map[int]*Cave)

	// The yourbasic/graph package works on int graphes... To avoid
	// implementing a custom graph, we assign an int to every cave.
	count := 0
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if _, ok := caves[parts[0]]; !ok {
			cave := Cave{parts[0], count, IsLower(parts[0]), 0}
			caves[parts[0]] = &cave
			caveDictornary[count] = &cave
			count++
		}
		if _, ok := caves[parts[1]]; !ok {
			cave := Cave{parts[1], count, IsLower(parts[1]), 0}
			caves[parts[1]] = &cave
			caveDictornary[count] = &cave
			count++
		}
	}

	fmt.Println("\n------------------ Caves ------------------")
	for _, cave := range caves {
		fmt.Println("Cave ", cave.number, ": ", cave.id, " [small: ", cave.small, "]")
	}
	fmt.Println("-------------------------------------------\n")

	caveSystem := graph.New(len(caves))

	for _, line := range lines {
		parts := strings.Split(line, "-")
		caveSystem.AddBoth(caves[parts[0]].number, caves[parts[1]].number)
	}

	return caveSystem, caves, caveDictornary
}

// Using a recursive Depth First Traversal to find all paths through the graph.
func findAllPaths(caveSystem *graph.Mutable, caves map[string]*Cave, dict map[int]*Cave) [][]*Cave {
	// Reset visits
	for id := range caves {
		caves[id].visits = 0
	}

	var paths [][]*Cave
	paths = visitPaths(paths, caveSystem, caves["start"], caves["end"], []*Cave{caves["start"]}, dict)

	return paths
}

func visitPaths(all [][]*Cave, caveSystem *graph.Mutable, from *Cave, to *Cave, paths []*Cave, dict map[int]*Cave) [][]*Cave {
	// fmt.Println("Visiting path from: ", from.id, " to: ", to.id)

	if from.id == to.id {
		all = append(all, paths)
		// fmt.Print("---> Found path: ")
		// for _, path := range paths {
		// 	fmt.Print(path.id, " ")
		// }
		// fmt.Println()
		return all
	}

	from.visits++

	graph.Sort(caveSystem).Visit(from.number, func(toInt int, _ int64) (skip bool) {
		// fmt.Println(color.Green("Visiting: ", dict[toInt].id, " from: ", from.id))
		cave := dict[toInt]
		// Only big caves can be visited twice
		if cave.visits == 0 || !cave.small {
			paths = append(paths, dict[toInt])
			all = visitPaths(all, caveSystem, dict[toInt], to, paths, dict)
			paths = paths[:len(paths)-1]
		}
		return
	})

	from.visits--

	return all
}

func Part1(caveSystem *graph.Mutable, caves map[string]*Cave, dict map[int]*Cave) int {
	paths := findAllPaths(caveSystem, caves, dict)

	return len(paths)
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day11"))
	fmt.Print("======================\n")

	graph1, caves1, dict1 := readInput("example1.txt")
	graph2, caves2, dict2 := readInput("example2.txt")
	graph3, caves3, dict3 := readInput("example3.txt")
	graphInput, cavesInput, dictInput := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | How many paths through this cave system are there that visit small caves at most once?")
	exampleResult1Part1 := strconv.Itoa(Part1(graph1, caves1, dict1))
	fmt.Printf(color.Yellow("[Example Input 1]: %s \n"), color.Teal(exampleResult1Part1))

	exampleResult2Part1 := strconv.Itoa(Part1(graph2, caves2, dict2))
	fmt.Printf(color.Yellow("[Example Input 2]: %s \n"), color.Teal(exampleResult2Part1))

	exampleResult3Part1 := strconv.Itoa(Part1(graph3, caves3, dict3))
	fmt.Printf(color.Yellow("[Example Input 3]: %s \n"), color.Teal(exampleResult3Part1))

	resultPart1 := strconv.Itoa(Part1(graphInput, cavesInput, dictInput))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	// fmt.Println("* Part 2 | What is the first step during which all octopuses flash?")
	// exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	// fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	// resultPart2 := strconv.Itoa(Part2(input))
	// fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
