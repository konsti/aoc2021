package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

func readInput(filename string) []string {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Failed to read input file")

	return lines
}

func parsePart1Boards(lines []string) []*Board {
	// Each board has 5 lines and an empty line at the end (only the last board has no empty line)
	numberOfBoards := (len(lines) + 1) / 6

	var boards []*Board
	stringBoards := make([][]string, numberOfBoards)

	// The boards are defined after the third line of the input
	iterator := 0
	for _, line := range lines {
		if line == "" {
			iterator++
			continue
		}

		stringBoards[iterator] = append(stringBoards[iterator], line)
	}

	for _, board := range stringBoards {
		numberString := strings.Join(board, " ")
		var numbers []int

		for _, number := range strings.Fields(numberString) {
			numberInt, err := strconv.Atoi(number)
			logging.FailOnError(err, "Failed to convert string to int")
			numbers = append(numbers, numberInt)
		}

		board, err := NewBoard(numbers)
		logging.FailOnError(err, "Failed to create board")

		boards = append(boards, board)
	}

	return boards
}

func Part1(lines []string) int {
	// The game numbers are defined in the first line of the input
	gameNumbers := strings.Split(lines[0], ",")

	// The game boards are defined in the rest of the lines
	boards := parsePart1Boards(lines[2:])

	for _, number := range gameNumbers {
		numberInt, err := strconv.Atoi(number)
		logging.FailOnError(err, "Failed to convert gamenumber to int")

		for i := 0; i < len(boards); i++ {
			if _, ok := boards[i].byNumber[numberInt]; ok {
				boards[i].byNumber[numberInt].checked = true
			}

			if boards[i].HasWon() {
				fmt.Printf("Winner Number: %v\n", color.Teal(number))
				return boards[i].GetScore(numberInt)
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 4"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | What will your final score be if you choose that board?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))
}
