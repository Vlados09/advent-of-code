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
	oxygen_rating := get_rating(input, true)
	co2_rating := get_rating(input, false)
	fmt.Println(co2_rating * oxygen_rating)
}

func get_rating(arr [][]int, popular bool) int64 {
	idx := 0
	for len(arr) > 1 {
		arr = filter_by_bit_at_index(arr, idx, popular)
		idx += 1
	}
	final_binary := ""
	for _, bit := range arr[0] {
		final_binary += strconv.Itoa(bit)
	}
	final_int, _ := strconv.ParseInt(final_binary, 2, 64)
	return final_int
}

func filter_by_bit_at_index(arr [][]int, idx int, popular bool) [][]int {
	popular_bit, arr0, arr1 := popular_bit_at_index(arr, idx)
	if !popular {
		arr0, arr1 = arr1, arr0
	}
	if popular_bit == 1 || popular_bit == -1 {
		return arr1
	} else {
		return arr0
	}
}

func popular_bit_at_index(arr [][]int, idx int) (int, [][]int, [][]int) {

	count0 := 0
	count1 := 0
	var arr0 [][]int
	var arr1 [][]int
	var popular_bit int

	for _, sub_arr := range arr {
		num := sub_arr[idx]
		if num == 0 {
			arr0 = append(arr0, sub_arr)
			count0 += 1
		} else {
			arr1 = append(arr1, sub_arr)
			count1 += 1
		}
	}

	if count0 > count1 {
		popular_bit = 0
	} else if count0 < count1 {
		popular_bit = 1
	} else {
		popular_bit = -1
	}

	return popular_bit, arr0, arr1
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line_arr_char := strings.Split(scanner.Text(), "")
		var line_arr_int []int
		for _, c := range line_arr_char {
			c_i, _ := strconv.Atoi(c)
			line_arr_int = append(line_arr_int, c_i)
		}
		lines = append(lines, line_arr_int)
	}
	return lines, scanner.Err()
}
