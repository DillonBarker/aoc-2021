package main

import (
	"aoc/commons/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Fish struct {
	timer int
	new   bool
}

func main() {
	fishes := getFishes()
	for day := 1; day <= 80; day++ {
		for i, fish := range fishes {
			if fish.timer == 0 {
				if fish.new {
					fishes[i].timer = 6
					fishes[i].new = false
					fishes = append(fishes, newFish())
					continue
				} else {
					fishes[i].timer = 6
					fishes = append(fishes, newFish())
					continue
				}
			}
			fishes[i].timer--
		}
	}
	fmt.Println(len(fishes))
}

func newFish() *Fish {
	var fish = new(Fish)
	fish.timer = 8
	fish.new = true
	return fish
}

func getFishes() []*Fish {
	input, _ := os.Open("./day6/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	sc.Scan()
	inputs := strings.Split(sc.Text(), ",")
	fishes := []*Fish{}
	for _, input := range inputs {
		inputAsInt, _ := helpers.ToInt(input)
		var fish = new(Fish)
		fish.timer = inputAsInt
		fish.new = false
		fishes = append(fishes, fish)
	}
	return fishes
}
