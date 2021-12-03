package main

import (
	"fmt"
	"log"

	"aoc/commons/helpers"
)

var zeroByte byte = 48
var oneByte byte = 49

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func main() {
	bits, err := helpers.ReadFile("day3")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	gamma := [5]int{1, 0, 1, 1, 0}

	fmt.Println(gamma, bits)
	for i, digit := range gamma {
		if digit == 49 {
			// remove bit in bits which does not have a 1 at this position
			for x, bit := range bits {
			}
		}
		if digit == 48 {
			// remove bit in bits which does not have a 0 at this position
		}
		if err != nil {
			log.Fatalf("%d: %s", i, err)
		}
	}
}
