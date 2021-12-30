package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	report := fileContents("day03/input.txt")
	gamma := calculateGamma(report)
	epsilon := mapEpsilon(gamma)
	fmt.Println("part 1 =", intArrayToDecimal(gamma)*intArrayToDecimal(epsilon))

	oxygenRating := calculateRating(report, 1)
	co2rating := calculateRating(report, 0)
	fmt.Println("part 2 =", stringToDecimal(oxygenRating)*stringToDecimal(co2rating))
}

func stringToDecimal(text string) int64 {
	decimal, _ := strconv.ParseInt(text, 2, 64)
	return decimal
}

func intArrayToDecimal(arr []int) int64 {
	var binary string
	for i := range arr {
		binary += strconv.Itoa(arr[i])
	}
	return stringToDecimal(binary)
}

func fileContents(filepath string) []string {
	var contents []string

	file, _ := os.Open(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	return contents
}

func Abs(number int) int {
	if number > 0 {
		return number
	} else {
		return -number
	}
}

func getsKept(keep int, ones int, zeros int) int {
	if ones > zeros || ones == zeros {
		return keep
	} else {
		return Abs(keep - 1)
	}
}

func calculateRating(report []string, keep int) string {
	var rating []string

	for len(report) != 1 {
		row := report[0]

		for x, _ := range row {
			var ones int

			for _, line := range report {
				bitstr := string(line[x])
				bit, _ := strconv.Atoi(bitstr)
				ones += bit
			}

			lineCount := len(report)
			zeros := lineCount - ones
			keeping := getsKept(keep, ones, zeros)
			rating = nil

			for _, line := range report {
				bitstr := string(line[x])
				bit, _ := strconv.Atoi(bitstr)

				if bit == keeping {
					rating = append(rating, line)
				}
			}

			report = rating

			if len(rating) == 1 {
				return rating[0]
			}
		}
	}
	return report[0]
}

func calculateGamma(report []string) []int {
	var power []int
	var lines int

	for _, consumption := range report {
		lines += 1

		if len(power) == 0 {
			power = make([]int, len(consumption))
		}

		bits := []rune(consumption)

		for i := 0; i < len(bits); i++ {
			bit, _ := strconv.Atoi(string(bits[i]))
			power[i] += bit
		}
	}
	return mapGamma(power, lines/2)
}

func mapGamma(totalConsumption []int, threshold int) []int {
	mapped := make([]int, len(totalConsumption), len(totalConsumption))

	for i := 0; i < len(mapped); i++ {
		if totalConsumption[i] > threshold {
			mapped[i] = 1
		}
	}
	return mapped
}

func mapEpsilon(gamma []int) []int {
	mapped := make([]int, len(gamma), len(gamma))

	for i := 0; i < len(mapped); i++ {
		if gamma[i] == 0 {
			mapped[i] = 1
		}
	}
	return mapped
}
