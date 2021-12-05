package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	input_len := len(input)
	line_len := len(input[0])
	var counts []int
	for i := 0; i < line_len; i++ {
		counts = append(counts, 0)
	}
	for _, line := range input {
		for i, char := range strings.Split(line, "") {
			char_int, _ := strconv.Atoi(char)
			counts[i] += char_int
		}
	}
	gamma_rate := ""
	epsilon_rate := ""
	for _, c := range counts {
		if c > (input_len / 2) {
			gamma_rate += "1"
			epsilon_rate += "0"
		} else {
			gamma_rate += "0"
			epsilon_rate += "1"
		}
	}
	gamma_rate_int, _ := strconv.ParseInt(gamma_rate, 2, 64)
	epsilon_rate_int, _ := strconv.ParseInt(epsilon_rate, 2, 64)
	fmt.Println(gamma_rate_int * epsilon_rate_int)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return lines, err
		}
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}
