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

func readInput(filename string) []SignalPattern {
	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	var signalPatterns []SignalPattern

	for _, line := range lines {
		parts := strings.Split(line, "|")
		signalPatterns = append(signalPatterns, SignalPattern{
			uniqueSignalPatterns: strings.Fields(strings.TrimSpace(parts[0])),
			fourDigitOutputValue: strings.Fields(strings.TrimSpace(parts[1])),
		})
	}

	return signalPatterns
}

type wiringCandidates []rune

func (c wiringCandidates) Contains(r rune) bool {
	for _, candidate := range c {
		if candidate == r {
			return true
		}
	}
	return false
}

func (c wiringCandidates) Exclude(e wiringCandidates) wiringCandidates {
	var result wiringCandidates
	for _, candidate := range c {
		if !e.Contains(candidate) {
			result = append(result, candidate)
		}
	}
	return result
}

func ExistsInAll(candidates ...wiringCandidates) wiringCandidates {
	var result wiringCandidates
	var test rune
	for _, r := range candidates[0] {
		test = r
		exists := true
		for _, c := range candidates {
			if !c.Contains(test) {
				exists = false
			}
		}
		if exists {
			result = append(result, test)
		}
	}
	return result
}

// The sort algorithm is taken from https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

// Number of active wires for each digit
// 0 = 6
// 1 = 2 unique!
// 2 = 5
// 3 = 5
// 4 = 4 unique!
// 5 = 5
// 6 = 6
// 7 = 3 unique!
// 8 = 7 unique!
// 9 = 6
type SignalPattern struct {
	uniqueSignalPatterns []string
	fourDigitOutputValue []string
	displayWiring        map[string]string
}

func findDisplayWiring(pattern SignalPattern) SignalPattern {
	var candidates = map[string]wiringCandidates{}

	count5 := 0
	count6 := 0
	for _, uniqueSignalPattern := range pattern.uniqueSignalPatterns {
		var key string
		if len(uniqueSignalPattern) == 5 {
			key = "5" + strconv.Itoa(count5)
			count5++
		} else if len(uniqueSignalPattern) == 6 {
			key = "6" + strconv.Itoa(count6)
			count6++
		} else {
			key = strconv.Itoa(len(uniqueSignalPattern)) + "0"
		}

		candidates[key] = append(candidates[key], []rune(uniqueSignalPattern)...)
	}

	hCandidates := ExistsInAll(candidates["50"], candidates["51"], candidates["52"])
	top := candidates["30"].Exclude(candidates["20"])
	bottom := hCandidates.Exclude(top).Exclude(candidates["40"])
	middle := hCandidates.Exclude(bottom).Exclude(top)
	rightCandidates := candidates["20"]
	topleft := candidates["40"].Exclude(rightCandidates).Exclude(middle)
	bottomleft := candidates["70"].Exclude(rightCandidates).Exclude(hCandidates).Exclude(topleft)
	var topright wiringCandidates

	// Find number 6
	if candidates["60"].Contains(middle[0]) && candidates["60"].Contains(bottomleft[0]) {
		topright = rightCandidates.Exclude(candidates["60"])
	} else if candidates["61"].Contains(middle[0]) && candidates["61"].Contains(bottomleft[0]) {
		topright = rightCandidates.Exclude(candidates["61"])
	} else {
		topright = rightCandidates.Exclude(candidates["62"])
	}

	bottomright := candidates["20"].Exclude(topright)

	// fmt.Println(top, middle, bottom, topleft, bottomleft, topright, bottomright)

	newDisplayWiring := map[string]string{
		"0": SortString(string(top[0]) + string(topleft[0]) + string(topright[0]) + string(bottomleft[0]) + string(bottomright[0]) + string(bottom[0])),
		"1": SortString(string(topright[0]) + string(bottomright[0])),
		"2": SortString(string(top[0]) + string(topright[0]) + string(middle[0]) + string(bottomleft[0]) + string(bottom[0])),
		"3": SortString(string(top[0]) + string(topright[0]) + string(middle[0]) + string(bottomright[0]) + string(bottom[0])),
		"4": SortString(string(topleft[0]) + string(middle[0]) + string(topright[0]) + string(bottomright[0])),
		"5": SortString(string(top[0]) + string(topleft[0]) + string(middle[0]) + string(bottomright[0]) + string(bottom[0])),
		"6": SortString(string(top[0]) + string(topleft[0]) + string(middle[0]) + string(bottomleft[0]) + string(bottomright[0]) + string(bottom[0])),
		"7": SortString(string(top[0]) + string(topright[0]) + string(bottomright[0])),
		"8": SortString(string(top[0]) + string(topleft[0]) + string(topright[0]) + string(bottomleft[0]) + string(bottomright[0]) + string(bottom[0]) + string(middle[0])),
		"9": SortString(string(top[0]) + string(topleft[0]) + string(topright[0]) + string(bottomright[0]) + string(bottom[0]) + string(middle[0])),
	}

	pattern.displayWiring = reverseMap(newDisplayWiring)

	// fmt.Println(pattern)
	// fmt.Println(pattern.displayWiring)

	return pattern
}

func Part1(input []SignalPattern) int {
	count := 0

	for _, signalPattern := range input {
		for _, digit := range signalPattern.fourDigitOutputValue {
			if len(digit) == 2 || len(digit) == 4 || len(digit) == 3 || len(digit) == 7 {
				count++
			}
		}
	}

	return count
}

func Part2(input []SignalPattern) int {
	sum := 0

	for _, signalPattern := range input {
		displayWiring := findDisplayWiring(signalPattern)
		number := ""
		for _, digit := range signalPattern.fourDigitOutputValue {
			number += displayWiring.displayWiring[SortString(digit)]
		}
		numberInt, err := strconv.Atoi(number)
		logging.FailOnError(err, "could not convert number to int")
		sum += numberInt
	}

	return sum
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day8"))
	fmt.Print("======================\n\n")

	exampleInput := readInput("example.txt")
	input := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | In the output values, how many times do digits 1, 4, 7, or 8 appear?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	fmt.Println("* Part 2 | What do you get if you add up all of the output values?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleInput))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(input))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
