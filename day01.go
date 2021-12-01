package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day01/input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	taken := false
	current := 0
	previous := 0
	increases := 0

	for scanner.Scan() {
		if !taken {
			i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
			previous = int(i)
			taken = true
			continue
		}

		i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		current = int(i)

		if current > previous {
			increases += 1
		}

		previous = current
	}

	fmt.Println(increases)
}
