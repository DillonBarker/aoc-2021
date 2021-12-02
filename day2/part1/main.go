package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"aoc/commons/read"
)

func toInt(line string) (int, error) {
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
	for i, line := range lines {
		slice := strings.Split(line, " ")
		sliceAsInt, err := toInt(slice[1])

		if slice[0] == "forward" {
			horCtr = horCtr + sliceAsInt
		}

		if slice[0] == "down" {
			verCtr += sliceAsInt
		}

		if slice[0] == "up" {
			verCtr += sliceAsInt
		}

		if err != nil {
			fmt.Println(err, i)
		}
	}
	fmt.Println(horCtr * verCtr)
}
