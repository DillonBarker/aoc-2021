package main

import (
	"aoc/commons/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	matrix := getMatrix()
	lowPoints := checkSurrounding(matrix)
	fmt.Println(lowPoints)
	fmt.Println(sum(lowPoints) + len(lowPoints))
}

func getMatrix() [][]int {
	input, _ := os.Open("./day9/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	matrix := [][]int{}
	for i := 0; i < 100; i++ {
		line := []int{}
		sc.Scan()
		array := strings.Split(sc.Text(), "")
		for _, each := range array {
			int, _ := helpers.ToInt(each)
			line = append(line, int)
		}
		matrix = append(matrix, line)
	}
	return matrix
}

func checkSurrounding(matrix [][]int) []int {
	numbers := []int{}
	for y, line := range matrix {
		for x, number := range line {
			// up
			if y != 0 {
				if number > matrix[y-1][x] {
					continue
				}
			}
			// down
			if y < len(matrix)-1 {
				if number > matrix[y+1][x] {
					continue
				}
			}
			// left
			if x != 0 {
				if number > matrix[y][x-1] {
					continue
				}
			}
			// right
			if x < len(matrix[0])-1 {
				if number > matrix[y][x+1] {
					continue
				}
			}
			if number != 9 {
				numbers = append(numbers, number)
			}

		}
	}
	return numbers
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
