package main

import (
	"aoc/commons/helpers"
	"fmt"
	"strings"
)

const SIZE int = 1000

func main() {
	vents, _ := helpers.ReadFileIntoLines("day5")
	grid := ventGrid()
	for _, line := range vents {
		coords := findVents(line)
		grid = markCoordsOnGrid(grid, coords)
	}
	for _, line := range grid {
		fmt.Println(line)
	}
	overlapping := findOverlapping(grid)
	fmt.Println(overlapping)
}

func findOverlapping(grid [SIZE][SIZE]int) int {
	ctr := 0
	for _, line := range grid {
		for _, cell := range line {
			if cell >= 2 {
				ctr++
			}
		}
	}
	return ctr
}

func markCoordsOnGrid(grid [SIZE][SIZE]int, coords [][]int) [SIZE][SIZE]int {
	for _, coord := range coords {
		grid[coord[1]][coord[0]] += 1
	}
	return grid
}

func findVents(line string) [][]int {
	splitLine := strings.Split(line, " -> ")
	start := strings.Split(splitLine[0], ",")
	end := strings.Split(splitLine[1], ",")
	coords := getCoords(start, end)
	return coords
}

func getCoords(start []string, end []string) [][]int {
	startIntX, _ := helpers.ToInt(start[0])
	startIntY, _ := helpers.ToInt(start[1])
	endIntX, _ := helpers.ToInt(end[0])
	endIntY, _ := helpers.ToInt(end[1])
	coords := [][]int{}
	if startIntX == endIntX || startIntY == endIntY {
		if startIntX < endIntX {
			fmt.Println(startIntX, startIntY, endIntX, endIntY)
			for x := startIntX; x < endIntX+1; x++ {
				coord := []int{}
				coord = append(coord, x)
				coord = append(coord, startIntY)
				coords = append(coords, coord)
			}
			return coords
		}

		if startIntX > endIntX {
			fmt.Println(startIntX, startIntY, endIntX, endIntY)

			for x := startIntX; x > endIntX-1; x-- {
				coord := []int{}
				coord = append(coord, x)
				coord = append(coord, startIntY)
				coords = append(coords, coord)
			}
			return coords
		}

		if startIntY < endIntY {
			fmt.Println(startIntX, startIntY, endIntX, endIntY)

			for y := startIntY; y < endIntY+1; y++ {
				coord := []int{}
				coord = append(coord, startIntX)
				coord = append(coord, y)
				coords = append(coords, coord)
			}
			return coords
		}

		if startIntY > endIntY {
			fmt.Println(startIntX, startIntY, endIntX, endIntY)

			for y := startIntY; y > endIntY-1; y-- {
				coord := []int{}
				coord = append(coord, startIntX)
				coord = append(coord, y)
				coords = append(coords, coord)
			}
			return coords
		}
		return coords
	}
	//		if x and y increase
	if startIntX < endIntX && startIntY < endIntY {
		fmt.Println(startIntX, startIntY, endIntX, endIntY)

		for x := startIntX; x < endIntX+1; x++ {
			coord := []int{}
			coord = append(coord, x)
			coord = append(coord, startIntY)
			coords = append(coords, coord)
			startIntY++
		}
		return coords
	}
	// //    if x and y decrease
	if startIntX > endIntX && startIntY > endIntY {
		fmt.Println(startIntX, startIntY, endIntX, endIntY)

		for x := startIntX; x > endIntX-1; x-- {
			coord := []int{}
			coord = append(coord, x)
			coord = append(coord, startIntY)
			coords = append(coords, coord)
			startIntY--
		}
		return coords
	}
	// // // if x increases and y decreases
	if startIntX < endIntX && startIntY > endIntY {
		for y := startIntY; y > endIntY-1; y-- {
			coord := []int{}
			coord = append(coord, startIntX)
			coord = append(coord, y)
			coords = append(coords, coord)
			startIntX++
		}
		return coords
	}
	// // // // if y increases and x decreases
	if startIntX > endIntX && startIntY < endIntY {
		fmt.Println(startIntX, startIntY, endIntX, endIntY)

		for y := startIntY; y < endIntY+1; y++ {
			coord := []int{}
			coord = append(coord, startIntX)
			coord = append(coord, y)
			coords = append(coords, coord)
			startIntX--
		}
		return coords
	}
	return coords
}

func ventGrid() [SIZE][SIZE]int {
	grid := [SIZE][SIZE]int{}
	return grid
}
