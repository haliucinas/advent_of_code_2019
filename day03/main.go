package main

import (
	"fmt"
	"math"
	"strconv"

	aoc "github.com/haliucinas/advent_of_code_2019"
)

type Point struct {
	X int
	Y int
}

var (
	dx = map[rune]int{'L': -1, 'R': 1, 'U': 0, 'D': 0}
	dy = map[rune]int{'L': 0, 'R': 0, 'U': 1, 'D': -1}
)

func main() {
	input, err := aoc.GetDayInput(3)
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := aoc.InputToSlice(input, "\n")
	A := aoc.InputToSlice(slice[0], ",")
	B := aoc.InputToSlice(slice[1], ",")
	AP, err := transform(A)
	if err != nil {
		fmt.Println(err)
		return
	}
	BP, err := transform(B)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Part1: %d\n", part1(AP, BP))
	fmt.Printf("Part2: %d\n", part2(AP, BP))
}

func part1(A, B map[Point]int) int {
	min := math.MaxInt32
	for k, v := range A {
		if v > 0 && B[k] > 0 {
			dist := aoc.Abs(k.X) + aoc.Abs(k.Y)
			if dist < min {
				min = dist
			}
		}
	}
	return min
}

func part2(A, B map[Point]int) int {
	min := math.MaxInt32
	for k, v := range A {
		if v > 0 && B[k] > 0 {
			dist := v + B[k]
			if dist < min {
				min = dist
			}
		}
	}
	return min
}

func transform(args []string) (map[Point]int, error) {
	m := make(map[Point]int)
	x, y, step := 0, 0, 0
	for _, arg := range args {
		cmd := rune(arg[0])
		param, err := strconv.Atoi(arg[1:])
		if err != nil {
			return nil, err
		}
		for i := 0; i < param; i++ {
			step++
			x += dx[cmd]
			y += dy[cmd]
			m[Point{x, y}] = step
		}
	}
	return m, nil
}
