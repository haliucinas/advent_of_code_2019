package main

import (
	"fmt"

	aoc "github.com/haliucinas/advent_of_code_2019"
)

func main() {
	input, err := aoc.GetDayInput(5)
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
	return intCode(1, input)
}

func part2(input []int) int {
	return intCode(5, input)
}

func intCode(input int, code []int) int {
	code = append(code[:0:0], code...)
	out := make([]int, 0)
	ip := 0
	for ip > -1 {
		p1m, p2m := code[ip]/100%10 == 0, code[ip]/1000%10 == 0
		switch code[ip] % 100 {
		case 1:
			code[code[ip+3]] = arg(code, ip+1, p1m) + arg(code, ip+2, p2m)
			ip += 4
		case 2:
			code[code[ip+3]] = arg(code, ip+1, p1m) * arg(code, ip+2, p2m)
			ip += 4
		case 3:
			code[code[ip+1]] = input
			ip += 2
		case 4:
			out = append(out, code[code[ip+1]])
			ip += 2
		case 5:
			if arg(code, ip+1, p1m) > 0 {
				ip = arg(code, ip+2, p2m)
				continue
			}
			ip += 3
		case 6:
			if arg(code, ip+1, p1m) == 0 {
				ip = arg(code, ip+2, p2m)
				continue
			}
			ip += 3
		case 7:
			code[code[ip+3]] = aoc.BoolToInt(arg(code, ip+1, p1m) < arg(code, ip+2, p2m))
			ip += 4
		case 8:
			code[code[ip+3]] = aoc.BoolToInt(arg(code, ip+1, p1m) == arg(code, ip+2, p2m))
			ip += 4
		case 99:
			ip = -1
		}
	}
	if len(out) == 0 {
		return 0
	}
	return out[len(out)-1]
}

func arg(code []int, ip int, mode bool) int {
	if mode {
		return code[code[ip]]
	}
	return code[ip]
}
