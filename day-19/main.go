package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var ruleSets []string
var testCases []string

func match(line string) bool {
	// build out regex
	return true
}

func parseTestCases(line string) {
	testCases = append(testCases, line)
}

func parseRuleSet(line string) {
	lineSplit := strings.Split(line, ":")
	// fmt.Println(lineSplit[0])

	if len(lineSplit) == 2 {
		// reached the definition rule. Ex: 4: "a"
	}
	fmt.Println(lineSplit[0])
}

func parseInput(inputFilePath string) []string {
	file, _ := os.Open(inputFilePath)
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	parseTestCasesInput := true

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			fmt.Println("end of rule")
			parseTestCasesInput = false
		}

		if parseTestCasesInput {
			parseTestCases(line)
		} else {
			// parse rule set
			parseRuleSet(line)
			// ruleSets = append(ruleSets, line)
		}
	}

	fmt.Println(ruleSets, testCases)
	return lines
}

// solving the challenge: https://adventofcode.com/2020/day/19
func solve(inputFilePath string) {
	parseInput(inputFilePath)

	matches := 0
	fmt.Println("Number of matches ", matches)
}

func main() {
	// read input file
	inputFilePath := os.Args[1:]
	if len(inputFilePath) < 1 {
		log.Fatal("Missing input file")
	}

	solve(inputFilePath[0])
}
