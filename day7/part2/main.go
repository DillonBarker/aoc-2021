package main

import (
	"aoc/commons/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	crabs := getInput()
	optimalPos := meanCrab(crabs)
	fuel := 0
	for _, crab := range crabs {
		if crab != optimalPos {
			diff := diff(crab, optimalPos)
			for i := 0; i < diff; i++ {
				fuel += i + 1
			}
		}
	}
	fmt.Println(fuel)
}

func meanCrab(crabs []int) int {
	n := len(crabs)
	sum := 0
	for i := 0; i < n; i++ {
		sum += (crabs[i])
	}
	avg := (float64(sum)) / (float64(n))
	avgInt := int(avg)
	fmt.Println(avg)
	return avgInt
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
