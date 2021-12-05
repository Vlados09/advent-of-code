package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count += 1
		}
	}
	fmt.Println(count)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return lines, err
		}
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}
