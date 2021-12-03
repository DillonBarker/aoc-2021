package main

import (
	"fmt"
	"log"
	"strconv"

	"aoc/commons/helpers"
)

func main() {
	lines, err := helpers.ReadFile("day3")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var gammaAns string
	var epsilonAns string
	for i := 0; i < 5; i++ {
		oneCtr := 0
		zeroCtr := 0
		for x, line := range lines {
			var zeroByte byte = 48
			var oneByte byte = 49
			if line[i] == zeroByte {
				zeroCtr++
			}
			if line[i] == oneByte {
				oneCtr++
			}
			if err != nil {
				log.Fatalf("for loop: %s with %d", err, x)
			}
		}
		if oneCtr > zeroCtr {
			gammaAns += "1"
			epsilonAns += "0"
		}
		if zeroCtr > oneCtr {
			gammaAns += "0"
			epsilonAns += "1"
		}

	}
	gamma, err := strconv.ParseInt(gammaAns, 2, 64)
	epsilon, err := strconv.ParseInt(epsilonAns, 2, 64)
	fmt.Println(gammaAns)
	fmt.Println(gamma * epsilon)
}
