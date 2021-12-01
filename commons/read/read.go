package read

import (
	"bufio"
	"os"
)

func ReadFile(day string) ([]string, error) {
	file, err := os.Open("../../" + day + "/input.txt")
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
