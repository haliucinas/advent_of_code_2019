package main

import (
	"fmt"
	"math"

	aoc "github.com/haliucinas/advent_of_code_2019"
)

func main() {
	input, err := aoc.GetDayInput(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := aoc.InputToSlice(input, "\n")
	digits, err := aoc.SliceToDigits(slice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Part1: %d\n", part1(digits))
	fmt.Printf("Part2: %d\n", part2(digits))
}

func part1(input []int) (sum int) {
	for _, num := range input {
		sum += transform(num)
	}
	return sum
}

func part2(input []int) (sum int) {
	for _, num := range input {
		fuel := transform(num)
		for fuel > 0 {
			sum += fuel
			fuel = transform(fuel)
		}
	}
	return sum
}

func transform(num int) int {
	return int(math.Floor(float64(num)/3) - 2)
}
