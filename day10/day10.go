package main

import (
	"fmt"
	"strconv"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"

	"github.com/lucsky/cuid"
)

type Delimiter struct {
	open  rune
	close rune
}

type Node struct {
	id        string
	delimiter Delimiter
	open      rune
	close     rune
	children  []*Node
	parent    *Node
}

func Delimiters() map[rune]Delimiter {
	return map[rune]Delimiter{
		'(': {'(', ')'},
		'[': {'[', ']'},
		'{': {'{', '}'},
		'<': {'<', '>'},
	}
}

func isOpening(r rune) bool {
	return r == '(' || r == '[' || r == '{' || r == '<'
}

func buildTree(input string) (*Node, rune) {
	rootNode := &Node{
		id: cuid.Slug(),
	}
	currentNode := rootNode

	for _, r := range input {
		// fmt.Println(currentNode, string(r))
		if isOpening(r) {
			node := Node{id: cuid.Slug(), delimiter: Delimiters()[r], open: r}
			currentNode.children = append(currentNode.children, &node)
			node.parent = currentNode
			currentNode = &node
		} else if r == currentNode.delimiter.close {
			currentNode.close = r
			currentNode = currentNode.parent
		} else {
			return nil, r
		}
	}

	return rootNode, ' '
}

func readInput(filename string) []string {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	return lines
}

func Part1(input []string) int {
	score := 0

	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	for _, line := range input {
		_, corruptRune := buildTree(line)
		if corruptRune != ' ' {
			score += scores[corruptRune]
		}
	}

	return score
}

func Part2(input []string) int {
	return 0
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day10"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | What is the total syntax error score for those errors?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	// fmt.Println("* Part 2 | What do you get if you multiply together the sizes of the three largest basins?")
	// exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	// fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	// resultPart2 := strconv.Itoa(Part2(input))
	// fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
