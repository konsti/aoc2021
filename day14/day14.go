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
	template string
}

func readInput(filename string) ([]ElementPair, map[string]string) {
	var polymerTemplate []ElementPair
	pairInsertionRules := make(map[string]string)

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
		} else {
			for index, char := range line {
				if index != len(line)-1 {
					polymerTemplate = append(polymerTemplate, ElementPair{string(char) + string(line[index+1])})
				}
			}
		}
	}

	return polymerTemplate, pairInsertionRules
}

func countPairs(template []ElementPair, rules map[string]string, steps int) map[string]int {
	countMap := make(map[string]int)

	for _, pair := range template {
		countMap[pair.template]++
	}

	// fmt.Println(countMap)

	for i := 0; i < steps; i++ {
		newCountMap := make(map[string]int)
		for pair, count := range countMap {
			newCountMap[string(pair[0])+rules[pair]] += count
			newCountMap[rules[pair]+string(pair[1])] += count
		}
		countMap = newCountMap
		// fmt.Println(countMap)
	}

	return countMap
}

type Count struct {
	Element string
	Value   int
}
type CountList []Count

func (p CountList) Len() int           { return len(p) }
func (p CountList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p CountList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func countPairsToCountList(template []ElementPair, countPairs map[string]int) CountList {
	countElements := make(map[string]int)

	for pair, value := range countPairs {
		countElements[string(pair[0])] += value
	}
	lastTemplateElement := string(template[len(template)-1].template[1])
	countElements[lastTemplateElement]++

	// Sort the final count map and calculate the result
	countList := make(CountList, len(countElements))
	i := 0
	for key, value := range countElements {
		countList[i] = Count{key, value}
		i++
	}
	sort.Sort(countList)

	return countList
}

func Part1(template []ElementPair, rules map[string]string) int {
	countPairs := countPairs(template, rules, 10)
	countList := countPairsToCountList(template, countPairs)

	fmt.Println(countList)

	return countList[len(countList)-1].Value - countList[0].Value
}

func Part2(template []ElementPair, rules map[string]string) int {
	countPairs := countPairs(template, rules, 40)
	countList := countPairsToCountList(template, countPairs)

	fmt.Println(countList)

	return countList[len(countList)-1].Value - countList[0].Value
}

func main() {
	fmt.Println(color.Purple("Advent of Code - Day14"))
	fmt.Print("======================\n\n")

	exampleTemplate, exampleRules := readInput("example.txt")
	inputTemplate, inputRules := readInput("input.txt")

	// Part 1

	fmt.Println("* Part 1 | What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?")
	exampleResultPart1 := strconv.Itoa(Part1(exampleTemplate, exampleRules))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart1))

	resultPart1 := strconv.Itoa(Part1(inputTemplate, inputRules))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart1))

	// Part 2

	fmt.Println("* Part 2 | What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?")
	exampleResultPart2 := strconv.Itoa(Part2(exampleTemplate, exampleRules))
	fmt.Printf(color.Yellow("[Example Input]: %s \n"), color.Teal(exampleResultPart2))

	resultPart2 := strconv.Itoa(Part2(inputTemplate, inputRules))
	fmt.Printf(color.Green("[Real Input]: %s \n\n"), color.Teal(resultPart2))
}
