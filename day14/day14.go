package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/konsti/aoc2021/utils/color"
	"github.com/konsti/aoc2021/utils/input"
	"github.com/konsti/aoc2021/utils/logging"
)

type ElementPair struct {
	template   string
	generation int
}

func readInput(filename string) ([]ElementPair, map[string]string, map[string]bool) {
	var polymerTemplate []ElementPair
	pairInsertionRules := make(map[string]string)
	potentialElements := make(map[string]bool)

	lines, err := input.ReadLines(filename)
	logging.FailOnError(err, "Error reading input file")

	parseRules := false
	for _, line := range lines {
		if line == "" {
			parseRules = true
			continue
		}
		if parseRules {
			r := regexp.MustCompile(`(\w{2})\s->\s(\w{1})`)
			matches := r.FindAllStringSubmatch(line, -1)
			pairInsertionRules[matches[0][1]] = matches[0][2]
			potentialElements[matches[0][2]] = true
		} else {
			for index, char := range line {
				potentialElements[string(char)] = true
				if index != len(line)-1 {
					polymerTemplate = append(polymerTemplate, ElementPair{string(char) + string(line[index+1]), 0})
				}
			}
		}
	}

	return polymerTemplate, pairInsertionRules, potentialElements
}

type CountPair struct {
	Key   string
	Value int
}
type CountPairList []CountPair

func (p CountPairList) Len() int           { return len(p) }
func (p CountPairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p CountPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func polymerize(pair ElementPair, rules map[string]string, count map[string]int, stop int, isFirstElement bool) map[string]int {
	A := pair.template[0:1]
	B := pair.template[1:2]

	if pair.generation >= stop {
		// we only count the second element of each final generation, as pairs are overlapping
		count[B]++
		if isFirstElement {
			count[A]++
		}
		return count
	}

	newPair1 := ElementPair{A + rules[A+B], pair.generation + 1}
	newPair2 := ElementPair{rules[A+B] + B, pair.generation + 1}

	polymerize(newPair1, rules, count, stop, isFirstElement)
	polymerize(newPair2, rules, count, stop, false)

	return count
}

func Part1(template []ElementPair, rules map[string]string, elementSet map[string]bool) int {
	countElementMap := make(map[string]int)
	for key := range elementSet {
		countElementMap[key] = 0
	}

	for index, pair := range template {
		polymerize(pair, rules, countElementMap, 10, index == 0)
	}

	fmt.Println(countElementMap)

	// Sort the final count map and calculate the result
	countList := make(CountPairList, len(countElementMap))
	i := 0
	for key, value := range countElementMap {
		countList[i] = CountPair{key, value}
		i++
	}
	sort.Sort(countList)

	return countList[len(countList)-1].Value - countList[0].Value
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day14"))
	fmt.Print("======================\n\n")

	exampleTemplate, exampleRules, exampleElementSet := readInput("example.txt")
	inputTemplate, inputRules, inputElementSet := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleTemplate, exampleRules, exampleElementSet))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(inputTemplate, inputRules, inputElementSet))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	// fmt.Println("* Part 2 | What code do you use to activate the infrared thermal imaging camera system?")
	// exampleResultPart2 := strconv.Itoa(Part2(exampleTemplate, exampleRules))
	// fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	// resultPart2 := strconv.Itoa(Part2(inputTemplate, inputRules))
	// fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
