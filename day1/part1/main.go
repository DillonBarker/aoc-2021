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
		if i == 0 {
			continue
		}

		lineAsInt, err := getLineAsInt(line)
		if err != nil {
			fmt.Println(err)
		}

		previousLineAsInt, err := getLineAsInt(lines[i-1])
		if err != nil {
			fmt.Println(err)
		}

		if lineAsInt > previousLineAsInt {
			ctr++
		}
	}
	fmt.Println(ctr)
}
