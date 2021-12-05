package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	depth    int
	position int
}

func main() {
	input, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	submarine := Submarine{depth: 0, position: 0}
	for _, line := range input {
		split_line := strings.Fields(line)
		instruction := split_line[0]
		num, err := strconv.Atoi(split_line[1])
		if err != nil {
			log.Fatalf("Atoi: %s", err)
		}
		switch instruction {
		case "forward":
			submarine.position += num
		case "down":
			submarine.depth += num
		case "up":
			submarine.depth -= num
		}
	}
	fmt.Println(submarine.depth * submarine.position)
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
