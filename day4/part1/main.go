package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inputs := getInputs()
	fmt.Println(inputs)
	getBoards()
}

// this is missing the last input
func getInputs() []string {
	content, _ := ioutil.ReadFile("day4/input.txt")
	a := strings.Split(string(content), ",")

	inputs := []string{}
	for _, item := range a {
		l := strings.Split(string(item), "\n")
		if len(l) == 1 {
			inputs = append(inputs, item)
		}
	}
	return inputs
}

// how da faq do i do dis
func getBoards() {
	content, _ := ioutil.ReadFile("day4/input.txt")
	a := strings.Split(string(content), ",")

	fmt.Println(a)
}
