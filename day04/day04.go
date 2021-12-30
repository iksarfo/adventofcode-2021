package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Filepath struct {
	location string
}

type board = [][]*int
type boards = map[int]board
type line = []*int
type ints = []int

func main() {
	filepath := Filepath{location: "day04/input.txt"}
	numbers := filepath.numbers()
	boards := filepath.boards()

	println(partOne(numbers, boards))
	println(partTwo(numbers, boards))
}

func partOne(numbers ints, boards boards) int {
	for _, number := range numbers {
		for _, board := range boards {
			numberDrawn(board, number)
			for y := 0; y < len(board); y++ {
				_, done := checkLine(board[y])
				if done == true {
					return sumRemaining(board, number)
				}
			}
			for x := 0; x < len(board); x++ {
				_, done := checkVertical(board, x)
				if done == true {
					return sumRemaining(board, number)
				}
			}
		}
	}

	panic("Failed part one")
}

func partTwo(numbers ints, boards boards) int {
	completed := make(map[int]bool)

	for _, number := range numbers {
		for bi, board := range boards {
			numberDrawn(board, number)
			for y := 0; y < len(board); y++ {
				_, done := checkLine(board[y])
				if done == true {
					completed[bi] = true
				}
				if len(completed) == len(boards) {
					return sumRemaining(board, number)
				}
			}
			for x := 0; x < len(board); x++ {
				_, done := checkVertical(board, x)
				if done == true {
					completed[bi] = true
				}
				if len(completed) == len(boards) {
					return sumRemaining(board, number)
				}
			}
		}
	}

	panic("Failed part two")
}

func sumRemaining(board board, number int) int {
	dim := len(board)

	sum := 0
	for y := 0; y < dim; y++ {
		s, _ := checkLine(board[y])
		sum += s
	}

	return sum * number
}

func numberDrawn(board board, n int) {
	for _, line := range board {
		for x, digit := range line {
			if digit == nil {
				continue
			}
			found := *digit
			if found == n {
				line[x] = nil
			}
		}
	}
}

func checkLine(line line) (int, bool) {
	done := true
	result := 0
	for _, number := range line {
		if number != nil {
			result += *number
			done = false
		}
	}
	return result, done
}

func checkVertical(board board, i int) (int, bool) {
	var result []*int
	for _, rank := range board {
		number := rank[i]
		result = append(result, number)
	}
	return checkLine(result)
}

func populateBoard(source []string) board {
	var board board

	zero := 0
	for y := 0; y < len(source); y++ {
		var xs line
		for x := 0; x < len(source); x++ {
			xs = append(xs, &zero)
		}
		board = append(board, xs)
	}

	for y, line := range source {
		nums := strings.Split(line, " ")

		d := 0
		for _, num := range nums {
			if len(num) > 0 {
				digit, _ := strconv.Atoi(num)
				ys := board[y]
				ys[d] = &digit
				d += 1
			}
		}
	}

	return board
}

func (filepath Filepath) numbers() ints {
	file, _ := os.Open(filepath.location)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ints ints
	for scanner.Scan() {
		line := scanner.Text()
		csv := strings.Split(line, ",")

		if len(csv) == 1 {
			continue
		}

		for _, number := range csv {
			integer, _ := strconv.Atoi(number)
			ints = append(ints, integer)
		}
	}
	return ints
}

func (filepath Filepath) boards() boards {
	boards := make(map[int]board)
	contents := filepath.boardText()

	var lines []string
	for _, line := range contents {
		if len(line) < 1 && len(lines) > 0 {
			board := populateBoard(lines)
			boards[len(boards)] = board
			lines = []string{}
			continue
		}

		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	board := populateBoard(lines)
	boards[len(boards)] = board

	return boards
}

func (filepath Filepath) boardText() []string {
	var contents []string

	file, _ := os.Open(filepath.location)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		csv := strings.Split(line, ",")

		if len(csv) > 1 {
			continue
		}

		contents = append(contents, line)
	}
	return contents
}
