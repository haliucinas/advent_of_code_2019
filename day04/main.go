package main

import (
	"fmt"
	"strconv"

	aoc "github.com/haliucinas/advent_of_code_2019"
)

func main() {
	input, err := aoc.GetDayInput(4)
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := aoc.InputToSlice(input, "-")
	start, err := strconv.Atoi(slice[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	end, err := strconv.Atoi(slice[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Part1: %d\n", part1(start, end))
	fmt.Printf("Part2: %d\n", part2(start, end))
}

func part1(start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if ok, _ := condition(i); ok {
			count++
		}
	}
	return count
}

func part2(start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if _, ok := condition(i); ok {
			count++
		}
	}
	return count
}

func condition(n int) (p1 bool, p2 bool) {
	prev := n % 10
	count := map[int]int{prev: 1}

	for n > 9 {
		n /= 10
		curr := n % 10

		count[curr]++

		if curr > prev {
			return false, false
		}

		if curr == prev {
			p1 = true
		}

		prev = curr
	}

	for _, c := range count {
		if c == 2 {
			p2 = true
		}
	}

	return p1, p2
}
