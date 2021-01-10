package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// algorithm to transform number
func transform(transformerNumber int, subjectNumber int) int {
	salt := 20201227

	return (transformerNumber * subjectNumber) % salt
}

// get loop size based on known public key
func loopSize(publicKey int, subjectNumber int) int {
	transformed := 1
	loopSize := 0

	for {
		transformed = transform(transformed, subjectNumber)
		loopSize++

		if transformed == publicKey {
			break
		}
	}

	return loopSize
}

// get encrypt key based on loop size
func encryptionKey(subjectNumber int, loopSize int) int {
	transformed := 1
	for i := 0; i < loopSize; i++ {
		transformed = transform(transformed, subjectNumber)
	}

	return transformed
}

// parse keys from input file
func parsePublicKeys(inputFilePath string) []int {
	file, _ := os.Open(inputFilePath)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInt, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, lineInt)
	}

	return lines
}

// solving the challenge: https://adventofcode.com/2020/day/25
func solve(inputFilePath string) {
	keys := parsePublicKeys(inputFilePath)
	cardPublicKey := keys[0]
	doorPublicKey := keys[1]

	subjectNumber := 7
	cardLoopSize := loopSize(cardPublicKey, subjectNumber)
	doorLoopSize := loopSize(doorPublicKey, subjectNumber)

	key := encryptionKey(doorPublicKey, cardLoopSize)
	key2 := encryptionKey(cardPublicKey, doorLoopSize)

	if key == key2 {
		fmt.Println("Secret key is: ", key)
	} else {
		fmt.Println("Bad input / algo")
	}
}

func main() {
	// read input file
	inputFilePath := os.Args[1:]
	if len(inputFilePath) < 1 {
		log.Fatal("Missing input file")
	}

	solve(inputFilePath[0])
}
