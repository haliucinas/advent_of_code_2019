package main

import (
	"fmt"
	"math"
	"strconv"

	aoc "github.com/haliucinas/advent_of_code_2019"
)

func main() {
	input, err := aoc.GetDayInput(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Part1: %d\n", part1(aoc.InputToSlice(input)))
}

func part1(input []string) (sum int) {
	for _, item := range input {
		num, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println(err)
			return
		}

		sum += int(math.Floor(float64(num)/3) - 2)
	}
	return sum
}
