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

func parseBoards(lines []string) []*Board {
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

func parseGameNumbers(gameNumbers []string) []int {
	var numbers []int

	for _, number := range gameNumbers {
		numberInt, err := strconv.Atoi(number)
		logging.FailOnError(err, "Failed to convert gamenumber to int")

		numbers = append(numbers, numberInt)
	}

	return numbers
}

func Part1(lines []string) int {
	// The game numbers are defined in the first line of the input
	gameNumbers := parseGameNumbers(strings.Split(lines[0], ","))

	// The game boards are defined in the rest of the lines
	boards := parseBoards(lines[2:])

	for _, number := range gameNumbers {
		for i := 0; i < len(boards); i++ {
			boards[i].Play(number)

			if boards[i].HasWon() {
				fmt.Printf("Winning Number: %v\n", color.Teal(number))
				return boards[i].GetScore(number)
			}
		}
	}

	return 0
}

func contains(s []*Board, e *Board) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Part2(lines []string) int {
	gameNumbers := parseGameNumbers(strings.Split(lines[0], ","))
	boards := parseBoards(lines[2:])

	var boardsInOrderOfWinning []*Board
	winningNumber := 0

	for _, number := range gameNumbers {
		for i := 0; i < len(boards); i++ {
			boards[i].Play(number)

			if boards[i].HasWon() && !contains(boardsInOrderOfWinning, boards[i]) {
				boardsInOrderOfWinning = append(boardsInOrderOfWinning, boards[i])
			}
		}
		if len(boardsInOrderOfWinning) == len(boards) {
			winningNumber = number
			break
		}
	}

	winningBoard := boardsInOrderOfWinning[len(boardsInOrderOfWinning)-1]

	fmt.Printf("Winning Number: %v\n", color.Teal(winningNumber))

	return winningBoard.GetScore(winningNumber)
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day 4"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | What will your final score be if you choose the winning board?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	fmt.Println("* Part 2 | What is the final score of the last board that wins?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input))
	fmt.Printf(color.Green("[Real Input]:	%s \n"), color.Teal(resultPart2))
}
