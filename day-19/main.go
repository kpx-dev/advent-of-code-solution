package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseInput(inputFilePath string) []string {
	file, _ := os.Open(inputFilePath)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// lineInt, _ := strconv.Atoi(scanner.Text())
		line := scanner.Text()
		fmt.Println(line)
		if len(line) == 0 {
			fmt.Println("end of rule")
		}
		// lines = append(lines)
	}

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
