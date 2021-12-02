package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
	lines, err := read.ReadFile("day2")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	horCtr := 0
	verCtr := 0
	aim := 0

	for i, line := range lines {
		slice := strings.Split(line, " ")
		fmt.Println(i, slice[0], slice[1])

		sliceAsInt, err := getLineAsInt(slice[1])

		if slice[0] == "forward" {
			horCtr = horCtr + sliceAsInt
			verCtr = verCtr + (aim * sliceAsInt)
		}

		if slice[0] == "down" {
			aim += sliceAsInt
		}

		if slice[0] == "up" {
			aim += sliceAsInt
		}

		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(horCtr * verCtr)
}
