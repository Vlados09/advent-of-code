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
	for i := 3; i < len(input); i++ {
		first_sum := sumArray(input[i-3 : i])
		second_sum := sumArray(input[i-2 : i+1])
		if second_sum > first_sum {
			count += 1
		}
	}
	fmt.Println(count)
}

func sumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
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
