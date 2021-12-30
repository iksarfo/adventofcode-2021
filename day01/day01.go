package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1, _ := part01()
	fmt.Println("part 1 = ", *part1)

	part2, _ := part02()
	fmt.Println("part 2 = ", *part2)
}

func part01() (*int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var previous, current *int
	var increases int

	for scanner.Scan() {
		if previous == nil {
			i, _ := strconv.Atoi(scanner.Text())
			previous = &i
			continue
		}

		next, _ := strconv.Atoi(scanner.Text())

		current = &next

		if *current > *previous {
			increases += 1
		}

		previous = current
	}

	return &increases, nil
}

type que struct {
	window []int
}

func (q *que) push(item int, max int) {
	if len(q.window) < max {
		q.window = append(q.window, item)
		return
	}

	for i := 1; i < max; i++ {
		q.window[i-1] = q.window[i]
	}

	q.window[max-1] = item
}

func part02() (*int, error) {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var increases int
	var items que
	windowSize := 4

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())

		items.push(i, windowSize)

		if len(items.window) < windowSize {
			continue
		}

		if items.window[windowSize-1] > items.window[0] {
			increases += 1
		}
	}

	return &increases, nil
}
