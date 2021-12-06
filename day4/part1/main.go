package main

import (
	"aoc/commons/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs := getInputs()
	boards := getBoards()
	currentInputs := []string{}
	for _, input := range inputs {
		currentInputs = append(currentInputs, input)
		for _, board := range boards {
			if checkIfWinningBoard(board, currentInputs) {
				fmt.Println("we have the winning board")
				calcAns(board, currentInputs)
			}
		}
	}
}

func sum(array []string) int {
	result := 0
	for _, v := range array {
		resultAsInt, _ := helpers.ToInt(v)
		result += resultAsInt
	}
	return result
}

func calcAns(board []string, currentInputs []string) {
	nums := []string{}
	for _, number := range board {
		if !contains(currentInputs, number) {
			nums = append(nums, number)
		}
	}
	lastInput, _ := helpers.ToInt(currentInputs[len(currentInputs)-1])
	fmt.Println(sum(nums) * lastInput)
}

func getInputs() []string {
	input, _ := os.Open("./day4/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	sc.Scan()
	inputs := strings.Split(sc.Text(), ",")
	return inputs
}

func getBoards() [][]string {
	input, _ := os.Open("./day4/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	sc.Scan()
	sc.Scan()

	boards := [][]string{}

	for x := 0; x < 100; x++ {
		board := []string{}
		for i := 0; i < 5; i++ {
			sc.Scan()
			if sc.Text() == "" {
				sc.Scan()
			}
			board = append(board, sc.Text())
		}
		splitBoard := []string{}
		for _, line := range board {
			splitLine := strings.Fields(line)
			splitBoard = append(splitBoard, splitLine...)
		}
		boards = append(boards, splitBoard)
	}
	return boards
}

func checkIfWinningBoard(board []string, currentInputs []string) bool {
	if check(board[0:5], currentInputs) {
		return true
	}
	if check(board[5:10], currentInputs) {
		return true
	}
	if check(board[10:15], currentInputs) {
		return true
	}
	if check(board[15:20], currentInputs) {
		return true
	}
	if check(board[20:25], currentInputs) {
		return true
	}
	return false
}

func check(board []string, currentInputs []string) bool {
	score := 0
	for _, input := range currentInputs {
		if contains(board, input) {
			score++
		}
	}
	return score == 5
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// Winning Conditions
// 22 13 17 11  0
//  8  2 23  4 24
// 21  9 14 16  7
//  6 10  3 18  5
//  1 12 20 15 19

// row
// 0:4, 5:9, 10:14, 15:19, 20:24
// column
// 0,5,10,15,20
// 1,6,11,16,21
// 2,7,12,17,22
// 3,8,13,18,23
// 4,9,14,19,24

// maybe go back to this approach if needed

// func createBoards() []map[string]string {
// 	boardMaps := []map[string]string{}
// 	boards := getBoards()
// 	fmt.Println(boards[0])
// 	for x := 0; x < len(boards); x++ {
// 		boardMap := make(map[string]string)
// 		for i := 0; i < 25; i++ {
// 			boardMap[boards[x][i]] = "test"
// 		}
// 		fmt.Println(boardMap)
// 		boardMaps = append(boardMaps, boardMap)
// 	}
// 	return boardMaps
// }
