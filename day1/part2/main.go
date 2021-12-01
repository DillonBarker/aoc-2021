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
		if i == 2000 || i == 1998 || i == 1999 || i == 1997 {
			continue
		}

		lineAsInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		lineAsInt2, err := strconv.Atoi(lines[i+1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		lineAsInt3, err := strconv.Atoi(lines[i+2])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		lineTotal := lineAsInt + lineAsInt2 + lineAsInt3

		nextLineAsInt, err := strconv.Atoi(lines[i+1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		nextLineAsInt2, err := strconv.Atoi(lines[i+2])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		nextLineAsInt3, err := strconv.Atoi(lines[i+3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		nextLineTotal := nextLineAsInt + nextLineAsInt2 + nextLineAsInt3

		if nextLineTotal > lineTotal {
			ctr++
		}

	}
	fmt.Println(ctr)
}
