package main

import (
	"aoc/commons/helpers"
	"fmt"
	"strings"
)

func main() {
	input, _ := helpers.ReadFileIntoLines("day8")
	arr := [][]string{}
	for _, line := range input {
		seperated := strings.Split(line, "|")
		arr = append(arr, seperated)
	}
	ctr := 0
	for _, line := range arr {
		splitLine := strings.Split(line[1], " ")
		for _, entry := range splitLine[1:] {
			splitEntry := strings.Split(entry, "")
			if len(splitEntry) == 2 {
				// if contains(splitEntry, "c") && contains(splitEntry, "f") {
				ctr++
				// }
			}
			if len(splitEntry) == 3 {
				// if contains(splitEntry, "d") && contains(splitEntry, "a") && contains(splitEntry, "b") {
				ctr++
				// }
			}
			if len(splitEntry) == 4 {
				// if contains(splitEntry, "a") && contains(splitEntry, "e") && contains(splitEntry, "f") && contains(splitEntry, "b") {
				ctr++
				// }
			}
			if len(splitEntry) == 7 {
				if contains(splitEntry, "a") && contains(splitEntry, "b") && contains(splitEntry, "c") && contains(splitEntry, "d") && contains(splitEntry, "e") && contains(splitEntry, "f") && contains(splitEntry, "g") {
					ctr++
				}
			}
		}
	}
	fmt.Println(ctr)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/*
1 cf

4 eafb

7 dab

8 abcdefg


*/
