package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/commons/helpers"
)

func main() {
	lines, err := helpers.ReadFileIntoLines("day2")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	horCtr := 0
	verCtr := 0
	aim := 0
	for i, line := range lines {
		slice := strings.Split(line, " ")
		sliceAsInt, err := helpers.ToInt(slice[1])

		switch slice[0] {
		case "forward":
			horCtr += sliceAsInt
			verCtr += (aim * sliceAsInt)
		case "down":
			aim += sliceAsInt
		case "up":
			aim -= sliceAsInt
		default:
			if err != nil {
				fmt.Println(err, i)
			}
		}
	}
	fmt.Println(horCtr * verCtr)
}
