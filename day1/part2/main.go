package main

import (
	"fmt"
	"log"

	"aoc/commons/helpers"
)

func main() {
	lines, err := helpers.ReadFileIntoLines("day1")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	ctr := 0
	for i, line := range lines {
		if i > 1996 {
			continue
		}

		lineAsInt, err := helpers.ToInt(line)
		lineAsInt2, err := helpers.ToInt(lines[i+1])
		lineAsInt3, err := helpers.ToInt((lines[i+2]))

		lineTotal := lineAsInt + lineAsInt2 + lineAsInt3

		nextLineAsInt, err := helpers.ToInt(lines[i+1])
		nextLineAsInt2, err := helpers.ToInt(lines[i+2])
		nextLineAsInt3, err := helpers.ToInt(lines[i+3])

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
