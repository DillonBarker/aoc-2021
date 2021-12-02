package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"aoc/commons/read"
)

func toInt(str string) (int, error) {
	lineAsInt, err := strconv.Atoi(str)
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
		sliceAsInt, err := toInt(slice[1])

		if slice[0] == "forward" {
			horCtr += sliceAsInt
			verCtr += (aim * sliceAsInt)
		}

		if slice[0] == "down" {
			aim += sliceAsInt
		}

		if slice[0] == "up" {
			aim += sliceAsInt
		}

		if err != nil {
			fmt.Println(err, i)
		}
	}
	fmt.Println(horCtr * verCtr)
}
