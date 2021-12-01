package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("../input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	ctr := 0
	for i, line := range lines {
		if i == 0 {
			continue
		}

		lineAsInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		previousLineAsInt, err := strconv.Atoi(lines[i-1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if lineAsInt > previousLineAsInt {
			ctr++
		}
	}
	fmt.Println(ctr)
}
