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
		if i == 0 {
			continue
		}

		lineAsInt, err := helpers.ToInt(line)
		previousLineAsInt, err := helpers.ToInt(lines[i-1])

		if err != nil {
			fmt.Println(err)
		}

		if lineAsInt > previousLineAsInt {
			ctr++
		}
	}
	fmt.Println(ctr)
}
