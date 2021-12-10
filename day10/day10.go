package main

import (
	"fmt"
	"sort"
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

// Tries to build the AST from the given input
// If chunks are not closed, it tries to close them and returns an array of used delimiters
// If some chunks close with an incorrect delimiter, it returns the first incorrect delimiter
func checkTree(input string) ([]rune, rune) {
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

	newDelimiters := []rune{}
	for ok := true; ok; ok = currentNode.close != currentNode.delimiter.close {
		newDelimiters = append(newDelimiters, currentNode.delimiter.close)
		currentNode = currentNode.parent
	}

	return newDelimiters, ' '
}

func readInput(filename string) []string {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	return lines
}

func Part1(input []string) int {
	score := 0

	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	for _, line := range input {
		_, corruptRune := checkTree(line)
		if corruptRune != ' ' {
			score += points[corruptRune]
		}
	}

	return score
}

func median(input []float64) float64 {
	sort.Float64s(input)
	median := float64(0)
	if len(input)%2 == 0 {
		median = (input[len(input)/2-1] + input[len(input)/2]) / 2
	} else {
		median = input[len(input)/2]
	}
	return median
}

func Part2(input []string) int {
	var scores []float64

	points := map[rune]float64{
		')': 1.0,
		']': 2.0,
		'}': 3.0,
		'>': 4.0,
	}

	for _, line := range input {
		newDelimiters, _ := checkTree(line)
		if newDelimiters != nil {
			score := 0.0
			for _, d := range newDelimiters {
				score = score*5.0 + points[d]
			}
			scores = append(scores, score)
		}
	}

	return int(median(scores))
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

	fmt.Println("* Part 2 | What is the middle score?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
