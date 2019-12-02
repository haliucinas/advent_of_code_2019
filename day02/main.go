package main

import (
	"fmt"

	aoc "github.com/haliucinas/advent_of_code_2019"
)

func main() {
	input, err := aoc.GetDayInput(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := aoc.InputToSlice(input, ",")
	digits, err := aoc.SliceToDigits(slice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Part1: %d\n", part1(digits))
	fmt.Printf("Part2: %d\n", part2(digits))
}

func part1(input []int) int {
	noun := 12
	verb := 2
	return intCode(noun, verb, input)
}

func part2(input []int) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if intCode(noun, verb, input) == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}

func intCode(noun, verb int, code []int) int {
	code = append(code[:0:0], code...)
	ip := 0
	code[1] = noun
	code[2] = verb
	for code[ip] != 99 {
		switch code[ip] {
		case 1:
			code[code[ip+3]] = code[code[ip+1]] + code[code[ip+2]]
			ip += 4
		case 2:
			code[code[ip+3]] = code[code[ip+1]] * code[code[ip+2]]
			ip += 4
		}
	}
	return code[0]
}
