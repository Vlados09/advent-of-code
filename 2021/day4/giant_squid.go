package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	entries  [](map[int]bool) // Use list of sets to
	n_marked []int
}

const BOARD_SIZE = 5

func main() {
	inputs, boards, err := readBoards("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	final_input, win_board := get_winning_board(inputs, boards)
	used_inputs := make(map[int]bool)
	i := 0
	input := inputs[i]
	for input != final_input {
		used_inputs[input] = true
		i++
		input = inputs[i]
	}
	used_inputs[final_input] = true
	total := 0
	for i := 0; i < BOARD_SIZE; i++ {
		entry := win_board.entries[i]
		for key := range entry {
			if !used_inputs[key] {
				total += key
			}
		}
	}
	fmt.Println(total * final_input)
}

func get_winning_board(inputs []int, boards []Board) (int, Board) {
	for _, input := range inputs {
		for _, board := range boards {
			for i, entry := range board.entries {
				if entry[input] {
					board.n_marked[i] += 1
					if board.n_marked[i] == BOARD_SIZE {
						return input, board
					}
				}
			}
		}
	}
	return inputs[len(inputs)], Board{}
}

func readBoards(path string) ([]int, []Board, error) {
	/**
		Read all of the boards in memory
	**/
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var boards []Board
	var inputs []int
	var board_arr [][]int

	scanner := bufio.NewScanner(file)

	n := 0
	for scanner.Scan() {
		line := scanner.Text()

		if n == 0 {
			inputs = string_to_int_arr(line, ",")
		} else if len(line) == 0 && len(board_arr) > 0 {
			new_board := build_board(board_arr)
			boards = append(boards, new_board)
			board_arr = nil
		} else if len(line) > 0 {
			row := string_to_int_arr(line, " ")
			board_arr = append(board_arr, row)
		}
		n++
	}
	if len(board_arr) > 0 {
		final_board := build_board(board_arr)
		boards = append(boards, final_board)
	}
	return inputs, boards, scanner.Err()
}

func build_board(board_arr [][]int) Board {
	/**
		Given a 2d array convert each row and column to set as
		defined by board representation
	**/
	var entries []map[int]bool
	var n_marked []int
	n_col := len(board_arr[0])
	for i := 0; i < n_col; i++ {
		col_entry := make(map[int]bool)
		entries = append(entries, col_entry)
		n_marked = append(n_marked, 0)
	}
	for _, row := range board_arr {
		row_entry := make(map[int]bool)
		for j, r := range row {
			row_entry[r] = true
			entries[j][r] = true
		}
		entries = append(entries, row_entry)
		n_marked = append(n_marked, 0)
	}
	return Board{entries: entries, n_marked: n_marked}
}

func string_to_int_arr(str string, delimiter string) (arr []int) {
	/**
		Convert list of ints encoded in sting to slice of integers
	**/
	var str_arr []string
	if delimiter == " " {
		str_arr = strings.Fields(str)
	} else {
		str_arr = strings.Split(str, delimiter)
	}
	var int_arr []int
	for _, num := range str_arr {
		num_int, _ := strconv.Atoi(num)
		int_arr = append(int_arr, num_int)
	}
	return int_arr
}
