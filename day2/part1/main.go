package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/commons/helpers"
)

func main() {
	lines, err := helpers.ReadFile("day2")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	horCtr := 0
	verCtr := 0
	for i, line := range lines {
		slice := strings.Split(line, " ")
		sliceAsInt, err := helpers.ToInt(slice[1])

		switch slice[0] {
		case "forward":
			horCtr += sliceAsInt
		case "down":
			verCtr += sliceAsInt
		case "up":
			verCtr -= sliceAsInt
		}

		if err != nil {
			fmt.Println(err, i)
		}
	}
	fmt.Println(horCtr * verCtr)
}
