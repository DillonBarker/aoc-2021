package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"aoc/commons/read"
)

func getLineAsInt(line string) (int, error) {
	lineAsInt, err := strconv.Atoi(line)
	if err != nil {
		os.Exit(2)
	}
	return lineAsInt, err
}

func main() {
	lines, err := read.ReadFile("day1")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	ctr := 0
	for i, line := range lines {
		if i > 1996 {
			continue
		}

		lineAsInt, err := getLineAsInt(line)
		lineAsInt2, err := getLineAsInt(lines[i+1])
		lineAsInt3, err := getLineAsInt((lines[i+2]))

		lineTotal := lineAsInt + lineAsInt2 + lineAsInt3

		nextLineAsInt, err := getLineAsInt(lines[i+1])
		nextLineAsInt2, err := getLineAsInt(lines[i+2])
		nextLineAsInt3, err := getLineAsInt(lines[i+3])

		nextLineTotal := nextLineAsInt + nextLineAsInt2 + nextLineAsInt3

		if nextLineTotal > lineTotal {
			ctr++
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(ctr)
}
