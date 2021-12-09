package main

import (
	"aoc/commons/helpers"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	crabs := getInput()
	optimalPos := medianCrab(crabs)
	fuel := 0
	for _, crab := range crabs {
		if crab != optimalPos {
			diff := diff(crab, optimalPos)
			fuel += diff
		}
	}
	fmt.Println(fuel)
}

func medianCrab(crabs []int) int {
	sortedCrabs := sortCrabs(crabs)

	middleCrab := len(sortedCrabs) / 2

	if middleCrab%2 != 0 {
		return sortedCrabs[middleCrab]
	} else {
		return (sortedCrabs[middleCrab-1] + sortedCrabs[middleCrab]) / 2
	}
}

func sortCrabs(crabs []int) []int { // assuming no tie
	sort.Ints(crabs)
	return crabs
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func getInput() []int {
	input, _ := os.Open("./day7/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	sc.Scan()
	inputs := strings.Split(sc.Text(), ",")
	nums := []int{}
	for _, input := range inputs {
		inputAsInt, _ := helpers.ToInt(input)
		nums = append(nums, inputAsInt)
	}
	return nums
}
