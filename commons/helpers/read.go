package helpers

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFileIntoLines(day string) ([]string, error) {
	file, err := os.Open(day + "/input.txt")
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

func ToInt(str string) (int, error) {
	lineAsInt, err := strconv.Atoi(str)
	if err != nil {
		os.Exit(2)
	}
	return lineAsInt, err
}
